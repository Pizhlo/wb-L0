package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_postgres "github.com/Pizhlo/wb-L0/internal/app/storage/postgres/mocks"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func runTestServer(t *testing.T, handler *Order) (chi.Router, *gomock.Controller, *mock_postgres.MockRepo) {
	r := chi.NewRouter()

	ctrl := gomock.NewController(t)
	db := mock_postgres.NewMockRepo(ctrl)

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetOrderByID(w, r)
	})

	return r, ctrl, db
}

func testRequest(t *testing.T, ts *httptest.Server, method,
	path string, body io.Reader) *http.Response {

	req, err := http.NewRequest(method, ts.URL+path, body)

	req.Close = true
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "PostmanRuntime/7.32.3")

	require.NoError(t, err)

	ts.Client()

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	return resp
}

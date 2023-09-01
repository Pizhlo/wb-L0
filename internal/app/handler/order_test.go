package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	"github.com/Pizhlo/wb-L0/internal/app/stan/data"
	"github.com/Pizhlo/wb-L0/internal/app/storage/cache"
	mock_postgres "github.com/Pizhlo/wb-L0/internal/app/storage/postgres/mocks"
	models "github.com/Pizhlo/wb-L0/internal/model"
	"github.com/Pizhlo/wb-L0/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetOrderByID_InvalidUUID(t *testing.T) {
	tests := []struct {
		name       string
		request    string
		method     string
		statusCode int
	}{
		{
			name:       "invalid UUID 123",
			request:    "/123",
			method:     http.MethodGet,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid UUID abc",
			request:    "/abc",
			method:     http.MethodGet,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid UUID asdaslk19280",
			request:    "/asdaslk19280",
			method:     http.MethodGet,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		db := mock_postgres.NewMockRepo(ctrl)

		order := cache.NewOrder()
		cache := cache.New(order)

		service := service.New(db, cache)
		handler := NewOrder(*service)

		r, ctrl, _ := runTestServer(t, handler)
		defer ctrl.Finish()

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp := testRequest(t, ts, tt.method, tt.request, nil)
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)
	}
}

func TestGetOrderByID_NotFound(t *testing.T) {
	tests := []struct {
		request    string
		idString   string
		method     string
		statusCode int
	}{
		{
			request:    "/106a2296-ef19-46ba-9d26-0f6a8ee617df",
			idString:   "106a2296-ef19-46ba-9d26-0f6a8ee617df",
			method:     http.MethodGet,
			statusCode: http.StatusNotFound,
		},
		{
			request:    "/e2620fc5-4675-4f85-859a-b89e9a689f9e",
			idString:   "e2620fc5-4675-4f85-859a-b89e9a689f9e",
			method:     http.MethodGet,
			statusCode: http.StatusNotFound,
		},
		{
			request:    "/cac37265-198a-4a7d-a716-8d95e11a2156",
			idString:   "cac37265-198a-4a7d-a716-8d95e11a2156",
			method:     http.MethodGet,
			statusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		id := uuid.MustParse(tt.idString)

		ctrl := gomock.NewController(t)
		db := mock_postgres.NewMockRepo(ctrl)

		order := cache.NewOrder()
		cache := cache.New(order)

		service := service.New(db, cache)
		handler := NewOrder(*service)

		r, ctrl, db := runTestServer(t, handler)
		defer ctrl.Finish()

		ts := httptest.NewServer(r)
		defer ts.Close()

		db.EXPECT().GetOrderByID(gomock.Any(), id).Return(&models.Order{}, errs.NotFound)

		resp := testRequest(t, ts, tt.method, tt.request, nil)
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)
	}
}

func TestGetOrderByID_DBErr(t *testing.T) {
	tests := []struct {
		request    string
		idString   string
		method     string
		statusCode int
	}{
		{
			request:    "/e2620fc5-4675-4f85-859a-b89e9a689f9e",
			idString:   "e2620fc5-4675-4f85-859a-b89e9a689f9e",
			method:     http.MethodGet,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		id := uuid.MustParse(tt.idString)

		ctrl := gomock.NewController(t)
		db := mock_postgres.NewMockRepo(ctrl)

		order := cache.NewOrder()
		cache := cache.New(order)

		service := service.New(db, cache)
		handler := NewOrder(*service)

		r, ctrl, db := runTestServer(t, handler)
		defer ctrl.Finish()

		ts := httptest.NewServer(r)
		defer ts.Close()

		pgErr := pgconn.PgError{}

		db.EXPECT().GetOrderByID(gomock.Any(), id).Return(&models.Order{}, &pgErr)

		resp := testRequest(t, ts, tt.method, tt.request, nil)
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)
	}
}

func TestGetOrderByID_FoundInCache(t *testing.T) {
	tests := []struct {
		request    string
		id         uuid.UUID
		method     string
		cache      *cache.Cache
		cacheNum   int // сколько записей добавить в кэш
		statusCode int
		result     models.Order
	}{
		{
			cacheNum:   3,
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
		{
			cacheNum:   100,
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
		{
			cacheNum:   500,
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		cache, orders := fillCache(tt.cacheNum)
		tt.cache = cache
		tt.result = orders[0] // заполняем кэш несколькими записями и запрашиваем любую из них

		ctrl := gomock.NewController(t)
		db := mock_postgres.NewMockRepo(ctrl)

		service := service.New(db, cache)
		handler := NewOrder(*service)

		tt.id = orders[0].OrderUIID
		tt.request = fmt.Sprintf("/%s", tt.id)

		r, ctrl, _ := runTestServer(t, handler)
		defer ctrl.Finish()

		ts := httptest.NewServer(r)
		defer ts.Close()

		resp := testRequest(t, ts, tt.method, tt.request, nil)
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)

		var orderBody models.Order
		err := json.NewDecoder(resp.Body).Decode(&orderBody)
		require.NoError(t, err)

		// вытаскиваем время из структур, чтобы нормально сравнить
		timeExpected := tt.result.DateCreated
		tt.result.DateCreated = time.Time{}

		timeActual := orderBody.DateCreated
		orderBody.DateCreated = time.Time{}

		assert.Equal(t, tt.result, orderBody)
		assert.True(t, timeExpected.Equal(timeActual))
		assert.Equal(t, "application/json", resp.Header.Get("Content-type"))
	}
}

func TestGetOrderByID_FoundInDB(t *testing.T) {
	tests := []struct {
		request    string
		id         uuid.UUID
		method     string
		statusCode int
	}{
		{
			method:     http.MethodGet,
			id:         uuid.New(),
			statusCode: http.StatusOK,
		},
		{
			method:     http.MethodGet,
			id:         uuid.New(),
			statusCode: http.StatusOK,
		},
		{
			method:     http.MethodGet,
			id:         uuid.New(),
			statusCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		db := mock_postgres.NewMockRepo(ctrl)

		orderCache := cache.NewOrder()
		service := service.New(db, cache.New(orderCache))
		handler := NewOrder(*service)

		r, ctrl, _ := runTestServer(t, handler)
		defer ctrl.Finish()

		ts := httptest.NewServer(r)
		defer ts.Close()

		tt.request = fmt.Sprintf("/%s", tt.id)

		db.EXPECT().GetOrderByID(gomock.Any(), tt.id).Return(&models.Order{
			OrderUIID: tt.id,
		}, nil)

		resp := testRequest(t, ts, tt.method, tt.request, nil)
		defer resp.Body.Close()

		assert.Equal(t, tt.statusCode, resp.StatusCode)

		assert.Equal(t, "application/json", resp.Header.Get("Content-type"))
	}
}

func fillCache(n int) (*cache.Cache, []models.Order) {
	order := cache.NewOrder()
	cache := cache.New(order)

	orders := []models.Order{}
	for i := 0; i < n; i++ {
		randomOrder := data.RandomOrder()
		cache.Order.Save(randomOrder.OrderUIID, randomOrder)
		orders = append(orders, randomOrder)
	}

	return cache, orders
}

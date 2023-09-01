package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"log"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func LoadMainPg(w http.ResponseWriter, r *http.Request, handler Order) {
	//tmpl := template.Must(template.New("").Parse())
}

func GetOrderByID(w http.ResponseWriter, r *http.Request, handler Order) {
	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Default().Printf("uuid %s is invalid: %s", id, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := handler.service.GetOrderByID(r.Context(), uuid)
	if err != nil {
		if errors.Is(err, errs.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Default().Println("error while getting order by ID: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		log.Default().Println("error while making json: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(orderJson)
	if err != nil {
		log.Default().Println("error while writing json: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

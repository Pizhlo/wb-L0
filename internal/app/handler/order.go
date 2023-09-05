package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"log"

	"github.com/Pizhlo/wb-L0/internal/app/errs"
	"github.com/Pizhlo/wb-L0/internal/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Order struct {
	service service.Order
}

func NewOrder(service service.Order) *Order {
	return &Order{service: service}
}

func (h *Order) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		log.Default().Printf("uuid is invalid: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, fmt.Sprintf("uuid is invalid: %s", err), http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, errs.NotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Default().Println("error while getting order by ID: ", err)
		http.Error(w, fmt.Sprintf("error while getting order by ID: %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		log.Default().Println("error while making json: ", err)
		http.Error(w, fmt.Sprintf("error while making json: %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(orderJson)
	if err != nil {
		log.Default().Println("error while writing json: ", err)
		http.Error(w, fmt.Sprintf("error while writing json: %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

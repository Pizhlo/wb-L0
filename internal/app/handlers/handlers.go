package internal

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/Pizhlo/wb-L0/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func GetOrderByID(w http.ResponseWriter, r *http.Request, service service.Service) {
	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Default().Printf("uuid %s is invalid: %s", id, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := service.Storage.GetOrderByID(r.Context(), uuid)
	if err != nil {
		log.Default().Println("error while getting order by ID: ", err)
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		log.Default().Println("error while making json: ", err)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(orderJson)
}

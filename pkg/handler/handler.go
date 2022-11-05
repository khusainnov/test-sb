package handler

import (
	"github.com/gorilla/mux"
	"github.com/khusainnov/sbercloud/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/config", h.Upload).Methods("POST")
	r.HandleFunc("/v1/config", h.GetConfig).Methods("GET")
	r.HandleFunc("/v1/config", h.DeleteConfig).Methods("DELETE")

	return r
}

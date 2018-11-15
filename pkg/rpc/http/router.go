package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter registers handlers to a router.
func NewRouter(h *SWAPIHandlerSet) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/film/{id}", h.FilmHandler).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", h.PersonHandler).Methods(http.MethodGet)
	r.HandleFunc("/planet/{id}", h.PlanetHandler).Methods(http.MethodGet)
	r.HandleFunc("/species/{id}", h.SpeciesHandler).Methods(http.MethodGet)
	r.HandleFunc("/starship/{id}", h.StarshipHandler).Methods(http.MethodGet)
	r.HandleFunc("/vehicle/{id}", h.VehicleHandler).Methods(http.MethodGet)
	return r
}

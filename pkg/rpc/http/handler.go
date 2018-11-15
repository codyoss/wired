package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/codyoss/wired/pkg/swapi"
	"github.com/gorilla/mux"
)

// SWAPIService is an abstraction for a swapi.Service
type SWAPIService interface {
	Film(id int) (*swapi.Film, error)
	Person(id int) (*swapi.Person, error)
	Planet(id int) (*swapi.Planet, error)
	Species(id int) (*swapi.Species, error)
	Starship(id int) (*swapi.Starship, error)
	Vehicle(id int) (*swapi.Vehicle, error)
}

// SWAPIHandlerSet holds a set of `http.Handler`s to retrieve data from a SWAPIService.
type SWAPIHandlerSet struct {
	s SWAPIService
}

// NewSWAPIHandlerSet creates a new SWAPIHandlerSet.
func NewSWAPIHandlerSet(s SWAPIService) *SWAPIHandlerSet {
	return &SWAPIHandlerSet{s}
}

// FilmHandler looks up information about a film given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) FilmHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	f, err := h.s.Film(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(f)
}

// PersonHandler looks up information about a person given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) PersonHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	p, err := h.s.Person(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// PlanetHandler looks up information about a planet given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) PlanetHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	p, err := h.s.Planet(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// SpeciesHandler looks up information about a species given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) SpeciesHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	s, err := h.s.Species(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// StarshipHandler looks up information about a starship given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) StarshipHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	ss, err := h.s.Starship(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(ss)
}

// VehicleHandler looks up information about a vehicle given an id. Results are written out as JSON.
func (h *SWAPIHandlerSet) VehicleHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(w, r)
	if !ok {
		return
	}
	v, err := h.s.Vehicle(id)
	if err != nil {
		handleError(w, err)
		return
	}
	json.NewEncoder(w).Encode(v)
}

func extractID(w http.ResponseWriter, r *http.Request) (int, bool) {
	vars := mux.Vars(r)
	sID := vars["id"]
	if sID == "" {
		handleError(w, errors.New("path parameter `id` missing"))
		return 0, false
	}

	id, err := strconv.Atoi(sID)
	if err != nil {
		handleError(w, errors.New("could not convert id to an int"))
		return 0, false
	}
	return id, true
}

func handleError(w http.ResponseWriter, e error) {
	em := NewErrorMessage(e)
	json.NewEncoder(w).Encode(em)
}

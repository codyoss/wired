package swapi

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/codyoss/wired/pkg/client"
)

var (
	filmURLTemplate     = "https://swapi.co/api/films/%d"
	personURLTemplate   = "https://swapi.co/api/people/%d"
	planetURLTemplate   = "https://swapi.co/api/planets/%d"
	speciesURLTemplate  = "https://swapi.co/api/species/%d"
	starshipURLTemplate = "https://swapi.co/api/starships/%d"
	vehicleURLTemplate  = "https://swapi.co/api/vehicles/%d"
)

// Service provides a way to enteract with the Star Wars API.
type Service struct {
	client client.Client
}

// NewService creates a new swapi service
func NewService(client client.Client) *Service {
	return &Service{
		client: client,
	}
}

// Film returns the film associated in the id provided or an error.
func (s *Service) Film(id int) (*Film, error) {
	f := &Film{}
	err := s.handleRequest(fmt.Sprintf(filmURLTemplate, id), f)
	return f, err
}

// Person returns the person associated in the id provided or an error.
func (s *Service) Person(id int) (*Person, error) {
	p := &Person{}
	err := s.handleRequest(fmt.Sprintf(personURLTemplate, id), p)
	return p, err
}

// Planet returns the planet associated in the id provided or an error.
func (s *Service) Planet(id int) (*Planet, error) {
	p := &Planet{}
	err := s.handleRequest(fmt.Sprintf(planetURLTemplate, id), p)
	return p, err
}

// Species returns the species associated in the id provided or an error.
func (s *Service) Species(id int) (*Species, error) {
	sp := &Species{}
	err := s.handleRequest(fmt.Sprintf(speciesURLTemplate, id), sp)
	return sp, err
}

// Starship returns the starship associated in the id provided or an error.
func (s *Service) Starship(id int) (*Starship, error) {
	ss := &Starship{}
	err := s.handleRequest(fmt.Sprintf(starshipURLTemplate, id), ss)
	return ss, err
}

// Vehicle returns the vehicle associated in the id provided or an error.
func (s *Service) Vehicle(id int) (*Vehicle, error) {
	v := &Vehicle{}
	err := s.handleRequest(fmt.Sprintf(vehicleURLTemplate, id), v)
	return v, err
}

func (s *Service) handleRequest(url string, v interface{}) error {
	b, err := s.client.GET(url)
	if err != nil {
		return err
	}

	err = decodeJSON(b, v)
	if err != nil {
		return err
	}

	return nil
}

func decodeJSON(b *bytes.Buffer, v interface{}) error {
	if v == nil {
		return fmt.Errorf("value of passed in struct must not be nil")
	}

	err := json.NewDecoder(b).Decode(v)
	if err != nil {
		return fmt.Errorf("error reading response")
	}

	return nil
}

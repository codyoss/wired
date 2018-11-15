package swapi_test

import (
	"testing"

	"github.com/codyoss/wired/pkg/swapi"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name  string
		s     *mockSWAPIService
		block func(c *swapi.CachedService)
	}{
		{"film", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Film(1) }},
		{"person", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Person(1) }},
		{"planet", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Planet(1) }},
		{"species", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Species(1) }},
		{"starship", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Starship(1) }},
		{"vehicle", &mockSWAPIService{}, func(c *swapi.CachedService) { c.Vehicle(1) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := swapi.NewCachedService(tt.s)
			if tt.s.Invoked != 0 {
				t.Error("Should have no innvocations")
			}

			tt.block(c)
			if tt.s.Invoked != 1 {
				t.Error("Should have 1 innvocations")
			}

			tt.block(c)
			if tt.s.Invoked != 1 {
				t.Error("Should have cached the result")
			}
		})
	}
}

type mockSWAPIService struct {
	Invoked int
}

func (s *mockSWAPIService) Film(id int) (*swapi.Film, error) {
	s.Invoked++
	return nil, nil
}
func (s *mockSWAPIService) Person(id int) (*swapi.Person, error) {
	s.Invoked++
	return nil, nil
}
func (s *mockSWAPIService) Planet(id int) (*swapi.Planet, error) {
	s.Invoked++
	return nil, nil
}
func (s *mockSWAPIService) Species(id int) (*swapi.Species, error) {
	s.Invoked++
	return nil, nil
}
func (s *mockSWAPIService) Starship(id int) (*swapi.Starship, error) {
	s.Invoked++
	return nil, nil
}
func (s *mockSWAPIService) Vehicle(id int) (*swapi.Vehicle, error) {
	s.Invoked++
	return nil, nil
}

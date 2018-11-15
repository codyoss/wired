// +build integration

package swapi_test

import (
	"testing"

	"github.com/codyoss/wired/pkg/client"
	"github.com/codyoss/wired/pkg/swapi"
)

func TestFilm(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Film(1)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Title != "A New Hope" {
		t.Errorf("got name %s, want %s", p.Title, "A New Hope")
	}
}

func TestPerson(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Person(1)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Name != "Luke Skywalker" {
		t.Errorf("got name %s, want %s", p.Name, "Luke Skywalker")
	}
}

func TestPlanet(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Planet(1)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Name != "Tatooine" {
		t.Errorf("got name %s, want %s", p.Name, "Tatooine")
	}
}

func TestSpecies(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Species(1)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Name != "Human" {
		t.Errorf("got name %s, want %s", p.Name, "Human")
	}
}

func TestStarship(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Starship(2)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Name != "CR90 corvette" {
		t.Errorf("got name %s, want %s", p.Name, "CR90 corvette")
	}
}

func TestVehicle(t *testing.T) {
	client := client.New()
	svc := swapi.NewService(client)

	p, err := svc.Vehicle(4)
	if err != nil {
		t.Errorf("not expecting error: %s", err)
	}

	if p.Name != "Sand Crawler" {
		t.Errorf("got name %s, want %s", p.Name, "Sand Crawler")
	}
}

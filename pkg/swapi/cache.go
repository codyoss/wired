package swapi

import (
	"github.com/hashicorp/golang-lru"
)

// SWAPIService is an abstraction for a swapi.Service
type SWAPIService interface {
	Film(id int) (*Film, error)
	Person(id int) (*Person, error)
	Planet(id int) (*Planet, error)
	Species(id int) (*Species, error)
	Starship(id int) (*Starship, error)
	Vehicle(id int) (*Vehicle, error)
}

// CachedService provides a caching layer for a SWAPI.
type CachedService struct {
	service       SWAPIService
	filmCache     *lru.Cache
	personCache   *lru.Cache
	planetCache   *lru.Cache
	speciesCache  *lru.Cache
	starshipCache *lru.Cache
	vehicleCache  *lru.Cache
}

// NewCachedService is a decorator for a Service the provides caching
func NewCachedService(service SWAPIService) *CachedService {
	filmCache, _ := lru.New(10)
	personCache, _ := lru.New(10)
	planetCache, _ := lru.New(10)
	speciesCache, _ := lru.New(10)
	starshipCache, _ := lru.New(10)
	vehicleCache, _ := lru.New(10)
	return &CachedService{
		service:       service,
		filmCache:     filmCache,
		personCache:   personCache,
		planetCache:   planetCache,
		speciesCache:  speciesCache,
		starshipCache: starshipCache,
		vehicleCache:  vehicleCache,
	}
}

// Film retrives a film from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Film(id int) (*Film, error) {
	f, ok := s.filmCache.Get(id)
	if ok {
		return f.(*Film), nil
	}

	film, err := s.service.Film(id)
	if err != nil {
		return nil, err
	}
	s.filmCache.Add(id, film)
	return film, nil
}

// Person retrives a person from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Person(id int) (*Person, error) {
	p, ok := s.personCache.Get(id)
	if ok {
		return p.(*Person), nil
	}

	person, err := s.service.Person(id)
	if err != nil {
		return nil, err
	}
	s.personCache.Add(id, person)
	return person, nil
}

// Planet retrives a Planet from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Planet(id int) (*Planet, error) {
	p, ok := s.planetCache.Get(id)
	if ok {
		return p.(*Planet), nil
	}

	planet, err := s.service.Planet(id)
	if err != nil {
		return nil, err
	}
	s.planetCache.Add(id, planet)
	return planet, nil
}

// Species retrives a species from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Species(id int) (*Species, error) {
	sp, ok := s.speciesCache.Get(id)
	if ok {
		return sp.(*Species), nil
	}

	species, err := s.service.Species(id)
	if err != nil {
		return nil, err
	}
	s.speciesCache.Add(id, species)
	return species, nil
}

// Starship retrives a starship from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Starship(id int) (*Starship, error) {
	ss, ok := s.starshipCache.Get(id)
	if ok {
		return ss.(*Starship), nil
	}

	starship, err := s.service.Starship(id)
	if err != nil {
		return nil, err
	}
	s.starshipCache.Add(id, starship)
	return starship, nil
}

// Vehicle retrives a vehicle from a cache or makes a call to the underlying service and caches its result.
func (s *CachedService) Vehicle(id int) (*Vehicle, error) {
	v, ok := s.vehicleCache.Get(id)
	if ok {
		return v.(*Vehicle), nil
	}

	vehicle, err := s.service.Vehicle(id)
	if err != nil {
		return nil, err
	}
	s.vehicleCache.Add(id, vehicle)
	return vehicle, nil
}

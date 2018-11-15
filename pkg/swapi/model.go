package swapi

// Film in the Star Wars series.
type Film struct {
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	URL           string   `json:"url"`
}

// A Person within the Star Wars universe.
type Person struct {
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	HairColor    string   `json:"hair_color"`
	SkinColor    string   `json:"skin_color"`
	EyeColor     string   `json:"eye_color"`
	BirthYear    string   `json:"birth_year"`
	Gender       string   `json:"gender"`
	Homeworld    string   `json:"homeworld"`
	FilmURLs     []string `json:"films"`
	SpeciesURLs  []string `json:"species"`
	VehicleURLs  []string `json:"vehicles"`
	StarshipURLs []string `json:"starships"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

// Planet is a big thing in the sky within the Star Wars universe.
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	ResidentURLs   []string `json:"residents"`
	FilmURLs       []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

// A Species is a type of person within the Star Wars universe.
type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	PeopleURLs      []string `json:"people"`
	FilmURLs        []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

// Starship is a flying ship from Star Wars.
type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	PilotURLs            []string `json:"pilots"`
	FilmURLs             []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

// Vehicle is a land vechicle in the Star Wars universe.
type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	PilotURLs            []string `json:"pilots"`
	FilmURLs             []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

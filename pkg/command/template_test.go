package command

import (
	"testing"

	"github.com/codyoss/wired/pkg/swapi"
)

func TestFormatFilm(t *testing.T) {
	f := &swapi.Film{Title: "Title", Director: "Director", Producer: "Producer"}
	names := []string{"Cody"}
	want := `Title was directed by Director and produced by Producer.
It stars the following characters:
- Cody
`
	got := formatFilm(f, names)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFormatPerson(t *testing.T) {
	p := &swapi.Person{Name: "Cody"}
	films := []string{"Film"}
	want := `Cody stars in the following films:
- Film
`
	got := formatPerson(p, films)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

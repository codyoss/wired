package command

import (
	"bytes"
	"html/template"

	"github.com/codyoss/wired/pkg/swapi"
)

const (
	filmTemplate = `{{.Film.Title}} was directed by {{.Film.Director}} and produced by {{.Film.Producer}}.
It stars the following characters:
{{range .Names}}{{println "-" .}}{{end}}`
	personTemplate = `{{.Person.Name}} stars in the following films:
{{range .Films}}{{println "-" .}}{{end}}`
)

func formatFilm(film *swapi.Film, names []string) string {
	var b bytes.Buffer
	data := struct {
		Film  *swapi.Film
		Names []string
	}{
		film,
		names,
	}

	t, err := template.New("film").Parse(filmTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&b, data)
	if err != nil {
		panic(err)
	}

	return b.String()
}

func formatPerson(person *swapi.Person, films []string) string {
	var b bytes.Buffer
	data := struct {
		Person *swapi.Person
		Films  []string
	}{
		person,
		films,
	}

	t, err := template.New("person").Parse(personTemplate)
	if err != nil {
		panic(err)
	}
	err = t.Execute(&b, data)
	if err != nil {
		panic(err)
	}

	return b.String()
}

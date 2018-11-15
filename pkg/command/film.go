package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filmCmd = &cobra.Command{
	Use:   "film",
	Short: "Get film information",
	Long: `This command gets information about a Star Wars film. It also returns
the names of the different characters that are in it.
Example:
	$ swctl film 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		id, ok := parseID(args)
		if !ok {
			cmd.Println("Could not parse an id to retrieve information for.")
			return
		}

		film, err := swapiService.Film(id)
		if err != nil {
			cmd.Println("Could not find a film for the ID provided.")
			return
		}

		ids := extractIDs(film.CharacterURLs)
		names := collectNames(ids, func(id int) string {
			p, err := swapiService.Person(id)
			if err != nil {
				return ""
			}
			return p.Name
		})

		fmt.Println(formatFilm(film, names))
	},
}

func init() {
	RootCmd.AddCommand(filmCmd)
}

package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "Get character information",
	Long: `This command gets information about a Star Wars character. It also returns
the names of the different films this character appears in.
Example:
	$ swctl character 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		id, ok := parseID(args)
		if !ok {
			cmd.Println("Could not parse an id to retrieve information for.")
			return
		}

		character, err := swapiService.Person(id)
		if err != nil {
			cmd.Println("Could not find a character for the ID provided.")
			return
		}

		ids := extractIDs(character.FilmURLs)
		names := collectNames(ids, func(id int) string {
			f, err := swapiService.Film(id)
			if err != nil {
				return ""
			}
			return f.Title
		})

		fmt.Println(formatPerson(character, names))
	},
}

func init() {
	RootCmd.AddCommand(characterCmd)
}

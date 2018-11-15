package command

import (
	"fmt"
	"os"

	"github.com/codyoss/wired/pkg/client"
	"github.com/codyoss/wired/pkg/swapi"
	"github.com/spf13/cobra"
)

// SWAPIService is an abstraction for a swapi.Service
type SWAPIService interface {
	Film(id int) (*swapi.Film, error)
	Person(id int) (*swapi.Person, error)
}

var swapiService SWAPIService

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "swctl",
	Short: "Star Wars CLI",
	Long:  `A command line utility for enteracting with the Star Wars API`,
}

func init() {
	c := client.New()
	s := swapi.NewService(c)
	swapiService = swapi.NewCachedService(s)
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

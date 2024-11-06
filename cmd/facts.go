package cmd

import (
	"log"

	"github.com/k3v-d3v/catfactninja-cli/catfacts"
	"github.com/spf13/cobra"
)

func NewCatFactsCmd(c *catfacts.CatFactApi) *cobra.Command {
	factCmd := &cobra.Command{
		Use:   "facts",
		Short: "Get a list of cat facts",
		Run: func(cmd *cobra.Command, args []string) {
			r := c.GetFacts()
			for i, f := range r {
				log.Printf("%d: %s\n", i+1, f.Fact)
			}
		},
	}

	return factCmd
}

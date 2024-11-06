package cmd

import (
	"log"

	"github.com/k3v-d3v/catfactninja-cli/catfacts"
	"github.com/spf13/cobra"
)

func NewCatFactCmd(c *catfacts.CatFactApi) *cobra.Command {
	factCmd := &cobra.Command{
		Use:   "fact",
		Short: "Get a random cat fact",
		Run: func(cmd *cobra.Command, args []string) {
			r := c.GetFact()
			log.Println(r.Fact)
		},
	}

	return factCmd
}

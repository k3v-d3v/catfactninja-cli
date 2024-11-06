package cmd

import (
	"log"

	"github.com/k3v-d3v/catfactninja-cli/catfacts"
	"github.com/spf13/cobra"
)

func NewBreedsCmd(c *catfacts.CatFactApi) *cobra.Command {
	breedsCmd := &cobra.Command{
		Use:   "breeds",
		Short: "Get a list of cat breeds",
		Run: func(cmd *cobra.Command, args []string) {
			r := c.GetBreeds()
			for i, b := range r {
				log.Printf("%d: %s\n", i+1, b.Breed)
			}
		},
	}
	return breedsCmd
}

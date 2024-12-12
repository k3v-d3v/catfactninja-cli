package cmd

import (
	"os"

	"github.com/k3v-d3v/catfactninja-cli/catfacts"
	"github.com/k3v-d3v/catfactninja-cli/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "catfactninja",
	Short: "Get some cat facts/breeds",
	Long:  `A simple CLI to get some cat facts from the Cat Facts Ninja API`,
}

func Execute() {
	cfg := config.NewConfig()
	c := catfacts.NewCatFactApi(cfg.BaseUrl)

	rootCmd.AddCommand(NewCatFactCmd(c))
	rootCmd.AddCommand(NewCatFactsCmd(c))
	rootCmd.AddCommand(NewBreedsCmd(c))
	rootCmd.AddCommand(NewHealthCheck())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

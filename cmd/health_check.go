package cmd

import (
	"log"
	"net/url"
	"os"

	"github.com/k3v-d3v/catfactninja-cli/health_check"
	"github.com/spf13/cobra"
)

func NewHealthCheck() *cobra.Command {
	return &cobra.Command{
		Use:   "health [URL, ...]",
		Short: "Check the health of a URL(s)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatal("Please provide a URL to check.")
			}

			urls := make([]string, len(args))

			var isInvalid = false

			// Perform URL Validation
			for _, u := range args {
				_, err := url.ParseRequestURI(u)
				if err != nil {
					log.Println("Invalid URL:", u)
					isInvalid = true
				}
			}

			if isInvalid {
				os.Exit(1)
			}

			copy(urls, args)
			hc := health_check.NewHealthCheck(urls)
			hc.WithChannel().Execute()
		},
	}
}

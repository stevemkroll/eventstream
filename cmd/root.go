package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = new(cobra.Command)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

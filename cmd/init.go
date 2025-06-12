package cmd

import (
	"fmt"
	"os"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed .version
var VERSION string

var rootCmd = &cobra.Command{
	Use:   "coda",
	Args:  cobra.ExactArgs(1),
	Short: "coda-cli is a wrapper for the coda workflow engine",
	Long:  "coda-cli is a wrapper for the coda workflow engine",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute(args []string) {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "help",
		Hidden: true,
	})

	rootCmd.ParseFlags(args)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

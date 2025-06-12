package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/yosev/coda"
)

var versionCmd = &cobra.Command{
	Use:                   "version",
	DisableFlagsInUseLine: true,
	Short:                 "coda cli and engine version",
	Run:                   version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("cli: %s (runtime: %s|%s) | engine: %s\n", VERSION, runtime.GOOS, runtime.GOARCH, coda.VERSION)
}

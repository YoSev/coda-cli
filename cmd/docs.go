package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yosev/coda"
)

var docsCmd = &cobra.Command{
	Use:                   "docs",
	DisableFlagsInUseLine: true,
	Short:                 "coda operations documentation",
	Run:                   docs,
}

var amount bool

func init() {
	docsCmd.Flags().BoolVarP(&amount, "amount", "a", false, "show amount of operations in each category")

	rootCmd.AddCommand(docsCmd)
}

func docs(cmd *cobra.Command, args []string) {
	c := coda.New()
	if amount {
		fmt.Printf("Total: %d\n", len(c.Fn.GetFns()))
		// print amount per category by looping over GetFns()
		categories := make(map[string]int)
		for _, operation := range c.Fn.GetFns() {
			categories[string(operation.Category)]++
		}
		for category, count := range categories {
			fmt.Printf("%s: %d\n", category, count)
		}
	} else {
		for key, operation := range c.Fn.GetFns() {
			fmt.Printf("Action: %s\nName: %s\nDescription: %s\nCategory: %s\n", key, operation.Name, operation.Description, operation.Category)
			if len(operation.Parameters) > 0 {
				fmt.Printf("Parameters:\n")
			}
			for _, v := range operation.Parameters {
				name := v.Name
				if v.Mandatory {
					name = "(*)" + name
				}
				fmt.Printf("  - %s: %s\n", name, v.Description)
			}
			fmt.Printf("\n")
		}
	}
}

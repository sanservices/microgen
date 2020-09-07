package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// text colors
const (
	CEND   = "\033[0m"
	CRED   = "\033[91m"
	CGREEN = "\033[32m"
	CBOLD  = "\033[1m"
)

// rootCmd is the root command,
// any top level command must be added to this one.
var rootCmd = &cobra.Command{
	Use:     "microgen",
	Short:   "Golang micro-service generator.",
	Long:    ``,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts cli commands using
// cobra's Execute function.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

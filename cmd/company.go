package cmd

import (
	"github.com/spf13/cobra"
)

// companyCmd represents the company command
var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Company command for action register or inactive",
}

func init() {
	rootCmd.AddCommand(companyCmd)
}

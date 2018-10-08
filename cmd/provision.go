package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// provisionCmd represents the provision command
var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision a new VM image",
	Long: `Provision a new VM image based on A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		fmt.Printf("provision called for %s", name)
		return nil
	},
}

func init() {
	provisionCmd.Flags().StringP("name", "n", "", "VM image name")
	rootCmd.AddCommand(provisionCmd)
}

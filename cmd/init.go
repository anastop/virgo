package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create sample configuration files and scripts",
	Long:  `Create sample configuration files and scripts.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile := "virgo.json"

		if _, err := os.Stat(configFile); err == nil {
			return fmt.Errorf("file %s already exists", configFile)
		}

		err := ioutil.WriteFile(configFile, []byte(sampleConfig), 0644)
		if err != nil {
			return fmt.Errorf("failed to create %s: %v", configFile, err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

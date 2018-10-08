package cmd

import (
	"fmt"
	"github.com/anastop/virgo/pkg/virgo"
	"log"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Create a new VM instance from an already existing specification",
	Long: `Create a new VM instance from an already existing specification.
This implies that the VM should have been already launched at least once in the past, 
either via 'provision' or 'launch'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		guest, err := cmd.Flags().GetString("guest")
		if err != nil {
			return fmt.Errorf("failed to parse 'guest' argument: %v", err)
		}

		l, err := virgo.NewLibvirtConn()
		if err != nil {
			return fmt.Errorf("failed to open Libvirt connection: %v", err)
		}
		defer func() {
			if err := l.Disconnect(); err != nil {
				log.Fatalf("failed to disconnect from Libvirt: %v", err)
			}
		}()

		if err := virgo.Start(l, guest); err != nil {
			return fmt.Errorf("failed to start guest %s: %v", guest, err)
		}
		return nil
	},
}

func init() {
	startCmd.Flags().StringP("guest", "g", "", "guest to start")
	rootCmd.AddCommand(startCmd)
}

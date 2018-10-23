package cmd

import (
	"fmt"
	"github.com/anastop/virgo/pkg/virgo"
	"log"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Shut down a running VM instance",
	Long:  `Shut down a running VM instance. Keep its current definition intact.`,
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

		if err := virgo.Stop(l, guest); err != nil {
			return fmt.Errorf("failed to stop guest %s: %v", guest, err)
		}
		return nil
	},
}

func init() {
	stopCmd.Flags().StringP("guest", "g", "", "guest to stop")
	stopCmd.MarkFlagRequired("guest")
	rootCmd.AddCommand(stopCmd)
}

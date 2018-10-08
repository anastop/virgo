package cmd

import (
	"fmt"
	"github.com/anastop/virgo/pkg/virgo"
	"log"

	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Fully destroy a VM by undefining it and removing its image",
	Long:  `Fully destroy a domain by undefining it and removing its image`,
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

		if err := virgo.Undefine(l, guest); err != nil {
			return fmt.Errorf("failed to undefine %s: %v", guest, err)
		}

		if err := virgo.Purge(l, guest); err != nil {
			return fmt.Errorf("failed to purge %s: %v", guest, err)
		}
		return nil
	},
}

func init() {
	purgeCmd.Flags().StringP("guest", "g", "", "guest to purge")
	rootCmd.AddCommand(purgeCmd)
}

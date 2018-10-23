package cmd

import (
	"fmt"
	"github.com/anastop/virgo/pkg/virgo"
	"log"

	"github.com/spf13/cobra"
)

// undefineCmd represents the undefine command
var undefineCmd = &cobra.Command{
	Use:   "undefine",
	Short: "Undefine a VM by removing its specification",
	Long: `Undefine a VM by removing its specification. Its image is not affected.
If it's running, the domain is first stopped.'`,
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
		return nil
	},
}

func init() {
	undefineCmd.Flags().StringP("guest", "g", "", "guest to undefine")
	undefineCmd.MarkFlagRequired("guest")
	rootCmd.AddCommand(undefineCmd)
}

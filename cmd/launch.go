package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/anastop/virgo/pkg/virgo"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Define and start a new VM instance",
	Long: `Define and start a new VM instance based on user-provided launch options.
The VM's image should have been already provisioned using the 'provision' command.
Any previous specification of the VM is overriden by the new launch options. 

The available launch options are presented in detail in virgo's main help message.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		guest, err := cmd.Flags().GetString("guest")
		if err != nil {
			return fmt.Errorf("failed to parse 'guest' argument: %v", err)
		}

		conf, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to parse config argument: %v", err)
		}

		data, err := ioutil.ReadFile(conf)
		if err != nil {
			return fmt.Errorf("failed to read config file %s: %v", conf, err)
		}

		gc := &virgo.GuestConf{}
		if err := json.Unmarshal(data, gc); err != nil {
			return fmt.Errorf("failed to unmarshal guest config: %v", err)
		}
		gc.Name = guest

		l, err := virgo.NewLibvirtConn()
		if err != nil {
			return fmt.Errorf("failed to open Libvirt connection: %v", err)
		}
		defer func() {
			if err := l.Disconnect(); err != nil {
				log.Fatalf("failed to disconnect from Libvirt: %v", err)
			}
		}()

		gc.RootImgPath, gc.ConfigIsoPath, err = virgo.GuestImagePaths(l, virgo.DefaultPool(), guest)
		if err != nil {
			return fmt.Errorf("failed to compute image paths for %s: %v", guest, err)
		}

		if err := virgo.LaunchGuest(l, gc); err != nil {
			return fmt.Errorf("launch failed: %v", err)
		}

		return nil
	},
}

func init() {
	launchCmd.Flags().StringP("guest", "g", "", "guest to launch")
	launchCmd.Flags().StringP("config", "c", "", "JSON file containing the launch options")
	rootCmd.AddCommand(launchCmd)
}

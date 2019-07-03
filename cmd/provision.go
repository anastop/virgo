package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/anastop/virgo/pkg/virgo"

	"github.com/spf13/cobra"
)

var provisionCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision a new VM image",
	Long: `Provision a new VM image based on a user-provided provision bash script and provisioning options.

The available provisioning options are presented in detail in virgo's main help message. 
The bash script can be any valid bash script and is executed with root permissions. 
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		guest := args[0]

		provisionScript, err := cmd.Flags().GetString("provision-script")
		if err != nil {
			return fmt.Errorf("failed to parse provision argument: %v", err)
		}

		initdScript, err := cmd.Flags().GetString("initd-script")
		if err != nil {
			return fmt.Errorf("failed to parse initd argument: %v", err)
		}

		conf, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to parse config argument: %v", err)
		}

		data, err := ioutil.ReadFile(conf)
		if err != nil {
			return fmt.Errorf("failed to read config file %s: %v", conf, err)
		}

		pc := &virgo.ProvisionConf{}
		if err := json.Unmarshal(data, pc); err != nil {
			return fmt.Errorf("failed to unmarshal provision config: %v", err)
		}
		pc.Name = guest

		gc := &virgo.GuestConf{}
		if err := json.Unmarshal(data, gc); err != nil {
			return fmt.Errorf("failed to unmarshal guest config: %v", err)
		}
		gc.Name = guest

		if provisionScript != "" {
			data, err = ioutil.ReadFile(provisionScript)
			if err != nil {
				return fmt.Errorf("failed to read provision script %s: %v", provisionScript, err)
			}
			pc.Provision = string(data)
		}

		if initdScript != "" {
			data, err = ioutil.ReadFile(initdScript)
			if err != nil {
				return fmt.Errorf("failed to read initd script %s: %v", initdScript, err)
			}
			pc.Initd = string(data)
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

		if err := virgo.Provision(l, pc, gc); err != nil {
			return fmt.Errorf("provision failed: %v", err)
		}

		return nil
	},
}

func init() {
	provisionCmd.Flags().StringP("provision-script", "p", "", "bash script to be used for provisioning")
	provisionCmd.Flags().StringP("initd-script", "i", "", "bash script to be used in init.d")
	provisionCmd.Flags().StringP("config", "c", "", "JSON file containing the provisioning options")
	provisionCmd.MarkFlagRequired("config")
	rootCmd.AddCommand(provisionCmd)
}

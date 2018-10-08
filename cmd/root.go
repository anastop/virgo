package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "virgo",
	Short: "virgo enables easy provisioning, configuration and management of Libvirt guests",
	Long: `virgo enables easy provisioning, configuration and management of Libvirt guests. 

For provisioning a new VM image, you should specify a JSON config file with provisioning
options, along with a provisioning script to be executed on image's first boot.

For launching a new VM instance from an already-provisioned image, you should specify a 
JSON config file with launch options. 

The example below shows all available provisioning and launch options, all in a single
JSON file (ignore the '#' lines which serve as comments). 

{
  # PROVISIONING OPTIONS
  "cloud_img_url": "https://cloud-images.ubuntu.com/releases/16.04/release/",
  "cloud_img_name": "ubuntu-16.04-server-cloudimg-amd64-disk1.img",
  "user": "nfvsap",
  "passwd": "nfvsap",
  "root_img_gb": 10,

  # LAUNCH OPTIONS
  "guest_memory_mb": 4096,
  "guest_num_vcpus": 8,
  "guest_hugepage_support": true,
  "guest_hugepage_size": 2,
  "guest_hugepage_size_unit": "M",
  "guest_hugepage_node_set": "0",
  "guest_net_ifs": [
    {"type": "bridge", "bridge": "virbr0"},
    {"type": "bridge", "bridge": "virbr0"},
    {
      "type": "vhostuser",
      "mac_addr": "de:ad:be:ef:01:23",
      "unix_socket_path": "/usr/local/var/run/openvswitch/dpdkvhostuser1",
      "queues": 2
    },
    {
      "type": "vhostuser",
      "mac_addr": "de:ad:be:ef:45:67",
      "unix_socket_path": "/usr/local/var/run/openvswitch/dpdkvhostuser2",
      "queues": 2
    }]
}

The provisioning script can be any valid bash script, and it's executed as the 
last step of cloud-init provisioning. 

PREREQUISITES
The following Linux utilities are required by virgo: 
- wget
- genisoimage
`}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

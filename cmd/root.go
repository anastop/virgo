package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sampleConfig = `
{
  "cloud_img_url": "https://cloud-images.ubuntu.com/releases/18.04/release/",
  "cloud_img_name": "ubuntu-18.04-server-cloudimg-amd64.img",
  "user": "guest",
  "passwd": "guest",
  "root_img_gb": 10,

  "guest_memory_mb": 4096,
  "guest_num_vcpus": 8,
  "guest_num_sockets": 2,
  "guest_num_cores_per_socket": 2,
  "guest_num_threads_per_core": 2,
  "guest_numa_nodes": [ 
    {"id": 0, "cpus": "0-3", "memory_mb": 2048 },
    {"id": 1, "cpus": "4-7", "memory_mb": 2048 }
  ],
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
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "virgo",
	Short: "virgo enables easy provisioning, configuration and management of Libvirt guests",
	Long: `virgo enables easy provisioning, configuration and management of Libvirt guests. 

Most virgo commands accept a single argument, the name of the VM they act upon. Every command
has its own flags. 

For provisioning a new VM image, you should specify a JSON config file with provisioning
options. Additionally, you may specify a provisioning script to be executed on image's first boot,
and/or an initd script with commands to be executed on every boot. 

For launching a new VM instance from an already-provisioned image, you should specify a 
JSON config file with launch options. 

The example below shows all available provisioning and launch options, all in a single JSON file
(these groups of options are separated by an empty line).

` + sampleConfig + `

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

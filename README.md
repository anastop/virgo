# virgo

virgo allows you to quickly provision and spin VMs locally, leveraging cloud-init and
Libvirt.

![virgo](./virgo.png)

## Features

- Runs on Linux baremetal machines and leverages Libvirt
- Allows easy VM provisioning based on user-provided provisioning scripts and simple configuration options (uses [cloud-init](https://cloudinit.readthedocs.io/en/latest/) under the hood)
- Allows easy VM creation with flexible configuration options
- Supports [vhost-user network interfaces](https://libvirt.org/formatdomain.html#elementVhostuser), to allow a VM to connect e.g. with a  DPDK-based vswitch

Provisioning options:
- cloud image used for provisioning (currently tested with Ubuntu 16.04)
- user credentials
- custom provisioning script to be used during VM creation
- custom init.d script to be installed permanently

VM configuration options: 
- number of vCPUs
- guest memory
- hugepage backing options
- network interfaces: support for `bridge`-type and `vhostuser`-type interfaces

## Installation

You can build virgo from source:
```console
$ go get github.com/anastop/virgo
$ go install
```

Or download the latest binary from the "Releases" page.

### Dependencies

virgo makes use of the following utilities: 
- wget
- genisoimage

## Usage 

Provision a new VM called "foo":
```console
$ sudo virgo provision --config guest_config.json --provision-script provision.sh --initd-script initd.sh --guest foo
```
"foo" will shutdown after provisioning. 

Edit `guest_config.json` to change VM's parameters (e.g. #vCPUs), and launch "foo":
```console
$ sudo virgo --config guest_config.json --guest foo
```

Usage:
```console
$ virgo --help

virgo enables easy provisioning, configuration and management of Libvirt guests.

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

Usage:
  virgo [command]

Available Commands:
  help        Help about any command
  launch      Define and start a new VM instance
  provision   Provision a new VM image
  purge       Fully destroy a VM by undefining it and removing its image
  start       Create a new VM instance from an already existing specification
  stop        Shut down a running VM instance
  undefine    Undefine a VM by removing its specification

Flags:
  -h, --help   help for virgo

Use "virgo [command] --help" for more information about a command.
```



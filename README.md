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
- cloud image used for provisioning (currently tested with Ubuntu 16.04 & 18.04)
- user credentials
- custom provisioning script to be used during VM creation
- custom init.d script to be installed permanently

VM configuration options: 
- number and topology of vCPUs
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

```console 
$ virgo init
```

Edit the created `virgo.json` according to your needs.

Optionally, create additional `provision.sh` and `initd.sh` files to be used as provisioning 
and initd scripts, respectively.

Provision a new VM called "foo":

```console
$ sudo virgo provision foo --config virgo.json [--provision-script provision.sh] [--initd-script initd.sh]
```

"foo" will shutdown after provisioning. 

Edit `virgo.json` to change VM's parameters (e.g. #vCPUs), and launch "foo":

```console
$ sudo virgo launch foo --config virgo.json
```

To find out more, run `virgo -h`. 

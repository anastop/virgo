# virgo

Flexible provisioning and management of virtual machines 

![virgo](./virgo.png)

## Features

- Runs on Linux baremetal machines and leverages Libvirt
- Allows easy VM provisioning based on user-provided provisioning scripts and simple configuration options (uses [cloud-init](https://cloudinit.readthedocs.io/en/latest/) under the hood)
- Allows easy VM creation with flexible configuration options
- Supports [vhost-user network interfaces](https://libvirt.org/formatdomain.html#elementVhostuser), to allow a VM to connect e.g. with a  DPDK-based vswitch

Provisioning options: 
- cloud image used for provisioning (currently tested with Ubuntu 16.04)
- user credentials

VM configuration options: 
- number of vCPUs
- guest memory
- hugepage backing options
- network interfaces: support for `bridge`-type and `vhostuser`-type interfaes

## Installation

```console
$ go get github.com/anastop/virgo
$ go install
```

## Usage 

```console
$ virgo --help
```



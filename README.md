# virgo

Flexible provisioning and management of virtual machines 

Features:
- runs on Linux baremetal machines, using Libvirt under the hood
- allows easy VM provisioning, based on user-provided Bash script and simple JSON configuration (uses [cloud-init](https://cloudinit.readthedocs.io/en/latest/) under the hood)
- allows easy VM creation with rich configuration options
- supports [vhost-user network interfaces](https://libvirt.org/formatdomain.html#elementVhostuser), to allow a VM to be connected to e.g. a DPDK-based switch

VM provisioning options: 
- `cloud_img_url`, `cloud_img_name`: cloud image to use for provisioning
- `user`, `passwd`: user credentials to provision within the VM
- `root_img_gb`: provisioned image size

VM launching options: 
- `guest_memory_mb`: guest memory size
- `guest_num_vcpus`: guest number of vCPUs
- `guest_hugepage_support`, `guest_hugepage_size`, `guest_hugepage_size_unit`, `guest_hugepage_node_set`: hugepage backing options
- `guest_net_ifs`: network interfaces configuration
    - "bridge" interfaces (config options: `bridge`)
    - "vhostuser" interfaces (config options: `mac_addr`, `unix_socket_path`, `queues`)



#!/usr/bin/env bash

apt-get install -y gcc make libvirt-bin python python-pip
mkdir -p /opt/nfv
cd /opt/nfv

# install DPDK
wget http://fast.dpdk.org/rel/dpdk-16.07.tar.xz
tar -xvf dpdk-16.07.tar.xz
export RTE_SDK=$(pwd)/dpdk-16.07
export RTE_TARGET=x86_64-native-linuxapp-gcc
cd $RTE_SDK
make config T=$RTE_TARGET
make install T=$RTE_TARGET DESTDIR=dpdk-install

# install T-Rex
mkdir -p /opt/nfv/trex
cd /opt/nfv/trex
wget --no-cache http://trex-tgn.cisco.com/trex/release/v2.45.tar.gz
tar zxvf v2.45.tar.gz

# install scapy
pip install scapy

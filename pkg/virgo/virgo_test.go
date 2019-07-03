package virgo

import (
	"testing"
)

func TestDomXMLStr(t *testing.T) {
	s := &GuestConf{
		Name:              "foo",
		MemoryMB:          8192,
		NumVcpus:          8,
		NumSockets:        2,
		NumCoresPerSocket: 2,
		NumThreadsPerCore: 2,
		NUMANodes: []NUMANode{
			{Id: 0, Cpus: "0-3", MemoryMB: 4096},
			{Id: 1, Cpus: "4-7", MemoryMB: 4096},
		},
		HugepageSupport:  true,
		HugepageSize:     2,
		HugepageSizeUnit: "MB",
		HugepageNodeSet:  "0,1",
		RootImgPath:      "foo.img",
		ConfigIsoPath:    "foo.iso",
		NetIfs: []NetIf{
			{Type: "bridge", Bridge: "virbr0"},
			{Type: "bridge", Bridge: "virbr0"},
			{
				Type:           "vhostuser",
				MacAddr:        "de:ad:be:ef:01:23",
				UnixSocketPath: "/usr/local/var/run/openvswitch/dpdkvhostuser1",
				Queues:         2,
			},
			{
				Type:           "vhostuser",
				MacAddr:        "de:ad:be:ef:45:67",
				UnixSocketPath: "/usr/local/var/run/openvswitch/dpdkvhostuser2",
				Queues:         2,
			},
		},
	}

	xml, err := domXML(s)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Domain XML string: %s", xml)
}

func TestCreateUserDataFile(t *testing.T) {
	p := &ProvisionConf{Name: "test",
		CloudImgURL:  "https://cloud-images.ubuntu.com/releases/16.04/release/",
		CloudImgName: "ubuntu-16.04-server-cloudimg-amd64-disk1.img",
		User:         "nfvsap",
		Passwd:       "nfvsap",
	}

	p.Provision = `#/bin/bash
echo Hello
echo Hello again`

	p.Initd = p.Provision

	if err := createUserDataFile("user-data", p); err != nil {
		t.Fatal(err)
	}
}

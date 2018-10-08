package virgo

import (
	"testing"
)

func TestDomXMLStr(t *testing.T) {
	s := &GuestConf{
		MemoryMB:         2,
		NumVcpus:         2,
		HugepageSupport:  true,
		HugepageSize:     2,
		HugepageSizeUnit: "MB",
		HugepageNodeSet:  "0",
		rootImgPath:      "foo.img",
		configIsoPath:    "foo.iso",
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

	t.Logf("Domain XML string: %s", domXMLStr("foo", s))
}

func TestCreateUserDataFile(t *testing.T) {
	if err := createUserDataFile("user-data", "nfvsap", "nfvsap", "echo hello\nls *"); err != nil {
		t.Fatal(err)
	}
}

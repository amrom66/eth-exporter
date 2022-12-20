package pkg

import (
	"time"

	"github.com/google/gopacket/pcap"
)

var (
	Device      string
	snapshotLen int32         = 1024
	promiscuous bool          = false
	timeout     time.Duration = 15 * time.Second
	handle      *pcap.Handle

	Token    string
	Url      string
	Bucket   string
	Org      string
	Instance string
)

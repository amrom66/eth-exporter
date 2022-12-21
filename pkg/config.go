package pkg

import (
	"time"

	"github.com/google/gopacket/pcap"
)

var (
	Device      string
	snapshotLen int32         = 1024
	promiscuous bool          = false
	timeout     time.Duration = 1 * time.Second
	handle      *pcap.Handle
	err         error

	Token    string
	Url      string
	Bucket   string
	Org      string
	Instance string
)

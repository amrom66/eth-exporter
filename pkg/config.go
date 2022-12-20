package pkg

import (
	"time"

	"github.com/google/gopacket/pcap"
)

var (
	device      string = "eth0"
	snapshotLen int32  = 1024
	promiscuous bool   = false
	err         error
	timeout     time.Duration = 15 * time.Second
	handle      *pcap.Handle

	token    string = "kb5IYxUwMSkETAFDz6YRj_FYdiFB9Tec_AiKej1h1s7Q9x_g6ol74YwUoAqUHZQ8UXa1UnCaLGCK3DoYM6PN1w=="
	url      string = "http://10.4.2.135:8086"
	bucket   string = "amrom"
	org      string = "amrom"
	instance string = "10.4.2.135"
)

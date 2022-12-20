package pkg

import (
	"log"
	"time"

	"github.com/golang/glog"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Capture() {
	handle, err := pcap.OpenLive(Device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	glog.Infoln("start capture", time.Now())
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			var mm = mypacket{
				ip.SrcIP.String(),
				ip.DstIP.String(),
				ip.Protocol.String(),
			}
			flush2influxdb(Url, Token, Org, Bucket, mm)
		}
	}
	glog.Infoln("end capture", time.Now())
}

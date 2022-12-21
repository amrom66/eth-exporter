package pkg

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func Capture() {
	ch := make(chan gopacket.Packet, 1000)

	Init()

	handle, err = pcap.OpenLive(Device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	go func() {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			ch <- packet
		}
	}()

	// flush data to influxdb
	go func() {
		for {
			packet, ok := <-ch
			if !ok {
				fmt.Println("channel has been clonsed.")
				break
			}
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)
				var mm = mypacket{
					ip.SrcIP.String(),
					ip.DstIP.String(),
					ip.Protocol.String(),
				}
				flush2influxdb(mm)
			}
		}
	}()

	select {}
}

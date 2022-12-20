package pkg

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client

func flush2influxdb(url string, token string, org string, bucket string, mypomypacket mypacket) {
	client = influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPI(org, bucket)
	p := influxdb2.NewPointWithMeasurement("mypacket").
		AddTag("instance", instance).
		AddField("src", mypomypacket.src).
		AddField("dst", mypomypacket.dst).
		AddField("protocol", mypomypacket.protocol).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
}

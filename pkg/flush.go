package pkg

import (
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client

func flush2influxdb(mm mypacket) {
	writeAPI := client.WriteAPI(Org, Bucket)
	p := influxdb2.NewPointWithMeasurement("mypacket").
		AddTag("instance", Instance).
		AddField("src", mm.src).
		AddField("dst", mm.dst).
		AddField("protocol", mm.protocol).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
}

func Init() {
	client = influxdb2.NewClient(Url, Token)
	client.Options().SetRetryInterval(1000)
}

# eth-exporter
collect src ip and dst ip.


## build

```shell

sudo apt-get install libpcap-dev

go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s"

```
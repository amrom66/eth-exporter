# eth-exporter
collect src ip and dst ip.


## build

```shell

go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s"

```
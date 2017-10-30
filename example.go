package main

import (
	"github.com/lflxp/sflowtool/collected"
	"time"
	"flag"
)

var Con collected.Collected = collected.Collected{
	DeviceName:  "en0",
	SnapShotLen: 1024,
	Promiscuous: true,
	Timeout:     30 * time.Second,
}

func main() {
	wait := make(chan int)
	item := flag.String("t","all","类型:all(sflowSample|Counter),counter(SflowCounter),sample(SflowSample),netflow")
	protocol := flag.String("s","udp","协议")
	port := flag.String("p","6343","端口")
	eth := flag.String("e","en0","网卡名")
	flag.Parse()

	Con.DeviceName = *eth

	if *item == "all" {
		SflowAll(*protocol,*port)
	} else if *item == "counter" {
		SflowCounter(*protocol,*port)
	} else if *item == "sample" {
		SflowSample(*protocol,*port)
	} else if *item == "netflow" {
		NetflowV5(*protocol,*port)
	}

	<-wait
}

func SflowCounter(protocol,port string) {
	Con.ListenSflowCounter(protocol,port)
}

func SflowSample(protocol,port string) {
	Con.ListenSFlowSample(protocol,port)
}

//include SFlowSample and SflowCounter
func SflowAll(protocol,port string) {
	Con.ListenSflowAll(protocol,port)
}

func NetflowV5(protocol,port string) {
	Con.ListenNetFlowV5(protocol,port)
}

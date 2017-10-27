package main

import (
	"github.com/lflxp/sflowtool/collected"
	"time"
)

var Con collected.Collected = collected.Collected{
	DeviceName:  "en0",
	SnapShotLen: 1024,
	Promiscuous: false,
	Timeout:     30 * time.Second,
}

func main() {
	//SflowAll()
	//SflowSample()
	//SflowCounter()
	NetflowV5()
	time.Sleep(60*time.Second)
}

func SflowCounter() {
	Con.ListenSflowCounter("udp","9999")
}

func SflowSample() {
	Con.ListenSFlowSample("udp","9999")
}

//include SFlowSample and SflowCounter
func SflowAll() {
	Con.ListenSflowAll("udp","9999")
}

func NetflowV5() {
	Con.ListenNetFlowV5("udp","6343")
}
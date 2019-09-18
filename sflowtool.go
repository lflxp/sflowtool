package main

import (
	"flag"
	//	"github.com/google/gopacket/afpacket"
	"net"
	"time"

	"github.com/lflxp/sflowtool/collected"
	"github.com/lflxp/sflowtool/parse"
	log "github.com/sirupsen/logrus"
)

var Con collected.Collected = collected.Collected{
	DeviceName:  "en0",
	SnapShotLen: 65535,
	Promiscuous: true,
	Timeout:     30 * time.Second,
}

func main() {
	wait := make(chan int)
	item := flag.String("t", "all", "类型:all(sflowSample|Counter),counter(SflowCounter),sample(SflowSample),netflow")
	protocol := flag.String("s", "udp", "协议")
	port := flag.String("p", "6343", "端口")
	eth := flag.String("ee", "en0", "网卡名")
	udp := flag.Bool("udp", false, "是否开启udp数据传输,默认不开启")
	udport := flag.String("host", "127.0.0.1:6666", "udp SFlowSample And Netflow 传输主机:端口")
	counterport := flag.String("chost", "127.0.0.1:7777", "udp CounterSample 传输主机:端口")
	esurl := flag.String("es", "http://127.0.0.1:9200", "elasticsearch 5.6 接口地址")
	ises := flag.Bool("ises", false, "是否开启output到elasticsearch")
	debug := flag.Bool("debug", false, "是否开启debug model")
	index := flag.String("index", "sflow", "es index name, example: sflow-2019-09-06")
	flag.Parse()

	Con.DeviceName = *eth
	Con.Host = *udport
	Con.Udpbool = *udp
	Con.CounterHost = *counterport
	Con.EsPath = *esurl
	Con.IsEs = *ises
	Con.Index = *index

	// 初始化es index
	if Con.IsEs {
		log.Infof("开启es通道 EsPath: %s Index: %s\n", Con.EsPath, Con.Index)
		parse.InitEs(Con.EsPath, Con.Index)
	}

	// 是否开启udp数据转发
	if *udp {
		Conn, err := net.Dial("udp", *udport)
		defer Conn.Close()
		if err != nil {
			panic(err)
		}
	}

	// 设置日志级别
	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Info("日志模式 DEBUG")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("日志模式 INFO")
	}

	// 启动命令
	if *item == "all" {
		SflowAll(*protocol, *port)
	} else if *item == "counter" {
		SflowCounter(*protocol, *port)
	} else if *item == "sample" {
		SflowSample(*protocol, *port)
	} else if *item == "netflow" {
		NetflowV5(*protocol, *port)
	}

	<-wait
}

func SflowCounter(protocol, port string) {
	Con.ListenSflowCounter(protocol, port)
}

func SflowSample(protocol, port string) {
	Con.ListenSFlowSample(protocol, port)
}

//include SFlowSample and SflowCounter
func SflowAll(protocol, port string) {
	Con.ListenSflowAll(protocol, port)
}

func NetflowV5(protocol, port string) {
	Con.ListenNetFlowV5(protocol, port)
}

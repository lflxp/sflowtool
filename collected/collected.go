package collected

import (
	"time"
	"github.com/google/gopacket/pcap"
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	. "github.com/lflxp/sflowtool/sflowV5"
	. "github.com/lflxp/sflowtool/netflowV5"
	"net"
)

type Collected struct {
	DeviceName      string //设备名称
	SnapShotLen     int32
	SnapShotLenUint uint32
	Promiscuous     bool //是否开启混杂模式
	Timeout         time.Duration
	Udpbool 	bool   //是否开启udp传输
	Host 		string //udp 发送客户端及端口 127.0.0.1:8888
}

func (this *Collected) SendUdp(result string) {
	conn,err := net.Dial("udp",this.Host)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	conn.Write([]byte(result))
}

func (this *Collected) ListenSFlowSample(protocol, port string) {
	//Open Device
	handle, err := pcap.OpenLive(this.DeviceName, this.SnapShotLen, this.Promiscuous, this.Timeout)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	defer handle.Close()

	//Set filter
	var filter string = fmt.Sprintf("%s and port %s", protocol, port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		beego.Error(err)
	}
	beego.Informational(fmt.Sprintf("Only capturing %s port %s packets.", protocol, port))
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		p := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Default)
		if p.ErrorLayer() != nil {
			fmt.Println("failed :", p.ErrorLayer().Error())
		}
		udpLayer := p.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//fmt.Println("UDP layer detected.")
			udp, _ := udpLayer.(*layers.UDP)
			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
			if pp.ErrorLayer() != nil {
				fmt.Println("failed :", pp.ErrorLayer().Error(),udp.Payload)
			}
			if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				go func(datas []layers.SFlowFlowSample){
					for _, y := range datas {
					//beego.Critical(len(y.Records),y.RecordCount)
					tmp := NewFlowSamples()
					tmp.InitOriginData(p)
					tmp.InitFlowSampleData(y)
					for _, yy := range y.Records {
						if g1, ok1 := yy.(layers.SFlowRawPacketFlowRecord); ok1 {
							tmp.ParseLayers(g1.Header)

							b, err := json.Marshal(tmp)
							if err != nil {
								fmt.Println(err.Error())
							}
							if this.Udpbool {
								this.SendUdp(string(b))
							} else {
								fmt.Println(string(b))
							}
						}
					}
				}
				}(got.FlowSamples)
			}
		}
	}
}

func (this *Collected) ListenSflowCounter(protocol, port string) {
	//Open Device
	handle, err := pcap.OpenLive(this.DeviceName, this.SnapShotLen, this.Promiscuous, this.Timeout)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	defer handle.Close()

	//Set filter
	var filter string = fmt.Sprintf("%s and port %s", protocol, port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		beego.Error(err)
	}
	beego.Informational(fmt.Sprintf("Only capturing %s port %s packets.", protocol, port))
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		p := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Default)
		if p.ErrorLayer() != nil {
			fmt.Println("failed :", p.ErrorLayer().Error())
		}

		//ipLayer := packet.Layer(layers.LayerTypeIPv4)
		//if ipLayer != nil {
		//	ip, _ := ipLayer.(*layers.IPv4)
		//	fmt.Println(ip.SrcIP.String(),ip.DstIP.String())
		//	//fmt.Println(p.Dump())
		//}

		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//fmt.Println("UDP layer detected.")
			udp, _ := udpLayer.(*layers.UDP)

			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
			if pp.ErrorLayer() != nil {
				fmt.Println("failed :", pp.ErrorLayer().Error(),udp.Payload)
			}
			if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				beego.Error(udp.Payload)
				go func(datas []layers.SFlowCounterSample){
					if len(datas) > 0 {
						tmp := NewCounterFlow()
						tmp.InitOriginData(p)
						for _, y := range datas {
							//beego.Critical(len(y.Records),y.RecordCount)
							tmp.InitCounterSample(y)
						}

						b, err := json.Marshal(tmp)
						if err != nil {
							fmt.Println(err.Error())
						}

						if this.Udpbool {
							this.SendUdp(string(b))
						} else {
							fmt.Println(string(b))
						}
					}
				}(got.CounterSamples)
			}
		}

		sflow := packet.Layer(layers.LayerTypeSFlow)
		if sflow != nil {
			fmt.Println("SFLOW layer detected")
		}
	}
}

func (this *Collected) ListenSflowAll(protocol, port string) {
	//Open Device
	handle, err := pcap.OpenLive(this.DeviceName, this.SnapShotLen, this.Promiscuous, this.Timeout)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	defer handle.Close()

	//Set filter
	var filter string = fmt.Sprintf("%s and port %s", protocol, port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		beego.Error(err)
	}
	beego.Informational(fmt.Sprintf("Only capturing %s port %s packets.", protocol, port))
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		p := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Default)
		if p.ErrorLayer() != nil {
			fmt.Println("failed :", p.ErrorLayer().Error())
		}
		udpLayer := p.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//fmt.Println("UDP layer detected.")
			udp, _ := udpLayer.(*layers.UDP)
			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
			if pp.ErrorLayer() != nil {
				fmt.Println("failed :", pp.ErrorLayer().Error())
			} else if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				go func(Sample []layers.SFlowFlowSample,Counter []layers.SFlowCounterSample) {
					if len(Sample) > 0 {
						for _, y := range Sample {
							//beego.Critical(len(y.Records),y.RecordCount)
							tmp := NewFlowSamples()
							tmp.InitOriginData(p)
							tmp.InitFlowSampleData(y)
							for _, yy := range y.Records {
								if g1, ok1 := yy.(layers.SFlowRawPacketFlowRecord); ok1 {
									tmp.ParseLayers(g1.Header)

									b, err := json.Marshal(tmp)
									if err != nil {
										fmt.Println(err.Error())
									}

									if this.Udpbool {
										this.SendUdp(string(b))
									} else {
										fmt.Println(string(b))
									}
								}
							}
						}
					}

					if len(Counter) > 0 {
						tmp := NewCounterFlow()
						tmp.InitOriginData(p)
						for _, y := range Counter {
							//beego.Critical(len(y.Records),y.RecordCount)
							tmp.InitCounterSample(y)
						}

						b, err := json.Marshal(tmp)
						if err != nil {
							fmt.Println(err.Error())
						}

						if this.Udpbool {
							this.SendUdp(string(b))
						} else {
							fmt.Println(string(b))
						}
					}
				} (got.FlowSamples,got.CounterSamples)
			}
		}
	}
}

func (this *Collected) ListenNetFlowV5(protocol, port string) {
	//Open Device
	handle, err := pcap.OpenLive(this.DeviceName, this.SnapShotLen, this.Promiscuous, this.Timeout)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	defer handle.Close()

	//Set filter
	var filter string = fmt.Sprintf("%s and port %s", protocol, port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		beego.Error(err)
	}
	beego.Informational(fmt.Sprintf("Only capturing %s port %s packets.", protocol, port))
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//beego.Informational("1")
		go func(packet gopacket.Packet) {
			//beego.Informational("2")
			//beego.Error("############开始解析#############")
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer != nil {
				//fmt.Println("UDP layer detected.")
				udp, _ := udpLayer.(*layers.UDP)

				tmp := NetFlowV5{}

				for _,x := range tmp.PayLoadToNetFlowV5(udp.Payload, packet.NetworkLayer().NetworkFlow().Src().String()) {
					this.SendUdp(x)
				}
				//beego.Error(len(data))
				//fmt.Println(data)
			}
		}(packet)
	}
}
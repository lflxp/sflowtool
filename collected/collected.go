package collected

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	. "github.com/lflxp/sflowtool/netflowV5"
	. "github.com/lflxp/sflowtool/sflowV5"
	//"github.com/Cistern/sflow"
	"net"
	"strings"
	"time"
)

type Collected struct {
	DeviceName      string //设备名称
	SnapShotLen     int32
	SnapShotLenUint uint32
	Promiscuous     bool //是否开启混杂模式
	Timeout         time.Duration
	Udpbool         bool   //是否开启udp sample and netflow传输
	Host            string //udp 发送客户端及端口 127.0.0.1:8888
	CounterHost     string //udp counter 传输
}

func (this *Collected) SendUdp(result string, counter bool) {
	if counter {
		conn, err := net.Dial("udp", this.CounterHost)
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(result))
	} else {
		conn, err := net.Dial("udp", this.Host)
		defer conn.Close()
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(result))
	}
}

func (this *Collected) CheckInfo(ppp []byte) {
	p := gopacket.NewPacket(ppp, layers.LayerTypeEthernet, gopacket.Default)
	if p.ErrorLayer() != nil {
		fmt.Println("failed :", p.ErrorLayer().Error())
	}
	eth := p.Layer(layers.LayerTypeEthernet)
	if eth != nil {
		fmt.Println("Ethernet layer detected.")
		ethernetPacket, _ := eth.(*layers.Ethernet)
		fmt.Println("Source MAC:", ethernetPacket.SrcMAC)
		fmt.Println("Destionation MAC:", ethernetPacket.DstMAC)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
		fmt.Println()
	}

	Dq := p.Layer(layers.LayerTypeDot1Q)
	if Dq != nil {
		fmt.Println("LayerTypeDot1Q layer detected.")
		dd, _ := Dq.(*layers.Dot1Q)
		fmt.Println("DropEligible:", dd.DropEligible)
		fmt.Println("Priority:", dd.Priority)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Dot1Q type: ", dd.Type)
		fmt.Println("VLANIdentifier: ", dd.VLANIdentifier)
		fmt.Println()
	}

	icmq := p.Layer(layers.LayerTypeICMPv4)
	if icmq != nil {
		fmt.Println("LayerTypeICMPv4 layer detected.")
		ic, _ := icmq.(*layers.ICMPv4)
		fmt.Println("Checksum:", ic.Checksum)
		fmt.Println("Id:", ic.Id)
		// Ethernet type is typically IPv4 but could be ARP or other
		fmt.Println("Seq: ", ic.Seq)
		fmt.Println("TypeCode: ", ic.TypeCode.String())
		fmt.Println()
	}

	// Let's see if the packet is IP ∂(even though the ether type told us)
	ipLayer := p.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		fmt.Println("Protocol: ", ip.Protocol)
		fmt.Println()
	}

	// Let's see if the packet is TCP
	udpLayer := p.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		fmt.Println("UDP layer detected.")
		udp, _ := udpLayer.(*layers.UDP)

		// TCP layer variables:
		// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
		// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
		fmt.Printf("From port %d to %d\n", udp.SrcPort, udp.DstPort)
		fmt.Println("Checksum number: ", udp.Checksum)
		fmt.Println()

		pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
		if pp.ErrorLayer() == nil {
			fmt.Println("UDP SFLOW detected")
		}
	}
	sflowlayer := p.Layer(layers.LayerTypeSFlow)
	if sflowlayer != nil {
		fmt.Println("SFLOW layer detected")
	}

	// When iterating through packet.Layers() above,
	// if it lists Payload layer then that is the same as
	// this applicationLayer. applicationLayer contains the payload
	applicationLayer := p.ApplicationLayer()
	if applicationLayer != nil {
		fmt.Println("Application layer/Payload found.")
		fmt.Printf("%d %s\n", len(applicationLayer.Payload()), applicationLayer.Payload())
		// Search for a string inside the payload
		if strings.Contains(string(applicationLayer.Payload()), "HTTP") {
			fmt.Println("HTTP found!")
		}
		fmt.Println(applicationLayer.Payload())
		fmt.Println()
	}

	// Check for errors
	if err := p.ErrorLayer(); err != nil {
		fmt.Println("Error decoding some part of the packet:", err)
	}

	for _, x := range p.Layers() {
		fmt.Println(x.LayerType().String())
	}

	fmt.Println()
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
			fmt.Println("failed LayerTypeEthernet:", p.ErrorLayer().Error())
		}
		udpLayer := p.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//fmt.Println("UDP layer detected.")
			udp, _ := udpLayer.(*layers.UDP)
			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
			//if pp.ErrorLayer() != nil {
			//	//fmt.Println("failed LayerTypeSFlow:", pp.ErrorLayer().Error())
			//	go func(data []byte) {
			//		s := &layers.SFlowDatagram{}
			//		s.DecodeSampleFromBytes(data, gopacket.NilDecodeFeedback)
			//
			//		//b, err := json.Marshal(s)
			//		//if err != nil {
			//		//	fmt.Println(err.Error())
			//		//}
			//		//fmt.Println(data,string(b))
			//		for n, y := range s.FlowSamples {
			//			tmp := NewFlowSamples()
			//			tmp.InitOriginData(p)
			//			tmp.InitFlowSampleData(y)
			//			for num, yy := range y.Records {
			//				if g1, ok1 := yy.(layers.SFlowRawPacketFlowRecord); ok1 {
			//					tmp.ParseLayers(g1.Header)
			//					fmt.Println(data,s.AgentUptime,n,num,tmp.Data.Datagram.SrcIP,g1.Header.NetworkLayer().NetworkFlow().Src().String(),g1.Header.NetworkLayer().NetworkFlow().Dst().String(), tmp.SFlowRawPacketFlowRecord.Header.Bytes)
			//					//b, err := json.Marshal(tmp)
			//					//if err != nil {
			//					//	fmt.Println(err.Error())
			//					//}
			//					//if this.Udpbool {
			//					//	this.SendUdp(string(b),false)
			//					//} else {
			//					//	fmt.Println(s.AgentUptime,n,num,string(b))
			//					//}
			//				}
			//			}
			//		}
			//	}(udp.Payload)
			//}
			if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				//b, err := json.Marshal(got)
				//if err != nil {
				//	fmt.Println(err.Error())
				//}
				//fmt.Println(string(b))
				go func(datas []layers.SFlowFlowSample,got *layers.SFlowDatagram) {
					for n, y := range datas {
						//beego.Critical(len(y.Records),y.RecordCount)
						tmp := NewFlowSamples()
						tmp.InitOriginData(p)
						tmp.InitFlowSampleData(y)
						for num, yy := range y.Records {
							if g1, ok1 := yy.(layers.SFlowRawPacketFlowRecord); ok1 {
								tmp.ParseLayers(g1.Header)
								fmt.Println(got.AgentUptime,n,num,tmp.Data.Datagram.SrcIP, tmp.SFlowRawPacketFlowRecord.Header.Bytes)
								//b, err := json.Marshal(tmp)
								//if err != nil {
								//	fmt.Println(err.Error())
								//}
								//if this.Udpbool {
								//	this.SendUdp(string(b),false)
								//} else {
								//	fmt.Println(string(b))
								//}
							}
						}
					}
				}(got.FlowSamples,got)
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
		//fmt.Println(packet.Dump())
		p := gopacket.NewPacket(packet.Data(), layers.LayerTypeEthernet, gopacket.Default)
		if p.ErrorLayer() != nil {
			fmt.Println("failed LayerTypeEthernet:", p.ErrorLayer().Error())
		}

		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//fmt.Println("UDP layer detected.")
			udp, _ := udpLayer.(*layers.UDP)

			pp := gopacket.NewPacket(udp.Payload, layers.LayerTypeSFlow, gopacket.Default)
			if pp.ErrorLayer() != nil {

				go func(data []byte) {
					s := &layers.SFlowDatagram{}
					s.DecodeCounterFromBytes(data, gopacket.NilDecodeFeedback)
					if len(s.CounterSamples) != 0 {
						//beego.Error("Error out of bounds ")
						tmp := NewCounterFlow()
						tmp.InitOriginData(p)
						tmp.InitCounterSampleStruct(s)

						b, err := json.Marshal(tmp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if this.Udpbool {
							this.SendUdp(string(b), true)
						} else {
							fmt.Println(string(b))
						}
					}

				}(udp.Payload)
			}
			if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				go func(datas []layers.SFlowCounterSample) {
					if len(datas) > 0 {
						//beego.Error(udp.Payload)
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
							this.SendUdp(string(b), true)
						} else {
							fmt.Println(string(b))
						}
					}
				}(got.CounterSamples)
			}
		}

		//sflow := packet.Layer(layers.LayerTypeSFlow)
		//if sflow != nil {
		//	fmt.Println("SFLOW layer detected")
		//}
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
				//fmt.Println("failed :", pp.ErrorLayer().Error())
				go func(data []byte) {
					s := &layers.SFlowDatagram{}
					s.DecodeSampleFromBytes(data, gopacket.NilDecodeFeedback)
					for _, y := range s.FlowSamples {
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
									this.SendUdp(string(b), false)
								} else {
									fmt.Println(string(b))
								}
							}
						}
					}

					sc := &layers.SFlowDatagram{}
					sc.DecodeCounterFromBytes(data, gopacket.NilDecodeFeedback)
					if len(s.CounterSamples) != 0 {
						//beego.Error("Error out of bounds ")
						tmp := NewCounterFlow()
						tmp.InitOriginData(p)
						tmp.InitCounterSampleStruct(sc)

						b, err := json.Marshal(tmp)
						if err != nil {
							fmt.Println(err.Error())
						}
						if this.Udpbool {
							this.SendUdp(string(b), true)
						} else {
							fmt.Println(string(b))
						}
					}
				}(udp.Payload)
			} else if got, ok := pp.ApplicationLayer().(*layers.SFlowDatagram); ok {
				go func(Sample []layers.SFlowFlowSample, Counter []layers.SFlowCounterSample) {
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
										this.SendUdp(string(b), false)
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
							this.SendUdp(string(b), true)
						} else {
							fmt.Println(string(b))
						}
					}
				}(got.FlowSamples, got.CounterSamples)
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

				for _, x := range tmp.PayLoadToNetFlowV5(udp.Payload, packet.NetworkLayer().NetworkFlow().Src().String()) {
					this.SendUdp(x, false)
				}
				//beego.Error(len(data))
				//fmt.Println(data)
			}
		}(packet)
	}
}

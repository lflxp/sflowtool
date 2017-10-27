# Installation
Install latest version using Golang (recommended)

> go get github.com/lflxp/sflowtool

# sflowtool
sflow V5 and Netflow V5 parse by golang

# SflowV5 Struct

![](https://github.com/lflxp/sflow/master/SflowV5.png)

# NetFlowV5 Struct

![](https://github.com/lflxp/sflow/master/NetflowV5.png)

# Installation

go get [github.com/lflxp/sflowtool](https://github.com/lflxp/sflowtool)

# Usage

```
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
	SflowAll()
	//SflowSample()
	//SflowCounter()
	//NetflowV5()
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

```

## OutPut

those functions output json used by logstash to collected

# Example

> SFlowSample
>> SFlowSample just only detectd 5 layers
>> SFlowRawPacketFlowRecord
>> SFlowExtendedSwitchFlowRecord
>> SFlowExtendedRouterFlowRecord
>> SFlowExtendedGatewayFlowRecord
>> SFlowExtendedUserFlow

```
{
                           "Dropped" => 0,
                   "OutputInterface" => 112,
                              "Data" => {
        "DatagramVersion" => 5,
            "AgentUptime" => 2031516848,
               "Datagram" => {
            "SrcPort" => "4493",
              "DstIP" => "1.6.1.1",
            "DstPort" => "9999(distinct)",
             "SrcMac" => "70:ba:ee:04:e5:b5",
              "SrcIP" => "1.6.1.2",
             "DstMac" => "70:4d:7b:e8:c8:ee"
        },
             "SubAgentID" => 2,
           "AgentAddress" => "1.6.1.1",
         "SequenceNumber" => 613264,
            "SampleCount" => 2
    },
                      "SampleLength" => 180,
                      "SamplingRate" => 1000,
             "SFlowExtendedUserFlow" => {
         "DestinationCharSet" => "",
        "SFlowBaseFlowRecord" => {
                    "Format" => "",
              "EnterpriseID" => "",
            "FlowDataLength" => 0
        },
          "DestinationUserID" => "",
              "SourceCharSet" => "",
               "SourceUserID" => ""
    },
                     "SourceIDIndex" => "112",
                      "EnterpriseID" => "Standard SFlow",
                              "path" => "/tmp/123",
              "InputInterfaceFormat" => 0,
                            "Format" => "Expanded Flow Sample",
                        "@timestamp" => 2017-10-27T06:14:41.138Z,
          "SFlowRawPacketFlowRecord" => {
                "FrameLength" => 68,
               "HeaderLength" => 64,
                     "Header" => {
                  "SrcPort" => "63422",
             "Ipv4_version" => 4,
                 "Ipv4_ihl" => 5,
                 "Ipv4_ttl" => 128,
              "FlowRecords" => 80,
            "Ipv4_protocol" => "TCP",
                 "Ipv4_tos" => 0,
                  "DstPort" => "80(http)",
                    "SrcIP" => "1.6.1.81",
                    "Bytes" => 68,
                  "Packets" => 1,
                    "DstIP" => "1.1.1.2",
                   "SrcMac" => "44:8d:5c:23:c8:9c",
                   "DstMac" => "d4:61:ff:35:ee:f7"
        },
        "SFlowBaseFlowRecord" => {
                    "Format" => "Raw Packet Flow Record",
              "EnterpriseID" => "Standard SFlow",
            "FlowDataLength" => 80
        },
             "HeaderProtocol" => "ETHERNET-ISO88023",
             "PayloadRemoved" => 4
    },
                    "InputInterface" => 109,
                    "SequenceNumber" => 4140992,
             "OutputInterfaceFormat" => 0,
                          "@version" => "1",
                              "host" => "lxpdeiMac.local",
     "SFlowExtendedSwitchFlowRecord" => {
         "SFlowBaseFlowRecord" => {
                    "Format" => "Extended Switch Flow Record",
              "EnterpriseID" => "Standard SFlow",
            "FlowDataLength" => 16
        },
        "OutgoingVLANPriority" => 0,
        "IncomingVLANPriority" => 0,
                "IncomingVLAN" => 120,
                "OutgoingVLAN" => 0
    },
                       "RecordCount" => 3,
                     "SourceIDClass" => "Single Interface",
                        "SamplePool" => 4151408156,
     "SFlowExtendedRouterFlowRecord" => {
             "NextHopSourceMask" => 21,
           "SFlowBaseFlowRecord" => {
                    "Format" => "Extended Router Flow Record",
              "EnterpriseID" => "Standard SFlow",
            "FlowDataLength" => 16
        },
        "NextHopDestinationMask" => 0,
                       "NextHop" => "1.6.1.5"
    },
    "SFlowExtendedGatewayFlowRecord" => {
                         "AS" => 0,
                  "LocalPref" => 0,
                     "ASPath" => nil,
                "Communities" => nil,
                    "NextHop" => "",
                     "PeerAS" => 0,
        "SFlowBaseFlowRecord" => {
                    "Format" => "",
              "EnterpriseID" => "",
            "FlowDataLength" => 0
        },
                   "SourceAS" => 0,
                "ASPathCount" => 0
    }
}
```

> SFlowCounter
>> SFlowCounterSample contain 3 layers
>> SFlowGenericInterfaceCounters
>> SFlowEthernetCounters
>> SFlowProcessorCounters

```
{
                             "Data" => {
        "DatagramVersion" => 5,
            "AgentUptime" => 3163458152,
               "Datagram" => {
            "SrcPort" => "9999(distinct)",
              "DstIP" => "1.6.1.1",
            "DstPort" => "9999(distinct)",
             "SrcMac" => "70:ba:ff:f4:e5:bf",
              "SrcIP" => "1.1.1.2",
             "DstMac" => "f0:4d:7b:bf:cf:ff"
        },
             "SubAgentID" => 2,
           "AgentAddress" => "4.4.4.53",
         "SequenceNumber" => 1270179,
            "SampleCount" => 3
    },
            "SFlowEthernetCounters" => {
                        "FCSErrors" => 0,
                    "SQETestErrors" => 0,
                     "SymbolErrors" => 0,
            "DeferredTransmissions" => 0,
                   "LateCollisions" => 0,
            "SingleCollisionFrames" => 0,
              "ExcessiveCollisions" => 0,
         "InternalMacReceiveErrors" => 0,
           "SFlowBaseCounterRecord" => {
                    "Format" => "Ethernet Interface Counters",
              "EnterpriseID" => "Standard SFlow",
            "FlowDataLength" => 52
        },
        "InternalMacTransmitErrors" => 0,
                    "FrameTooLongs" => 0,
                  "AlignmentErrors" => 0,
               "CarrierSenseErrors" => 0,
          "MultipleCollisionFrames" => 0
    },
                     "SampleLength" => 172,
                    "SourceIDIndex" => "72",
                     "EnterpriseID" => "Standard SFlow",
                             "path" => "/tmp/123",
                           "Format" => "Expanded Counter Sample",
                       "@timestamp" => 2017-10-27T07:29:31.012Z,
                   "SequenceNumber" => 2853,
                         "@version" => "1",
                             "host" => "lxpdeiMac.local",
           "SFlowProcessorCounters" => {
                     "OneMinCpu" => 0,
                   "TotalMemory" => 0,
        "SFlowBaseCounterRecord" => {
                    "Format" => "",
              "EnterpriseID" => "",
            "FlowDataLength" => 0
        },
                    "FiveMinCpu" => 0,
                    "FiveSecCpu" => 0,
                    "FreeMemory" => 0
    },
                      "RecordCount" => 2,
                    "SourceIDClass" => "Single Interface",
    "SFlowGenericInterfaceCounters" => {
                        "IfType" => 6,
             "IfInMulticastPkts" => 82195,
                 "IfOutDiscards" => 27715,
             "IfPromiscuousMode" => 2,
             "IfInBroadcastPkts" => 1,
                      "IfStatus" => 3,
                  "IfInDiscards" => 0,
        "SFlowBaseCounterRecord" => {
                    "Format" => "Generic Interface Counters",
              "EnterpriseID" => "Standard SFlow",
            "FlowDataLength" => 88
        },
                   "IfDirection" => 1,
            "IfOutBroadcastPkts" => 0,
                       "IfSpeed" => 10000000000,
                   "IfOutErrors" => 0,
                 "IfInUcastPkts" => 4108242914,
                    "IfInOctets" => 103765034886283,
            "IfOutMulticastPkts" => 82441,
                       "IfIndex" => 72,
                   "IfOutOctets" => 102304620226454,
                "IfOutUcastPkts" => 3022758445,
                    "IfInErrors" => 0,
             "IfInUnknownProtos" => 0
    }
}
```

> NetFlowV5

```
{
                "dst_as" => 0,
               "in_pkts" => 1,
        "first_switched" => 536214234,
           "l4_src_port" => 51228,
         "ipv4_next_hop" => "10.6.32.5",
    "sampling_algorithm" => 0,
             "unix_nsec" => 0,
              "unix_sec" => 1509089578,
                  "path" => "/tmp/123",
              "in_bytes" => 40,
              "protocol" => 6,
             "tcp_flags" => 17,
                  "host" => "1.6.1.6",
              "@version" => "1",
           "l4_dst_port" => 80,
                "src_as" => 0,
           "output_snmp" => 50,
              "dst_mask" => 0,
               "src_tos" => 0,
         "ipv4_dst_addr" => "1.121.1.2",
              "src_mask" => 32,
               "version" => 5,
                "uptime" => 536423914,
          "flow_records" => 25,
          "flow_seq_num" => 245097954,
         "ipv4_src_addr" => "1.2.3.4",
           "engine_type" => 0,
            "@timestamp" => 2017-10-27T07:33:10.012Z,
             "engine_id" => 1,
            "input_snmp" => 0,
         "last_switched" => 536214234,
     "sampling_interval" => 0
}

```
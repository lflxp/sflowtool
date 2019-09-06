# Installation
Install latest version using Golang (recommended)

> go get -insecure github.com/lflxp//sflowtool

# sflowtool
sflow V5 and Netflow V5 parse by golang

# SflowV5 Struct

![](https://github.com/lflxp/sflowtool/blob/master/SflowV5.png)

# NetFlowV5 Struct

![](https://github.com/lflxp/sflowtool/blob/master/NetflowV5.png)

# Installation

go get [github.com/lflxp//sflowtool](https://github.com/lflxp//sflowtool)

# Usage

> sudo ./sflowtool -ee enp1s0 -p 9999 -t all -ises

```
Usage of ./sflowtool:
  -chost string
    	udp CounterSample 传输主机:端口 (default "127.0.0.1:7777")
  -d string
    	Enable certain debug selectors
  -debug
    	是否开启debug model
  -e	Log to stderr and disable syslog/file output
  -ee string
    	网卡名 (default "en0")
  -es string
    	elasticsearch 5.6 接口地址 (default "http://127.0.0.1:9200")
  -graceful
    	listen on open fd (after forking)
  -host string
    	udp SFlowSample And Netflow 传输主机:端口 (default "127.0.0.1:6666")
  -index string
    	es index name, example: sflow-2019-09-06 (default "sflow")
  -ises
    	是否开启output到elasticsearch
  -p string
    	端口 (default "6343")
  -s string
    	协议 (default "udp")
  -socketorder string
    	previous initialization order - used when more than one listener was started
  -strict.perms
    	Strict permission checking on config files (default true)
  -t string
    	类型:all(sflowSample|Counter),counter(SflowCounter),sample(SflowSample),netflow (default "all")
  -udp
    	是否开启udp数据传输,默认不开启
  -v	Log at INFO level

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
   "Data": {
      "Datagram": {
         "SrcMac": "70:99:99:04:99:99",
         "DstMac": "70:4d:99:99:99:99",
         "SrcIP": "99.99.99.205",
         "DstIP": "99.99.99.8",
         "SrcPort": "9999(distinct)",
         "DstPort": "9999(distinct)"
      },
      "DatagramVersion": 5,
      "AgentAddress": "99.99.99.53",
      "SubAgentID": 2,
      "SequenceNumber": 1275756,
      "AgentUptime": 3164307152,
      "SampleCount": 2
   },
   "EnterpriseID": "Standard SFlow",
   "Format": "Expanded Flow Sample",
   "SampleLength": 244,
   "SequenceNumber": 1251869,
   "SourceIDClass": "Single Interface",
   "SourceIDIndex": "71",
   "SamplingRate": 20000,
   "SamplePool": 3990725044,
   "Dropped": 0,
   "InputInterfaceFormat": 0,
   "InputInterface": 71,
   "OutputInterfaceFormat": 0,
   "OutputInterface": 114,
   "RecordCount": 3,
   "SFlowRawPacketFlowRecord": {
      "SFlowBaseFlowRecord": {
         "EnterpriseID": "Standard SFlow",
         "Format": "Raw Packet Flow Record",
         "FlowDataLength": 144
      },
      "HeaderProtocol": "ETHERNET-ISO88023",
      "FrameLength": 1518,
      "PayloadRemoved": 4,
      "HeaderLength": 128,
      "Header": {
         "FlowRecords": 144,
         "Packets": 1,
         "Bytes": 1518,
         "SrcMac": "99:8c:40:99:99:99",
         "DstMac": "99:8c:40:99:99:ab",
         "SrcIP": "99.99.99.26",
         "DstIP": "99.99.99.57",
         "Ipv4_version": 4,
         "Ipv4_ihl": 5,
         "Ipv4_tos": 0,
         "Ipv4_ttl": 62,
         "Ipv4_protocol": "TCP",
         "SrcPort": "49165",
         "DstPort": "33851"
      }
   },
   "SFlowExtendedSwitchFlowRecord": {
      "SFlowBaseFlowRecord": {
         "EnterpriseID": "Standard SFlow",
         "Format": "Extended Switch Flow Record",
         "FlowDataLength": 16
      },
      "IncomingVLAN": 0,
      "IncomingVLANPriority": 0,
      "OutgoingVLAN": 0,
      "OutgoingVLANPriority": 0
   },
   "SFlowExtendedRouterFlowRecord": {
      "SFlowBaseFlowRecord": {
         "EnterpriseID": "Standard SFlow",
         "Format": "Extended Router Flow Record",
         "FlowDataLength": 16
      },
      "NextHop": "99.99.99.206",
      "NextHopSourceMask": 22,
      "NextHopDestinationMask": 21
   },
   "SFlowExtendedGatewayFlowRecord": {
      "SFlowBaseFlowRecord": {
         "EnterpriseID": "",
         "Format": "",
         "FlowDataLength": 0
      },
      "NextHop": "",
      "AS": 0,
      "SourceAS": 0,
      "PeerAS": 0,
      "ASPathCount": 0,
      "ASPath": null,
      "Communities": null,
      "LocalPref": 0
   },
   "SFlowExtendedUserFlow": {
      "SFlowBaseFlowRecord": {
         "EnterpriseID": "",
         "Format": "",
         "FlowDataLength": 0
      },
      "SourceCharSet": "",
      "SourceUserID": "",
      "DestinationCharSet": "",
      "DestinationUserID": ""
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
   "Data": {
      "Datagram": {
         "SrcMac": "99:99:ef:04:99:99",
         "DstMac": "99:99:7b:b8:99:99",
         "SrcIP": "99.99.99.205",
         "DstIP": "99.99.99.8",
         "SrcPort": "9999(distinct)",
         "DstPort": "9999(distinct)"
      },
      "DatagramVersion": 5,
      "AgentAddress": "99.99.99.53",
      "SubAgentID": 2,
      "SequenceNumber": 1280989,
      "AgentUptime": 3164899152,
      "SampleCount": 3
   },
   "EnterpriseID": "Standard SFlow",
   "Format": "Expanded Counter Sample",
   "SampleLength": 172,
   "SequenceNumber": 2865,
   "SourceIDClass": "Single Interface",
   "SourceIDIndex": "72",
   "RecordCount": 2,
   "SFlowGenericInterfaceCounters": {
      "SFlowBaseCounterRecord": {
         "EnterpriseID": "Standard SFlow",
         "Format": "Generic Interface Counters",
         "FlowDataLength": 88
      },
      "IfIndex": 72,
      "IfType": 6,
      "IfSpeed": 10000000000,
      "IfDirection": 1,
      "IfStatus": 3,
      "IfInOctets": 104160000662999,
      "IfInUcastPkts": 92171299,
      "IfInMulticastPkts": 82243,
      "IfInBroadcastPkts": 1,
      "IfInDiscards": 0,
      "IfInErrors": 0,
      "IfInUnknownProtos": 0,
      "IfOutOctets": 992414418961899,
      "IfOutUcastPkts": 9939958927,
      "IfOutMulticastPkts": 82489,
      "IfOutBroadcastPkts": 0,
      "IfOutDiscards": 28017,
      "IfOutErrors": 0,
      "IfPromiscuousMode": 2
   },
   "SFlowEthernetCounters": {
      "SFlowBaseCounterRecord": {
         "EnterpriseID": "Standard SFlow",
         "Format": "Ethernet Interface Counters",
         "FlowDataLength": 99
      },
      "AlignmentErrors": 0,
      "FCSErrors": 0,
      "SingleCollisionFrames": 0,
      "MultipleCollisionFrames": 0,
      "SQETestErrors": 0,
      "DeferredTransmissions": 0,
      "LateCollisions": 0,
      "ExcessiveCollisions": 0,
      "InternalMacTransmitErrors": 0,
      "CarrierSenseErrors": 0,
      "FrameTooLongs": 0,
      "InternalMacReceiveErrors": 0,
      "SymbolErrors": 0
   },
   "SFlowProcessorCounters": {
      "SFlowBaseCounterRecord": {
         "EnterpriseID": "",
         "Format": "",
         "FlowDataLength": 0
      },
      "FiveSecCpu": 0,
      "OneMinCpu": 0,
      "FiveMinCpu": 0,
      "TotalMemory": 0,
      "FreeMemory": 0
   }
}
```

> NetFlowV5

```
{
   "version": 5,
   "flow_records": 30,
   "uptime": 537043304,
   "unix_sec": 1509090197,
   "unix_nsec": 0,
   "flow_seq_num": 245226516,
   "engine_type": 0,
   "engine_id": 1,
   "sampling_interval": 0,
   "input_snmp": 50,
   "output_snmp": 0,
   "in_pkts": 1,
   "in_bytes": 476,
   "first_switched": 537025674,
   "last_switched": 537025674,
   "l4_src_port": 53,
   "l4_dst_port": 60657,
   "tcp_flags": 0,
   "protocol": 17,
   "src_tos": 0,
   "src_as": 0,
   "dst_as": 0,
   "src_mask": 0,
   "dst_mask": 32,
   "host": "99.99.99.6",
   "sampling_algorithm": 0,
   "ipv4_src_addr": "99.99.99.19",
   "ipv4_dst_addr": "99.99.99.25",
   "ipv4_next_hop": "0.0.0.0"
}
```

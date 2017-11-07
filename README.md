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


```
Usage of ./sflowtool:
  -chost string
    	udp CounterSample 传输主机:端口 (default "127.0.0.1:7777")
  -e string
    	网卡名 (default "en0")
  -graceful
    	listen on open fd (after forking)
  -host string
    	udp SFlowSample And Netflow 传输主机:端口 (default "127.0.0.1:6666")
  -p string
    	端口 (default "6343")
  -s string
    	协议 (default "udp")
  -socketorder string
    	previous initialization order - used when more than one listener was started
  -t string
    	类型:all(sflowSample|Counter),counter(SflowCounter),sample(SflowSample),netflow (default "all")
  -udp
    	是否开启udp数据传输,默认不开启

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

0 0 0 5 0 0 0 1 10 6 0 254 0 0 0 6 1 92 141 29 115 64 76 123 0 0 0 5 0 0 0 3 0 0 0 244 6 239 203 65 0 0 0 0 0 0 0 101 0 0 3 232 152 221 198 204 0 0 0 0 0 0 0 0 0 0 0 195 0 0 0 0 0 0 0 101 0 0 0 3 0 0 3 234 0 0 0 16 0 0 0 1 172 18 10 154 0 0 0 24 0 0 0 0 0 0 3 233 0 0 0 16 0 0 11 184 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 144 0 0 0 1 0 0 0 203 0 0 0 4 0 0 0 128 0 35 137 204 162 20 112 186 239 4 229 181 129 0 11 184 8 0 69 0 0 181 171 167 64 0 63 6 229 222 10 6 5 106 10 10 144 67 184 214 7 143 100 166 78 121 15 3 226 44 128 24 0 114 229 20 0 0 1 1 8 10 149 106 123 68 119 75 154 133 196 226 149 15 178 149 129 193 94 62 140 58 36 42 160 210 78 77 106 15 97 99 21 59 32 34 249 107 109 151 78 41 73 222 145 14 43 4 118 78 181 236 176 202 147 183 91 161 169 151 44 40 30 174 93 208 124 208 0 0 0 3 0 0 0 244 6 239 203 66 0 0 0 0 0 0 0 101 0 0 3 232 152 221 198 204 0 0 0 0 0 0 0 0 0 0 0 101 0 0 0 0 0 0 0 207 0 0 0 3 0 0 3 234 0 0 0 16 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 21 0 0 3 233 0 0 0 16 0 0 11 184 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 144 0 0 0 1 0 0 5 242 0 0 0 4 0 0 0 128 112 186 239 4 229 181 0 35 137 204 162 20 129 0 11 184 8 0 69 0 5 220 59 157 64 0 52 6 239 70 125 64 134 41 10 6 8 201 0 80 204 182 143 51 233 137 92 78 194 255 80 16 13 186 255 130 0 0 91 194 194 240 195 23 14 29 14 247 245 78 133 107 136 119 6 69 157 57 207 137 1 10 193 2 220 61 95 184 29 101 176 44 159 138 126 192 33 240 39 45 109 58 211 186 43 102 124 197 131 117 13 98 129 112 17 44 121 188 66 79 231 89 137 149 3 134 214 244 0 0 0 3 0 0 0 244 6 239 203 67 0 0 0 0 0 0 0 101 0 0 3 232 152 221 198 204 0 0 0 0 0 0 0 0 0 0 0 177 0 0 0 0 0 0 0 101 0 0 0 3 0 0 3 234 0 0 0 16 0 0 0 1 172 18 10 154 0 0 0 24 0 0 0 0 0 0 3 233 0 0 0 16 0 0 11 184 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 144 0 0 0 1 0 0 5 242 0 0 0 4 0 0 0 128 0 35 137 204 162 20 112 186 239 4 229 181 129 0 11 184 8 0 69 0 5 220 76 79 64 0 63 6 232 4 10 6 5 125 10 15 232 54 217 160 192 53 143 5 191 2 250 82 156 27 128 16 5 144 71 247 0 0 1 1 8 10 56 142 255 223 21 170 111 105 187 177 108 114 26 235 121 112 104 149 255 207 129 124 201 241 195 201 106 231 25 8 64 187 227 218 45 86 245 81 98 233 210 193 75 89 128 176 148 121 71 172 32 82 173 253 100 225 62 195 152 251 99 191 37 71 155 97 0 0 0 3 0 0 0 188 6 239 203 68 0 0 0 0 0 0 0 101 0 0 3 232 152 221 198 204 0 0 0 0 0 0 0 0 0 0 0 101 0 0 0 0 0 0 0 212 0 0 0 3 0 0 3 234 0 0 0 16 0 0 0 1 10 6 32 2 0 0 0 0 0 0 0 22 0 0 3 233 0 0 0 16 0 0 11 184 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 88 0 0 0 1 0 0 0 74 0 0 0 4 0 0 0 70 112 186 239 4 229 181 0 35 137 204 162 20 129 0 11 184 8 0 69 0 0 52 99 138 64 0 49 6 73 103 121 201 4 228 10 6 20 32 31 69 243 9 162 211 160 158 20 81 228 107 80 24 0 33 248 190 0 0 3 225 46 5 0 0 0 167 153 1 0 1 0 212 0 0
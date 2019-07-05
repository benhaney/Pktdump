package decoders

import (
  "pktdump/utils"
  "strings"
  "fmt"
)

type Ipv6 struct {
  HopLimit uint
  Protocol uint
  IpSrc [16]byte
  IpDst [16]byte
  Data Packet
}

func NewIpv6(data []byte) Ipv6 {
  i := Ipv6{}
  i.Decode(data)
  return i
}

func (i *Ipv6) Decode(data []byte) {
  i.HopLimit = uint(data[7])
  i.Protocol = uint(data[6])
  copy(i.IpSrc[:], data[8:24])
  copy(i.IpDst[:], data[24:40])
  switch i.Protocol {
    case 0x06:
      // TCP
      t := NewTcp(data[40:])
      i.Data = &t
    case 0x11:
      // UDP
      u := NewUdp(data[40:])
      i.Data = &u
  }
}

func (i *Ipv6) String() string {
  var str strings.Builder
  i.Append(&str, 0)
  return str.String()
}

func (i *Ipv6) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin IPv4 ==\n", ind))
  str.WriteString(fmt.Sprintf("%sHop Limit: %d\n", ind, i.HopLimit))
  str.WriteString(fmt.Sprintf("%sIP Address (src): %s\n", ind, utils.FmtIpv6(i.IpSrc)))
  str.WriteString(fmt.Sprintf("%sIP Address (dst): %s\n", ind, utils.FmtIpv6(i.IpDst)))
  i.Data.Append(str, indent + 2)
}

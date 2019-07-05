package decoders

import (
  "strings"
  "fmt"
)

type Ipv4 struct {
  Ihl uint
  Ttl uint
  Protocol uint
  IpSrc [4]byte
  IpDst [4]byte
  Data Packet
}

func NewIpv4(data []byte) Ipv4 {
  i := Ipv4{}
  i.Decode(data)
  return i
}

func (i *Ipv4) Decode(data []byte) {
  i.Ihl = uint((data[0] & 0x0F) * 4)
  i.Ttl = uint(data[8])
  i.Protocol = uint(data[9])
  copy(i.IpSrc[:], data[12:16])
  copy(i.IpDst[:], data[16:20])
  switch i.Protocol {
    case 0x06:
      // TCP
      t := NewTcp(data[i.Ihl:])
      i.Data = &t
    case 0x11:
      // UDP
  }
}

func (i *Ipv4) String() string {
  var str strings.Builder
  i.Append(&str, 0)
  return str.String()
}

func (i *Ipv4) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin IPv4 ==\n", ind))
  str.WriteString(fmt.Sprintf("%sTTL: %d\n", ind, i.Ttl))
  str.WriteString(fmt.Sprintf("%sIP Address (src): %s\n", ind, fmtIpv4(i.IpSrc)))
  str.WriteString(fmt.Sprintf("%sIP Address (dst): %s\n", ind, fmtIpv4(i.IpDst)))
  i.Data.Append(str, indent + 2)
}

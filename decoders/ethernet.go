package decoders

import (
  "pktdump/utils"
  "encoding/binary"
  "strings"
  "fmt"
)

type Ethernet struct {
  MacDst [6]byte
  MacSrc [6]byte
  Ethertype uint16
  Data Packet
}

func NewEthernet(data []byte) Ethernet {
  e := Ethernet{}
  e.Decode(data)
  return e
}

func (e *Ethernet) Decode(data []byte) {
  const preamble uint64 = 0xaaaaaaaaaaaaaaab
  if binary.BigEndian.Uint64(data[0:8]) == preamble {
    // This is a layer 1 ethernet frame, so shave off the preamble
    data = data[8:]
  }
  copy(e.MacDst[:], data[0:6])
  copy(e.MacSrc[:], data[6:12])
  e.Ethertype = binary.BigEndian.Uint16(data[12:14])
  switch e.Ethertype {
    case 0x0800:
      // IPv4
      i := NewIpv4(data[14:])
      e.Data = &i
    case 0x86DD:
      // IPv6
  }
}

func (e *Ethernet) String() string {
  var str strings.Builder
  e.Append(&str, 0)
  return str.String()
}

func (e *Ethernet) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin Ethernet ==\n", ind))
  str.WriteString(fmt.Sprintf("%sMAC (dst): %s\n", ind, utils.FmtMac(e.MacDst)))
  str.WriteString(fmt.Sprintf("%sMAC (src): %s\n", ind, utils.FmtMac(e.MacSrc)))
  str.WriteString(fmt.Sprintf("%sEthertype: %#04x\n", ind, e.Ethertype))
  e.Data.Append(str, indent + 2)
}

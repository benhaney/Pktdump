package decoders

import (
  "encoding/binary"
  "strings"
  "fmt"
)

type Udp struct {
  PortSrc uint16
  PortDst uint16
  Data Packet
}

func NewUdp(data []byte) Udp {
  t := Udp{}
  t.Decode(data)
  return t
}

func (t *Udp) Decode(data []byte) {
  t.PortSrc = binary.BigEndian.Uint16(data[0:2])
  t.PortDst = binary.BigEndian.Uint16(data[2:4])
  r := NewRaw(data[8:])
  t.Data = &r
}

func (t *Udp) String() string {
  var str strings.Builder
  t.Append(&str, 0)
  return str.String()
}

func (t *Udp) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin UDP ==\n", ind))
  str.WriteString(fmt.Sprintf("%sPort (src): %d\n", ind, t.PortSrc))
  str.WriteString(fmt.Sprintf("%sPort (dst): %d\n", ind, t.PortDst))
  t.Data.Append(str, indent + 2)
}

package decoders

import (
  "encoding/binary"
  "strings"
  "fmt"
)

type Tcp struct {
  Offset uint
  PortSrc uint16
  PortDst uint16
  Data Packet
}

func NewTcp(data []byte) Tcp {
  t := Tcp{}
  t.Decode(data)
  return t
}

func (t *Tcp) Decode(data []byte) {
  t.Offset = uint((data[12] >> 4) * 4)
  t.PortSrc = binary.BigEndian.Uint16(data[0:2])
  t.PortDst = binary.BigEndian.Uint16(data[2:4])
  r := NewRaw(data[t.Offset:])
  t.Data = &r
}

func (t *Tcp) String() string {
  var str strings.Builder
  t.Append(&str, 0)
  return str.String()
}

func (t *Tcp) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin TCP ==\n", ind))
  str.WriteString(fmt.Sprintf("%sPort (src): %d\n", ind, t.PortSrc))
  str.WriteString(fmt.Sprintf("%sPort (dst): %d\n", ind, t.PortDst))
  t.Data.Append(str, indent + 2)
}

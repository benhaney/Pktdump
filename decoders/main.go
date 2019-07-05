package decoders

import "strings"

type Packet interface {
  Decode([]byte)
  String() string
  Append(*strings.Builder, int)
}

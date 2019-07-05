package decoders

import (
  "fmt"
)

func fmtIpv4(d [4]byte) string {
  return fmt.Sprintf("%d.%d.%d.%d", d[0], d[1], d[2], d[3])
}

func fmtMac(d [6]byte) string {
  return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", d[0], d[1], d[2], d[3], d[4], d[5])
}

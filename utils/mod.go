package utils

import (
  "fmt"
)

func FmtIpv4(d [4]byte) string {
  return fmt.Sprintf("%d.%d.%d.%d", d[0], d[1], d[2], d[3])
}

func FmtIpv6(d [16]byte) string {
  return fmt.Sprintf("%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x",
    d[0], d[1], d[2], d[3], d[4], d[5], d[6], d[7], d[8], d[9], d[10], d[11], d[12], d[13], d[14], d[15])
}

func FmtMac(d [6]byte) string {
  return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", d[0], d[1], d[2], d[3], d[4], d[5])
}

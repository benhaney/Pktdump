package main

import (
  "io/ioutil"
  "os"
  "pktdump/decoders"
  "fmt"
)

func main() {
  var data []byte
  var err error
  data, err = ioutil.ReadAll(os.Stdin)
  if err != nil { panic(err) }
  frame := decoders.NewEthernet(data)
  fmt.Print(&frame)
}

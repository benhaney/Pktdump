package main

import (
  "io/ioutil"
  "os"
)

func main() {
  var data []byte
  var err error
  data, err = ioutil.ReadAll(os.Stdin)
  if err != nil { panic(err) }
}

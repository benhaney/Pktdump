package decoders

import (
  "strings"
  "fmt"
)

type Raw struct {
  Data []byte
}

func NewRaw(data []byte) Raw {
  r := Raw{}
  r.Decode(data)
  return r
}

func (r *Raw) Decode(data []byte) {
  r.Data = data
}

func (r *Raw) String() string {
  var str strings.Builder
  r.Append(&str, 0)
  return str.String()
}

func (r *Raw) Append(str *strings.Builder, indent int) {
  ind := strings.Repeat(" ", indent)
  str.WriteString(fmt.Sprintf("%s== Begin Raw ==\n\n", ind))
  str.WriteString(fmt.Sprintf("%s\n\n", string(r.Data)))
  str.WriteString(fmt.Sprintf("%s== End Raw ==\n", ind))
}

package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/pretty"
)

func Pretty(v interface{})  {
	bf := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(bf)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	bs := pretty.Pretty(bf.Bytes())
	fmt.Println(string(bs))
}

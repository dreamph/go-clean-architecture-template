package json

import (
	gojson "github.com/goccy/go-json"
)

var (
	Marshal   = gojson.Marshal
	Unmarshal = gojson.Unmarshal
)

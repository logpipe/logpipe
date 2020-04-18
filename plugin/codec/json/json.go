package json

import (
	"encoding/json"
	"errors"
	"github.com/tk103331/logpipe/core"
)

func init() {
	core.RegCodec("json", func(conf core.CodecConf) core.Codec {
		return &JSONCodec{}
	})
}

type JSONCodec struct {
}

func (*JSONCodec) Encode(event core.Event) (interface{}, error) {
	bytes, err := json.Marshal(event)
	return string(bytes), err
}
func (*JSONCodec) Decode(data interface{}) (core.Event, error) {
	event := core.Event{}
	if str, ok := data.(string); ok {
		err := json.Unmarshal([]byte(str), &event)
		return event, err
	}
	return event, errors.New("unsupported")
}

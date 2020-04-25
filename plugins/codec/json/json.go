package json

import (
	"encoding/json"
	"errors"
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin"
)

func init() {
	plugin.RegCodec(&JSONCodecBuilder{})
}

type JSONCodec struct {
}

func (*JSONCodec) Encode(event core.Event) (interface{}, error) {
	bytes, err := json.Marshal(event.Map())
	return string(bytes), err
}
func (*JSONCodec) Decode(data interface{}) (core.Event, error) {
	event := core.NewEvent(data)
	var fields map[string]interface{}
	if str, ok := data.(string); ok {
		err := json.Unmarshal([]byte(str), &fields)
		if err != nil {
			event.AddTag("_jsonparsefailure")
		} else {
			for k, v := range fields {
				event.AddField(k, v)
			}
		}
		return event, err
	}
	return event, errors.New("unsupported")
}

type JSONCodecBuilder struct {
}

func (b *JSONCodecBuilder) Kind() string {
	return "json"
}

func (b *JSONCodecBuilder) Build(value config.Value) core.Codec {
	return &JSONCodec{}
}

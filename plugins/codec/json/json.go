package json

import (
	"encoding/json"
	"errors"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/plugin"
)

func init() {
	plugin.RegCodec(&JSONCodecBuilder{})
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

type JSONCodecBuilder struct {
}

func (b *JSONCodecBuilder) Kind() string {
	return "json"
}

func (b *JSONCodecBuilder) Build(value config.Value) core.Codec {
	return &JSONCodec{}
}

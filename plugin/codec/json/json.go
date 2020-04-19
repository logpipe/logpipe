package json

import (
	"encoding/json"
	"errors"
	"github.com/tk103331/logpipe/core"
)

func init() {
	core.RegCodec("json", &JSONCodecBuilder{})
}

type JSONCodecBuilder struct {
}

func (J JSONCodecBuilder) NewConf() core.CodecConf {
	return JSONCodecConf{}
}

func (J JSONCodecBuilder) Build(conf core.CodecConf) core.Codec {
	panic("implement me")
}

type JSONCodecConf struct {
	core.CodecConf
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

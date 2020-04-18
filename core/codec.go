package core

type Codec interface {
	Encoder
	Decoder
}

type Encoder interface {
	Encode(event Event) (interface{}, error)
}

type Decoder interface {
	Decode(data interface{}) (Event, error)
}

type CodecConf struct {
	BaseConf
	Kind string
}

package core

type InputConf struct {
	BaseConf
	Name  string
	Kind  string
	Codec *CodecConf
}

func (c *InputConf) Load(value *Value) error {

	return value.Parse(c)
}

type Input interface {
	Start(ctx Context) error
	Stop() error
}

type BaseInput struct {
	Name  string
	Kind  string
	Codec Decoder
}

func (*BaseInput) Start(ctx Context) error {
	return nil
}
func (*BaseInput) Stop() error {
	return nil
}

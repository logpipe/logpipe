package core

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

type InputConf interface {
	Conf
	NameConf
	KindConf
}
type BaseInputConf struct {
	BaseConf
	BaseNameConf
	BaseKindConf
	Codec CodecConf
}

func (c *BaseInputConf) Load(value *Value) error {
	return value.Parse(c)
}

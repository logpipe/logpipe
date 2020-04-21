package core

type Input interface {
	Start(consumer func(event Event)) error
	Stop() error
}

type BaseInput struct {
	codec Codec
}

func (*BaseInput) Start(_ func(event Event)) error {
	return nil
}
func (*BaseInput) Stop() error {
	return nil
}

func (i *BaseInput) SetCodec(codec Codec) {
	i.codec = codec
}

func (i *BaseInput) Codec() Codec {
	return i.codec
}

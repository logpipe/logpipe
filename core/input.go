package core

type Input interface {
	Start(consumer func(event Event)) error
	Stop() error
}

type BaseInput struct {
	codec Codec
	ctx   *Context
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

func (i *BaseInput) Context() *Context {
	return i.ctx
}

func (i *BaseInput) SetContext(ctx *Context) {
	i.ctx = ctx
}

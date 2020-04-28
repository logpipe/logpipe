package core

type Output interface {
	Start() error
	Stop() error
	Output(event Event) error
}

type BaseOutput struct {
	codec Codec
	ctx   *Context
}

func (*BaseOutput) Start() error {
	return nil
}

func (*BaseOutput) Stop() error {
	return nil
}

func (*BaseOutput) Output(_ Event) error {
	return nil
}

func (i *BaseOutput) SetCodec(codec Codec) {
	i.codec = codec
}

func (i *BaseOutput) Codec() Codec {
	return i.codec
}

func (i *BaseOutput) Context() *Context {
	return i.ctx
}

func (i *BaseOutput) SetContext(ctx *Context) {
	i.ctx = ctx
}

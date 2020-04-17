package core

type Input interface {
	Start(ctx Context) error
	Stop(ctx Context) error
}

type BaseInput struct {
	Codec Decoder
}

func (*BaseInput) Start(ctx Context) error {
	return nil
}
func (*BaseInput) Stop(ctx Context) error {
	return nil
}

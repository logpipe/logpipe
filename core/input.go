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

package core

type Input interface {
	Start(consumer func(event Event)) error
	Stop() error
}

type BaseInput struct {
	Name  string
	Kind  string
	Codec Decoder
}

func (*BaseInput) Start(_ func(event Event)) error {
	return nil
}
func (*BaseInput) Stop() error {
	return nil
}

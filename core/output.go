package core

type Output interface {
	Start() error
	Stop() error
	Output(event Event) error
}

type BaseOutput struct {
	Name  string
	Kind  string
	Cond  []*Cond
	Codec Encoder
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

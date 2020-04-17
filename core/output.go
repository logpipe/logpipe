package core

type Output interface {
	Start() error
	Stop() error
	Output(event Event) error
}

type BaseOutput struct {
	Conf  string
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

package core

type OutputConf interface {
	Conf
	NameConf
	KindConf
}

type BaseOutputConf struct {
	BaseConf
	BaseNameConf
	BaseKindConf
	Codec CodecConf
}

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

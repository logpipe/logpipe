package config

type InputConf struct {
	name  string
	kind  string
	codec CodecConf
	spec  Value
}

func (i *InputConf) Name() string {
	return i.name
}
func (i *InputConf) Kind() string {
	return i.kind
}
func (i *InputConf) Spec() Value {
	return i.spec
}

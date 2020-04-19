package config

import "gopkg.in/yaml.v3"

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

func (i *InputConf) Codec() CodecConf {
	return i.codec
}

func (i *InputConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	i.name = value.GetString("name")
	i.kind = value.GetString("kind")
	err := value.Get("codec").Parse(&(i.codec))
	if err != nil {
		return err
	}
	i.spec = *value.Get("spec")
	return nil
}

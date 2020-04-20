package config

import "gopkg.in/yaml.v3"

type OutputConf struct {
	name  string
	kind  string
	cond  []CondConf
	codec CodecConf
	spec  Value
}

func (c *OutputConf) Name() string {
	return c.name
}

func (c *OutputConf) Kind() string {
	return c.kind
}

func (c *OutputConf) Spec() Value {
	return c.spec
}

func (i *OutputConf) Codec() CodecConf {
	return i.codec
}

func (i *OutputConf) Cond() []CondConf {
	return i.cond
}

func (i *OutputConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	i.name = value.GetString("name")
	i.kind = value.GetString("kind")

	err := value.Get("cond").Parse(&(i.cond))
	if err != nil {
		return err
	}
	err = value.Get("codec").Parse(&(i.codec))
	if err != nil {
		return err
	}
	i.spec = *value.Get("spec")
	return nil
}

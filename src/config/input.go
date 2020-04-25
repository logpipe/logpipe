package config

import "gopkg.in/yaml.v3"

type InputConf struct {
	name   string
	kind   string
	codec  CodecConf
	action []ActionConf
	spec   Value
}

func (c *InputConf) Name() string {
	return c.name
}

func (c *InputConf) Kind() string {
	return c.kind
}

func (c *InputConf) Spec() Value {
	return c.spec
}

func (c *InputConf) Codec() CodecConf {
	return c.codec
}

func (c *InputConf) Action() []ActionConf {
	return c.action
}

func (c *InputConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	c.name = value.GetString("name")
	c.kind = value.GetString("kind")
	err := value.Get("codec").Parse(&(c.codec))
	if err != nil {
		return err
	}
	err = value.Get("action").Parse(&(c.action))
	if err != nil {
		return err
	}
	c.spec = *value.Get("spec")
	return nil
}

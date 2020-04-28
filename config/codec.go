package config

import (
	"gopkg.in/yaml.v3"
)

type CodecConf struct {
	kind string
	spec Value
}

func (c *CodecConf) Kind() string {
	return c.kind
}

func (c *CodecConf) Spec() Value {
	return c.spec
}

func (c *CodecConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	c.kind = value.GetString("kind")
	err := value.Get("spec").Parse(c.spec)
	if err != nil {
		return err
	}
	return nil
}

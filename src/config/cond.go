package config

import (
	"gopkg.in/yaml.v3"
)

type CondConf struct {
	kind string
	spec Value
}

func (c *CondConf) Kind() string {
	return c.kind
}

func (c *CondConf) Spec() Value {
	return c.spec
}

func (c *CondConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node}
	c.kind = value.GetString("kind")
	err := value.Get("spec").Parse(&c.spec)
	if err != nil {
		return err
	}
	return nil
}

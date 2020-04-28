package config

import "gopkg.in/yaml.v3"

type ActionConf struct {
	kind string
	spec Value
}

func (c *ActionConf) Kind() string {
	return c.kind
}

func (c *ActionConf) Spec() Value {
	return c.spec
}

func (c *ActionConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	c.kind = value.GetString("kind")
	err := value.Get("spec").Parse(&c.spec)
	if err != nil {
		return err
	}
	return nil
}

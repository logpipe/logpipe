package config

import "gopkg.in/yaml.v3"

type FilterConf struct {
	name   string
	kind   string
	cond   []CondConf
	action []ActionConf
	spec   Value
}

func (c *FilterConf) Name() string {
	return c.name
}

func (c *FilterConf) Kind() string {
	return c.kind
}

func (c *FilterConf) Spec() Value {
	return c.spec
}

func (c *FilterConf) Cond() []CondConf {
	return c.cond
}

func (c *FilterConf) Action() []ActionConf {
	return c.action
}

func (c *FilterConf) UnmarshalYAML(node *yaml.Node) error {
	value := Value{node: node}
	c.name = value.GetString("name")
	c.kind = value.GetString("kind")
	err := value.Get("cond").Parse(&(c.cond))
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

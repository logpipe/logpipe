package config

type FilterConf struct {
	name string
	kind string
	spec Value
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

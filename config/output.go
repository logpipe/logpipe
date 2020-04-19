package config

type OutputConf struct {
	name  string
	kind  string
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

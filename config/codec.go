package config

type CodecConf struct {
	kind string
	spec Value
}

func (c *CodecConf) Kind() string {
	return c.kind
}

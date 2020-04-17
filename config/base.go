package config

type Conf interface {
	Load(value Value)
}

type BaseConf struct {
	Name  string
	Value Value
}

func (c *BaseConf) Load(value Value) {
	c.Value = value
}

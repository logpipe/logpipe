package config

type Conf interface {
	Load(value *Value)
}

type BaseConf struct {
	Name string
	Kind string
	Spec *Value
}

func (c *BaseConf) Load(value *Value) error {
	c.Spec = value
	return c.Spec.Parse(c)
}

func (c *BaseConf) GetValue(key string) *BaseConf {
	conf := BaseConf{}
	_ = conf.Load(c.Spec.GetValue(key))
	return &conf
}
func (c *BaseConf) GetArray(key string) []*BaseConf {
	values := c.Spec.GetArray(key)
	confs := make([]*BaseConf, len(values))
	for i, value := range values {
		conf := BaseConf{}
		_ = conf.Load(value)
		confs[i] = &conf
	}
	return confs
}

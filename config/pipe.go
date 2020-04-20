package config

type PipeConf struct {
	name    string
	async   bool
	inputs  []InputConf
	filters []FilterConf
	outputs []OutputConf
}

func (p *PipeConf) Load(value *Value) error {
	p.name = value.GetString("name")
	p.async = value.GetBool("async")
	err := value.Get("inputs").Parse(&(p.inputs))
	if err != nil {
		return err
	}
	err = value.Get("filters").Parse(&(p.filters))
	if err != nil {
		return err
	}
	err = value.Get("outputs").Parse(&(p.outputs))
	if err != nil {
		return err
	}
	return nil
}

func (p *PipeConf) Name() string {
	return p.name
}
func (p *PipeConf) Async() bool {
	return p.async
}
func (p *PipeConf) Inputs() []InputConf {
	return p.inputs
}
func (p *PipeConf) Filters() []FilterConf {
	return p.filters
}
func (p *PipeConf) Outputs() []OutputConf {
	return p.outputs
}

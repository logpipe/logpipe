package config

import "path/filepath"

type PipeConf struct {
	name    string
	file    string
	async   bool
	log     LogConf
	inputs  []InputConf
	filters []FilterConf
	outputs []OutputConf
	vars    map[string]interface{}
}

func (p *PipeConf) Load(value *Value) error {
	p.name = value.GetString("name")
	p.file = value.file
	p.async = value.GetBool("async")
	value.Get("log").Parse(&(p.log))
	if p.log.Level == "" {
		p.log.Level = GetAppConf().Log.Level
	}
	if p.log.Path == "" {
		p.log.Path = p.name
	}
	if !filepath.IsAbs(p.log.Path) {
		p.log.Path = filepath.Join(GetAppConf().Log.Path, p.log.Path)
	}
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
	err = value.Get("vars").Parse(p.vars)
	if err != nil {
		return err
	}
	return nil
}

func (p *PipeConf) Name() string {
	return p.name
}
func (p *PipeConf) File() string {
	return p.file
}
func (p *PipeConf) Async() bool {
	return p.async
}
func (p *PipeConf) Log() LogConf {
	return p.log
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
func (p *PipeConf) Vars() map[string]interface{} {
	return p.vars
}

package core

import (
	"log"
)

type kind struct {
	Kind string
}

type PipeConf struct {
	BaseConf
	Name    string
	inputs  []InputConf
	filters []FilterConf
	outputs []OutputConf
}

func (c *PipeConf) Load(value *Value) error {
	c.BaseConf.Load(value)
	// parse input
	input := value.Get("input")
	if input != nil && input.IsArray() {
		values := input.Array()
		c.inputs = make([]InputConf, len(values))
		for i, v := range values {
			kind := v.GetString("kind")
			if builder, ok := inputBuilders[kind]; ok {
				conf := builder.NewConf()
				conf.Load(v)
				c.inputs[i] = conf
			}
		}
	}
	// parse filter
	filter := value.Get("filter")
	if filter != nil && filter.IsArray() {
		values := filter.Array()
		c.filters = make([]FilterConf, len(values))
		for i, v := range values {
			kind := v.GetString("kind")
			if builder, ok := filterBuilders[kind]; ok {
				conf := builder.NewConf()
				conf.Load(v)
				c.filters[i] = conf
			}
		}
	}
	// parse output
	output := value.Get("output")
	if output != nil && output.IsArray() {
		values := output.Array()
		c.outputs = make([]OutputConf, len(values))
		for i, v := range values {
			kind := v.GetString("kind")
			if builder, ok := outputBuilders[kind]; ok {
				conf := builder.NewConf()
				conf.Load(v)
				c.outputs[i] = conf
			}
		}
	}
	return value.Parse(c)
}

type Pipe struct {
	name    string
	conf    PipeConf
	ctx     Context
	inputs  []Input
	filters []Filter
	outputs []Output
}

func (p *Pipe) Init(pipeConf PipeConf) {
	p.conf = pipeConf
	p.inputs = make([]Input, len(pipeConf.inputs))
	for i, conf := range pipeConf.inputs {
		kind := conf.GetKind()
		if builder, ok := inputBuilders[kind]; ok {
			p.inputs[i] = builder.Build(conf)
		}
	}
	p.filters = make([]Filter, len(pipeConf.filters))
	for i, conf := range pipeConf.filters {
		kind := conf.GetKind()
		if builder, ok := filterBuilders[kind]; ok {
			p.filters[i] = builder.Build(conf)
		}
	}
	p.outputs = make([]Output, len(pipeConf.outputs))
	for i, conf := range pipeConf.outputs {
		kind := conf.GetKind()
		if builder, ok := outputBuilders[kind]; ok {
			p.outputs[i] = builder.Build(conf)
		}
	}
	p.ctx = Context{pipe: p}
}

func (p *Pipe) Start() {
	for _, output := range p.outputs {
		if output != nil {
			err := output.Start()
			if err != nil {
				log.Println(err)
			}
		}
	}
	for _, input := range p.inputs {
		if input != nil {
			err := input.Start(p.ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (p *Pipe) Stop() {

	for _, input := range p.inputs {
		err := input.Stop()
		if err != nil {
			log.Println(err)
		}
	}

	for _, output := range p.outputs {
		err := output.Stop()
		if err != nil {
			log.Println(err)
		}
	}
}

func (p *Pipe) Input(event Event) {
	temp := event
	if p.filters != nil && len(p.filters) > 0 {
		for _, filter := range p.filters {
			if !temp.IsEmpty() {
				temp = filter.Filter(temp)
			}
		}
	}
	if p.outputs != nil && len(p.outputs) > 0 {
		for _, output := range p.outputs {
			if !temp.IsEmpty() {
				if output != nil {
					err := output.Output(temp)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

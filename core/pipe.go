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
	input := value.GetValue("input")
	var inputKinds []kind
	input.Parse(&inputKinds)
	c.inputs = make([]InputConf, len(inputKinds))
	for i, k := range inputKinds {
		if builder, ok := inputBuilders[k.Kind]; ok {
			conf := builder.NewConf()
			c.inputs[i] = conf
		}
	}
	input.Parse(&c.inputs)
	// parse filter
	filter := value.GetValue("filter")
	var filterKinds []kind
	filter.Parse(&filterKinds)
	c.filters = make([]FilterConf, len(filterKinds))
	for i, k := range filterKinds {
		if builder, ok := filterBuilders[k.Kind]; ok {
			conf := builder.NewConf()
			c.filters[i] = conf
		}
	}
	filter.Parse(&c.filters)
	// output filter
	output := value.GetValue("output")
	var outputKinds []kind
	output.Parse(&outputKinds)
	c.filters = make([]FilterConf, len(outputKinds))
	for i, k := range outputKinds {
		if builder, ok := outputBuilders[k.Kind]; ok {
			conf := builder.NewConf()
			c.outputs[i] = conf
		}
	}
	output.Parse(&c.filters)
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
		kind := conf.Kind
		if builder, ok := filterBuilders[kind]; ok {
			p.filters[i] = builder.Build(conf)
		}
	}
	p.outputs = make([]Output, len(pipeConf.outputs))
	for i, conf := range pipeConf.outputs {
		kind := conf.Kind
		if builder, ok := outputBuilders[kind]; ok {
			p.outputs[i] = builder.Build(conf)
		}
	}
}

func (p *Pipe) Start() {
	for _, output := range p.outputs {
		output.Start()
	}
	for _, input := range p.inputs {
		input.Start(p.ctx)
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
	for _, filter := range p.filters {
		if !temp.IsEmpty() {
			temp = filter.Filter(temp)
		}
	}
	for _, output := range p.outputs {
		if !temp.IsEmpty() {
			output.Output(temp)
		}
	}
}

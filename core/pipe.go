package core

import (
	"fmt"
	"go.uber.org/config"
	"log"
)

type PipeConf struct {
	BaseConf
	Name    string
	inputs  []*InputConf
	filters []*FilterConf
	outputs []*OutputConf
}

func (c *PipeConf) Load(value *Value) error {
	c.BaseConf.Load(value)

	var values []*config.Value
	value.GetValue("input").Parse(&values)
	fmt.Println(values)
	value.GetValue("input").Parse(&c.inputs)
	value.GetValue("filter").Parse(&c.filters)
	value.GetValue("output").Parse(&c.outputs)

	filterValues := value.GetArray("filter")
	c.filters = make([]*FilterConf, len(filterValues))
	for i, filterValue := range filterValues {
		filterConf := &FilterConf{}
		filterConf.Load(filterValue)
		c.filters[i] = filterConf
	}
	outputValues := value.GetArray("output")
	c.outputs = make([]*OutputConf, len(outputValues))
	for i, outputValue := range outputValues {
		outputConf := &OutputConf{}
		outputConf.Load(outputValue)
		c.outputs[i] = outputConf
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
		kind := conf.Kind
		if builder, ok := inputBuilders[kind]; ok {
			p.inputs[i] = builder(*conf)
		}
	}
	p.filters = make([]Filter, len(pipeConf.filters))
	for i, conf := range pipeConf.filters {
		kind := conf.Kind
		if builder, ok := filterBuilders[kind]; ok {
			p.filters[i] = builder(*conf)
		}
	}
	p.outputs = make([]Output, len(pipeConf.outputs))
	for i, conf := range pipeConf.outputs {
		kind := conf.Kind
		if builder, ok := outputBuilders[kind]; ok {
			p.outputs[i] = builder(*conf)
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

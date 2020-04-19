package core

import (
	"github.com/tk103331/logpipe/config"
	"log"
)

type Pipe struct {
	name    string
	conf    config.PipeConf
	ctx     Context
	inputs  []Input
	filters []Filter
	outputs []Output
}

func (p *Pipe) Init(pipeConf config.PipeConf) {
	p.conf = pipeConf
	p.inputs = make([]Input, len(pipeConf.Inputs))
	for i, conf := range pipeConf.Inputs {
		kind := conf.Kind()
		if builder, ok := inputBuilders[kind]; ok {
			p.inputs[i] = builder(conf)
		}
	}
	p.filters = make([]Filter, len(pipeConf.Filters))
	for i, conf := range pipeConf.Filters {
		kind := conf.Kind()
		if builder, ok := filterBuilders[kind]; ok {
			p.filters[i] = builder(conf)
		}
	}
	p.outputs = make([]Output, len(pipeConf.Outputs))
	for i, conf := range pipeConf.Outputs {
		kind := conf.Kind()
		if builder, ok := outputBuilders[kind]; ok {
			p.outputs[i] = builder(conf)
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

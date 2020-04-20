package engine

import (
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/plugin"
	"log"
)

type Pipe struct {
	name     string
	conf     config.PipeConf
	consumer func(event core.Event)
	inputs   []core.Input
	filters  []core.Filter
	outputs  []core.Output
}

func (p *Pipe) Init(pipeConf config.PipeConf) {
	p.conf = pipeConf
	p.inputs = make([]core.Input, len(pipeConf.Inputs))
	for i, conf := range pipeConf.Inputs {
		p.inputs[i] = plugin.BuildInput(conf)
	}
	p.filters = make([]core.Filter, len(pipeConf.Filters))
	for i, conf := range pipeConf.Filters {
		p.filters[i] = plugin.BuildFilter(conf)
	}
	p.outputs = make([]core.Output, len(pipeConf.Outputs))
	for i, conf := range pipeConf.Outputs {
		p.outputs[i] = plugin.BuildOutput(conf)
	}
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
			err := input.Start(p.input)
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

func (p *Pipe) input(event core.Event) {
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

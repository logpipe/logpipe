package engine

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin"
	"log"
)

type Pipe struct {
	name     string
	conf     config.PipeConf
	consumer func(event core.Event)
	inputs   []InputNode
	filters  []FilterNode
	outputs  []OutputNode
}

func (p *Pipe) Init(pipeConf config.PipeConf) {
	p.conf = pipeConf
	logger := core.NewLogger(pipeConf.Log().Path, pipeConf.Log().Level)
	p.inputs = make([]InputNode, len(pipeConf.Inputs()))
	for i, conf := range pipeConf.Inputs() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		input := plugin.BuildInput(conf)
		if container, ok := input.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		actions := core.BuildActions(conf.Action())
		p.inputs[i] = InputNode{input: input, action: actions}
	}
	p.filters = make([]FilterNode, len(pipeConf.Filters()))
	for i, conf := range pipeConf.Filters() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		filter := plugin.BuildFilter(conf)
		if container, ok := filter.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		cond := core.BuildConds(conf.Cond())
		actions := core.BuildActions(conf.Action())
		p.filters[i] = FilterNode{filter: filter, cond: cond, action: actions}
	}
	p.outputs = make([]OutputNode, len(pipeConf.Outputs()))
	for i, conf := range pipeConf.Outputs() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		output := plugin.BuildOutput(conf)
		if container, ok := output.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		cond := core.BuildConds(conf.Cond())
		p.outputs[i] = OutputNode{output: output, cond: cond}
	}
	if p.conf.Async() {
		p.consumer = func(event core.Event) {
			go p.input(event)
		}
	} else {
		p.consumer = func(event core.Event) {
			p.input(event)
		}
	}
}

func (p *Pipe) Start() {
	for _, node := range p.outputs {
		if node.output != nil {
			err := node.output.Start()
			if err != nil {
				log.Println(err)
			}
		}
	}
	for _, node := range p.inputs {
		if node.input != nil {
			err := node.input.Start(func(event core.Event) {
				if !event.IsEmpty() {
					if node.action != nil && len(node.action) > 0 {
						node.action.Exec(&event)
					}
					p.consumer(event)
				}
			})
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (p *Pipe) Stop() {

	for _, node := range p.inputs {
		if node.input != nil {
			err := node.input.Stop()
			if err != nil {
				log.Println(err)
			}
		}
	}

	for _, node := range p.outputs {
		if node.output != nil {
			err := node.output.Stop()
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (p *Pipe) input(event core.Event) {
	temp := event
	if p.filters != nil && len(p.filters) > 0 {
		for _, node := range p.filters {
			if !temp.IsEmpty() && node.filter != nil {
				if node.cond != nil && len(node.cond) > 0 {
					if !node.cond.Test(event) {
						continue
					}
				}
				temp = node.filter.Filter(temp)
				if node.action != nil && len(node.action) > 0 {
					node.action.Exec(&event)
				}
			}
		}
	}
	if p.outputs != nil && len(p.outputs) > 0 {
		for _, node := range p.outputs {
			if !temp.IsEmpty() && node.output != nil {
				if node.cond != nil && len(node.cond) > 0 {
					if !node.cond.Test(event) {
						continue
					}
				}
				err := node.output.Output(temp)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

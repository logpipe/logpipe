package engine

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/log"
	"github.com/logpipe/logpipe/plugin"
)

type pipe struct {
	name     string
	conf     config.PipeConf
	consumer func(event core.Event)
	inputs   []inputNode
	filters  []filterNode
	outputs  []outputNode
}

func (p *pipe) Init(pipeConf config.PipeConf) {
	log.Info("init pipe [%v] from %v", pipeConf.Name())
	p.conf = pipeConf
	logger := log.NewLogger(pipeConf.Log().Path, pipeConf.Log().Level)
	p.inputs = make([]inputNode, len(pipeConf.Inputs()))
	for i, conf := range pipeConf.Inputs() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		input := plugin.BuildInput(conf)
		if container, ok := input.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		actions := core.BuildActions(conf.Action())
		p.inputs[i] = inputNode{name: conf.Name(), ctx: ctx, input: input, action: actions}
	}
	p.filters = make([]filterNode, len(pipeConf.Filters()))
	for i, conf := range pipeConf.Filters() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		filter := plugin.BuildFilter(conf)
		if container, ok := filter.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		cond := core.BuildConds(conf.Cond())
		actions := core.BuildActions(conf.Action())
		p.filters[i] = filterNode{name: conf.Name(), ctx: ctx, filter: filter, cond: cond, action: actions}
	}
	p.outputs = make([]outputNode, len(pipeConf.Outputs()))
	for i, conf := range pipeConf.Outputs() {
		ctx := core.NewContext(p.name, conf.Name(), conf.Kind(), logger)
		output := plugin.BuildOutput(conf)
		if container, ok := output.(core.ContextContainer); ok {
			container.SetContext(ctx)
		}
		cond := core.BuildConds(conf.Cond())
		p.outputs[i] = outputNode{name: conf.Name(), ctx: ctx, output: output, cond: cond}
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

func (p *pipe) Start() {
	for _, node := range p.outputs {
		if node.output != nil {
			err := node.output.Start()
			if err != nil {
				log.Error("starting output plugin [%v] error: %v", node.name, err.Error())
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
				log.Error("starting input plugin [%v] error: %v", node.name, err.Error())
			}
		}
	}
}

func (p *pipe) Stop() {

	for _, node := range p.inputs {
		if node.input != nil {
			err := node.input.Stop()
			if err != nil {
				log.Error("stopping input plugin [%v] error: %v", node.name, err.Error())
			}
		}
	}

	for _, node := range p.outputs {
		if node.output != nil {
			err := node.output.Stop()
			if err != nil {
				log.Error("stopping output plugin [%v] error: %v", node.name, err.Error())
			}
		}
	}
}

func (p *pipe) input(event core.Event) {
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
					log.Error("outputting error: %v", err.Error())
				}
			}
		}
	}
}

package engine

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin/buildin"
)

func DebugPipe(conf config.PipeConf) {
	p := pipe{}
	p.Init(conf)
	if len(p.inputs) == 0 {
		p.inputs = make([]inputNode, 1)
		p.inputs[0] = inputNode{input: &buildin.StdinInput{}}
	}
	if len(p.outputs) == 0 {
		p.outputs = make([]outputNode, 1)
		p.outputs[0] = outputNode{output: &buildin.StdoutOutput{}}
	}
	p.Start()
}

func DebugInput(input core.Input) {
	debug(input, nil, nil)
}

func DebugFilter(filter core.Filter) {
	debug(nil, filter, nil)
}
func DebugOutput(output core.Output) {
	debug(nil, nil, output)
}
func debug(input core.Input, filter core.Filter, output core.Output) {
	p := pipe{}
	p.Init(config.PipeConf{})
	p.inputs = make([]inputNode, 1)
	if input == nil {
		input = &buildin.StdinInput{}
	}
	p.inputs[0] = inputNode{input: input}
	if filter != nil {
		p.filters = make([]filterNode, 1)
		p.filters[0] = filterNode{filter: filter}
	}
	p.outputs = make([]outputNode, 1)
	if output == nil {
		output = &buildin.StdoutOutput{}
	}
	p.outputs[0] = outputNode{output: output}
	p.Start()
}

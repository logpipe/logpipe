package engine

import "github.com/tk103331/logpipe/core"

var (
	inputs  = make(map[string]InputBuilder)
	filters = make(map[string]FilterBuilder)
	outputs = make(map[string]OutputBuilder)
	codecs  = make(map[string]CodecBuilder)
)

type InputBuilder func(ctx core.Context) core.Input
type FilterBuilder func(ctx core.Context) core.Filter
type OutputBuilder func(ctx core.Context) core.Output
type CodecBuilder func(ctx core.Context) core.Codec

func RegInput(name string, builder InputBuilder) {
	inputs[name] = builder
}

func RegFilter(name string, builder FilterBuilder) {
	filters[name] = builder
}

func RegOutput(name string, builder OutputBuilder) {
	outputs[name] = builder
}

func RegCodec(name string, builder CodecBuilder) {
	codecs[name] = builder
}

package core

import "github.com/tk103331/logpipe/config"

var (
	inputBuilders  = make(map[string]InputBuilder)
	filterBuilders = make(map[string]FilterBuilder)
	outputBuilders = make(map[string]OutputBuilder)
	codecBuilders  = make(map[string]CodecBuilder)
)

type InputBuilder func(name string, codec Codec, spec config.Value) Input
type FilterBuilder func(name string, spec config.Value) Filter
type OutputBuilder func(name string, codec Codec, spec config.Value) Output
type CodecBuilder func(spec config.Value) Codec

func RegInput(name string, builder InputBuilder) {
	inputBuilders[name] = builder
}

func RegFilter(name string, builder FilterBuilder) {
	filterBuilders[name] = builder
}

func RegOutput(name string, builder OutputBuilder) {
	outputBuilders[name] = builder
}

func RegCodec(name string, builder CodecBuilder) {
	codecBuilders[name] = builder
}

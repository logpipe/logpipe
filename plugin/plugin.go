package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
)

var (
	inputBuilders  = make(map[string]InputBuilder)
	filterBuilders = make(map[string]FilterBuilder)
	outputBuilders = make(map[string]OutputBuilder)
	codecBuilders  = make(map[string]CodecBuilder)
)

type InputBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Input
}

type FilterBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Filter
}
type OutputBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Output
}
type CodecBuilder interface {
	Kind() string
	Build(spec config.Value) core.Codec
}

func RegInput(builder InputBuilder) {
	inputBuilders[builder.Kind()] = builder
}

func BuildInput(conf config.InputConf) core.Input {
	if builder, ok := inputBuilders[conf.Kind()]; ok {
		codec := BuildCodec(conf.Codec())
		input := builder.Build(conf.Name(), conf.Spec())
		if container, ok := input.(core.CodecContainer); ok {
			container.SetCodec(codec)
		}
	}
	return nil
}

func RegFilter(builder FilterBuilder) {
	filterBuilders[builder.Kind()] = builder
}

func BuildFilter(conf config.FilterConf) core.Filter {
	if builder, ok := filterBuilders[conf.Kind()]; ok {
		return builder.Build(conf.Name(), conf.Spec())
	}
	return nil
}

func RegOutput(builder OutputBuilder) {
	outputBuilders[builder.Kind()] = builder
}

func BuildOutput(conf config.OutputConf) core.Output {
	if builder, ok := outputBuilders[conf.Kind()]; ok {
		codec := BuildCodec(conf.Codec())
		output := builder.Build(conf.Name(), conf.Spec())
		if container, ok := output.(core.CodecContainer); ok {
			container.SetCodec(codec)
		}
	}
	return nil
}

func RegCodec(builder CodecBuilder) {
	codecBuilders[builder.Kind()] = builder
}

func BuildCodec(conf config.CodecConf) core.Codec {
	if builder, ok := codecBuilders[conf.Kind()]; ok {
		return builder.Build(conf.Spec())
	}
	return nil
}

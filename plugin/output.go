package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/log"
)

type OutputBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Output
}

type outputBuilderWrapper struct {
	kind    string
	builder func(name string, spec config.Value) core.Output
}

func (o *outputBuilderWrapper) Kind() string {
	return o.kind
}

func (o *outputBuilderWrapper) Build(name string, spec config.Value) core.Output {
	return o.builder(name, spec)
}

func RegisterOutputBuilder(builder OutputBuilder) bool {
	outputLock.Lock()
	defer outputLock.Unlock()
	kind := builder.Kind()
	log.Info("register output plugin [%v]", kind)
	if _, ok := outputBuilders[kind]; ok {
		log.Error("register output plugin [%v] error: plugin already exist", kind)
		return false
	}
	outputBuilders[kind] = builder
	return true
}

func RegisterOutput(kind string, builder func(name string, spec config.Value) core.Output) bool {
	return RegisterOutputBuilder(&outputBuilderWrapper{kind: kind, builder: builder})
}

func BuildOutput(conf config.OutputConf) core.Output {
	outputLock.RLock()
	defer outputLock.RUnlock()
	if builder, ok := outputBuilders[conf.Kind()]; ok {
		codec := BuildCodec(conf.Codec())
		output := builder.Build(conf.Name(), conf.Spec())
		if container, ok := output.(core.CodecContainer); ok {
			container.SetCodec(codec)
		}
	}
	return nil
}

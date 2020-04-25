package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/log"
)

type InputBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Input
}

type inputBuilderWrapper struct {
	kind    string
	builder func(name string, spec config.Value) core.Input
}

func (i *inputBuilderWrapper) Kind() string {
	return i.kind
}

func (i *inputBuilderWrapper) Build(name string, spec config.Value) core.Input {
	return i.builder(name, spec)
}

func RegisterInputBuilder(builder InputBuilder) bool {
	inputLock.Lock()
	defer inputLock.Unlock()
	kind := builder.Kind()
	if _, ok := inputBuilders[kind]; ok {
		log.Error("register input plugin [%v] error: plugin already exist", kind)
		return false
	}
	inputBuilders[kind] = builder
	return true
}

func RegisterInput(kind string, builder func(name string, spec config.Value) core.Input) bool {
	return RegisterInputBuilder(&inputBuilderWrapper{kind: kind, builder: builder})
}

func BuildInput(conf config.InputConf) core.Input {
	inputLock.RLock()
	defer inputLock.RUnlock()
	if builder, ok := inputBuilders[conf.Kind()]; ok {
		codec := BuildCodec(conf.Codec())
		input := builder.Build(conf.Name(), conf.Spec())
		if container, ok := input.(core.CodecContainer); ok {
			container.SetCodec(codec)
		}
	}
	return nil
}

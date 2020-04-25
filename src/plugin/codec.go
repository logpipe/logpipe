package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/log"
)

type CodecBuilder interface {
	Kind() string
	Build(spec config.Value) core.Codec
}

type codecBuilderWrapper struct {
	kind    string
	builder func(spec config.Value) core.Codec
}

func (c *codecBuilderWrapper) Kind() string {
	return c.kind
}

func (c *codecBuilderWrapper) Build(spec config.Value) core.Codec {
	return c.builder(spec)
}

func RegisterCodecBuilder(builder CodecBuilder) bool {
	codecLock.Lock()
	defer codecLock.Unlock()
	kind := builder.Kind()
	if _, ok := codecBuilders[kind]; ok {
		log.Error("register codec plugin [%v] error: plugin already exist", kind)
		return false
	}
	codecBuilders[kind] = builder
	return true
}

func RegisterCodec(kind string, builder func(spec config.Value) core.Codec) bool {
	return RegisterCodecBuilder(&codecBuilderWrapper{kind: kind, builder: builder})
}

func BuildCodec(conf config.CodecConf) core.Codec {
	codecLock.RLock()
	defer codecLock.RUnlock()
	if builder, ok := codecBuilders[conf.Kind()]; ok {
		return builder.Build(conf.Spec())
	}
	return nil
}

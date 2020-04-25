package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/log"
)

type FilterBuilder interface {
	Kind() string
	Build(name string, spec config.Value) core.Filter
}

type filterBuilderWrapper struct {
	kind    string
	builder func(name string, spec config.Value) core.Filter
}

func (f *filterBuilderWrapper) Kind() string {
	return f.kind
}

func (f *filterBuilderWrapper) Build(name string, spec config.Value) core.Filter {
	return f.builder(name, spec)
}

func RegisterFilterBuilder(builder FilterBuilder) bool {
	filterLock.Lock()
	defer filterLock.Unlock()
	kind := builder.Kind()
	if _, ok := filterBuilders[kind]; ok {
		log.Error("register filter plugin [%v] error: plugin already exist", kind)
		return false
	}
	filterBuilders[kind] = builder
	return true
}

func RegisterFilter(kind string, builder func(name string, spec config.Value) core.Filter) {
	RegisterFilterBuilder(&filterBuilderWrapper{kind: kind, builder: builder})
}

func BuildFilter(conf config.FilterConf) core.Filter {
	filterLock.RLock()
	defer filterLock.RUnlock()
	if builder, ok := filterBuilders[conf.Kind()]; ok {
		return builder.Build(conf.Name(), conf.Spec())
	}
	return nil
}

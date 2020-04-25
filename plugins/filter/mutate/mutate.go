package mutate

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin"
	"strings"
)

var OPS = make(map[string]MutateOpBuilder)

func init() {

	initOps()

	plugin.RegFilter(&MutateFilterBuilder{})
}

type MutateFilter struct {
	Ops []MutateOp
}

func (m *MutateFilter) Filter(event core.Event) core.Event {
	for _, o := range m.Ops {
		o.Exec(&event)
	}
	return event
}

type MutateOp interface {
	Exec(event *core.Event) error
}

type MutateOpBuilder func(exp string) MutateOp

func initOps() {
	OPS["add_tag"] = func(exp string) MutateOp {
		return &AddTagOp{Tag: exp[6 : len(exp)-1]}
	}
	OPS["add_field"] = func(exp string) MutateOp {
		p := exp[6 : len(exp)-1]
		params := strings.Split(p, ",")
		return &AddFieldOp{Field: params[0], Value: params[1]}
	}
	OPS["set_field"] = func(exp string) MutateOp {
		p := exp[6 : len(exp)-1]
		params := strings.Split(p, ",")
		return &AddFieldOp{Field: params[0], Value: params[1]}
	}
}

type MutateFilterBuilder struct {
}

func (b *MutateFilterBuilder) Kind() string {
	return "mutate"
}

func (b *MutateFilterBuilder) Build(name string, spec config.Value) core.Filter {
	return &MutateFilter{}
}

package mutate

import (
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/engine"
	"strings"
)

var OPS = make(map[string]MutateOpBuilder)

func init() {

	initOps()

	engine.RegFilter("mutate", func(ctx core.Context) core.Filter {
		return &MutateFilter{}
	})
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

type AddTagOp struct {
	Tag string
}

func (o *AddTagOp) Exec(event *core.Event) error {
	event.AddTag(o.Tag)
	return nil
}

type AddFieldOp struct {
	Field string
	Value interface{}
}

func (o *AddFieldOp) Exec(event *core.Event) error {
	event.AddField(o.Field, o.Value)
	return nil
}

func initOps() {
	OPS["add_tag"] = func(exp string) MutateOp {
		return &AddTagOp{Tag: exp[6 : len(exp)-1]}
	}
	OPS["add_field"] = func(exp string) MutateOp {
		p := exp[6 : len(exp)-1]
		params := strings.Split(p, ",")
		return &AddFieldOp{Field: params[0], Value: params[1]}
	}
}

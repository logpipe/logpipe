package core

import (
	"github.com/logpipe/logpipe/config"
	"strings"
)

var conds = make(map[string]CondBuilder)

func init() {
	regCond(&HasTagCondBuilder{})
	regCond(&HasFieldCondBuilder{})
	regCond(&MatchFieldCondBuilder{})
}

func regCond(builder CondBuilder) {
	conds[builder.Kind()] = builder
}

func BuildCond(conf config.CondConf) Cond {
	if builder, ok := conds[conf.Kind()]; ok {
		return builder.Build(conf.Spec())
	}
	return nil
}
func BuildConds(confs []config.CondConf) Conds {
	conds := make([]Cond, len(confs))
	for i, c := range confs {
		conds[i] = BuildCond(c)
	}
	return conds
}

type Cond interface {
	Test(event Event) bool
}

type Conds []Cond

func (c *Conds) Test(event Event) bool {
	conds := ([]Cond)(*c)
	if len(conds) > 0 {
		for _, cond := range conds {
			if !cond.Test(event) {
				return false
			}
		}
	}
	return true
}

type CondBuilder interface {
	Kind() string
	Build(spec config.Value) Cond
}

type HasTagCond struct {
	tag string
}

func (c *HasTagCond) Test(event Event) bool {
	return event.HasTag(c.tag)
}

type HasTagCondBuilder struct {
}

func (*HasTagCondBuilder) Kind() string {
	return "has_tag"
}
func (*HasTagCondBuilder) Build(spec config.Value) Cond {
	tag := spec.GetString("tag")
	return &HasTagCond{tag: tag}
}

type HasFieldCond struct {
	field string
}

func (c *HasFieldCond) Test(event Event) bool {
	return event.HasField(c.field)
}

type HasFieldCondBuilder struct {
}

func (h *HasFieldCondBuilder) Kind() string {
	return "has_field"
}

func (h *HasFieldCondBuilder) Build(spec config.Value) Cond {
	field := spec.GetString("field")
	return &HasFieldCond{field: field}
}

type MatchFieldCond struct {
	field string
	op    string
	value interface{}
}

func (c *MatchFieldCond) Test(event Event) bool {
	v := event.GetField(c.field)
	if value, isStr := v.(string); isStr {
		switch c.op {
		case "equal":
			target, ok := c.value.(string)
			return ok && value == target
		case "not_equal":
			target, ok := c.value.(string)
			return ok && value != target
		case "start_with":
			target, ok := c.value.(string)
			return ok && strings.HasPrefix(value, target)
		case "end_with":
			target, ok := c.value.(string)
			return ok && strings.HasSuffix(value, target)
		case "contains":
			target, ok := c.value.(string)
			return ok && strings.Contains(value, target)
		case "length":
			target, ok := c.value.(int)
			return ok && len(value) == target
		case "lower_than":
			target, ok := c.value.(int)
			return ok && len(value) == target
		case "greater_than":
			target, ok := c.value.(int)
			return ok && len(value) == target
		default:
			return false
		}
	} else {

	}
	return false
}

type MatchFieldCondBuilder struct {
}

func (*MatchFieldCondBuilder) Kind() string {
	return "match_field"
}
func (*MatchFieldCondBuilder) Build(spec config.Value) Cond {
	field := spec.GetString("field")
	op := spec.GetString("op")
	var value interface{}
	_ = spec.Get("value").Parse(&value)
	return &MatchFieldCond{field: field, op: op, value: value}
}

package core

import (
	"strings"
)

type CondConf struct {
	BaseConf
	Kind string
}

type Cond interface {
	Name() string
	Test(event Event) bool
}

type HasTagCond struct {
	Tag string
}

func (*HasTagCond) Name() string {
	return "has_tag"
}
func (c *HasTagCond) Test(event Event) bool {
	return event.HasTag(c.Tag)
}

type HasFieldCond struct {
	Field string
}

func (*HasFieldCond) Name() string {
	return "has_field"
}
func (c *HasFieldCond) Test(event Event) bool {
	return event.HasField(c.Field)
}

type MatchFieldsCond struct {
	Field  string
	Op     string
	Values []interface{}
	origin string
}

func (*MatchFieldsCond) Name() string {
	return "match_fields"
}
func (c *MatchFieldsCond) Test(event Event) bool {
	v := event.GetField(c.Field)
	if value, isStr := v.(string); isStr {
		switch c.Op {
		case "=":
			target, ok := c.Values[0].(string)
			return ok && value == target
		case "!":
			target, ok := c.Values[0].(string)
			return ok && value != target
		case "^":
			target, ok := c.Values[0].(string)
			return ok && strings.HasPrefix(value, target)
		case "$":
			target, ok := c.Values[0].(string)
			return ok && strings.HasSuffix(value, target)
		case "@":
			target, ok := c.Values[0].(string)
			return ok && strings.Contains(value, target)
		case "#":
			target, ok := c.Values[0].(int)
			return ok && len(value) == target
		default:
			return false
		}
	} else {

	}
	return false
}

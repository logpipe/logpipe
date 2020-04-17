package add_field

import (
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/engine"
	"github.com/tk103331/logpipe/plugin/filter/add_tag"
)

func init() {
	engine.RegFilter("add_field", func(ctx core.Context) core.Filter {
		return &add_tag.AddTagFilter{}
	})
}

type AddFieldFilter struct {
	Field string
	Value interface{}
}

func (f *AddFieldFilter) Filter(event core.Event) core.Event {
	event.AddField(f.Field, f.Value)
	return event
}

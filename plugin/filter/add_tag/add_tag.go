package add_tag

import (
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/engine"
)

func init() {
	engine.RegFilter("add_tag", func(ctx core.Context) core.Filter {
		return &AddTagFilter{}
	})
}

type AddTagFilter struct {
	Tag string
}

func (f *AddTagFilter) Filter(event core.Event) core.Event {
	event.AddTag(f.Tag)
	return event
}

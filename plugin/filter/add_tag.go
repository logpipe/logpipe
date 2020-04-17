package filter

import "../../core"

type AddTagFilter struct {
	Tag string
}

func (f *AddTagFilter) Filter(event core.Event) core.Event {
	event.AddTag(f.Tag)
	return event
}

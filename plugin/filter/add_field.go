package filter

import "../../core"

type AddFieldFilter struct {
	Field string
	Value interface{}
}

func (f *AddFieldFilter) Filter(event core.Event) core.Event {
	event.AddField(f.Field, f.Value)
	return event
}

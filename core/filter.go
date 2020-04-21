package core

type Filter interface {
	Filter(event Event) Event
}

type BaseFilter struct {
}

func (f *BaseFilter) Filter(event Event) Event {
	return event
}

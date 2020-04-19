package core

type Filter interface {
	Filter(event Event) Event
}

type BaseFilter struct {
	Name string
	Kind string
	Cond []*Cond
}

func (f *BaseFilter) Filter(event Event) Event {
	return event
}

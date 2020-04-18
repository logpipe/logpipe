package core

type FilterConf struct {
	BaseConf
	Name string
	Kind string
	Cond *CondConf
}

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

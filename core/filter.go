package core

type Filter interface {
	Filter(event Event) Event
}

type BaseFilter struct {
	ctx *Context
}

func (f *BaseFilter) Filter(event Event) Event {
	return event
}

func (i *BaseFilter) Context() *Context {
	return i.ctx
}

func (i *BaseFilter) SetContext(ctx *Context) {
	i.ctx = ctx
}

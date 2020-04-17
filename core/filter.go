package core

type Filter interface {
	Filter(event Event) Event
}

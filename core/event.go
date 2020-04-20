package core

import (
	"sort"
	"time"
)

type Event struct {
	kind      string
	host      string
	timestamp int64
	source    interface{}
	fields    map[string]interface{}
	tags      []string
	empty     bool
}

func NewEvent(kind string, host string, source interface{}) Event {
	timestamp := time.Now().UnixNano()
	fields := make(map[string]interface{})
	tags := make([]string, 0)
	return Event{kind: kind, host: host, timestamp: timestamp, source: source, fields: fields, tags: tags}
}

func NewEmptyEvent() Event {
	return Event{empty: true}
}

func (e *Event) SetKind(kind string) {
	e.kind = kind
}
func (e *Event) Kind() string {
	return e.kind
}

func (e *Event) SetHost(host string) {
	e.host = host
}
func (e *Event) Host() string {
	return e.host
}

func (e *Event) Timestamp() int64 {
	return e.timestamp
}

func (e *Event) Source() interface{} {
	return e.source
}

func (e *Event) IsEmpty() bool {
	return e.empty
}

func (e *Event) HasTag(tag string) bool {
	if tag == "" {
		return true
	}
	for _, t := range e.tags {
		if t == tag {
			return true
		}
	}
	return false
}
func (e *Event) AddTag(tag string) {
	if e.HasTag(tag) {
		return
	}
	e.tags = append(e.tags, tag)
}
func (e *Event) RemoveTag(tag string) {
	index := -1
	for i, t := range e.tags {
		if t == tag {
			index = i
		}
	}
	e.tags = append(e.tags[0:index], e.tags[index:]...)
}
func (e *Event) SortTag(asc bool) {
	if asc {
		sort.Strings(e.tags)
	} else {
		reverse := sort.Reverse(sort.StringSlice(e.tags))
		sort.Sort(reverse)
	}
}

func (e *Event) HasField(field string) bool {
	if field == "" {
		return true
	}
	_, ok := e.fields[field]
	return ok
}
func (e *Event) GetField(field string) interface{} {
	if v, ok := e.fields[field]; ok {
		return v
	}
	return nil
}
func (e *Event) AddField(field string, value interface{}) {
	if e.HasField(field) {
		return
	}
	e.fields[field] = value
}
func (e *Event) SetField(field string, value interface{}) {
	e.fields[field] = value
}
func (e *Event) RemoveField(field string) interface{} {
	value := e.GetField(field)
	delete(e.fields, field)
	return value
}

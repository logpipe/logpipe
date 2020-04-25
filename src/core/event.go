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

func NewEvent(source interface{}) Event {
	timestamp := time.Now().UnixNano()
	fields := make(map[string]interface{})
	tags := make([]string, 0)
	return Event{timestamp: timestamp, source: source, fields: fields, tags: tags}
}

func NewEmptyEvent() Event {
	fields := make(map[string]interface{})
	tags := make([]string, 0)
	return Event{empty: true, fields: fields, tags: tags}
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

func (e *Event) Tags() []string {
	return e.tags[:]
}

func (e *Event) Fields() map[string]interface{} {
	fields := make(map[string]interface{}, len(e.fields))
	for k, v := range e.fields {
		fields[k] = v
	}
	return fields
}

func (e *Event) Map() map[string]interface{} {
	data := make(map[string]interface{})
	data["kind"] = e.kind
	data["host"] = e.host
	data["timestamp"] = e.timestamp
	data["source"] = e.source
	data["fields"] = e.fields
	data["tags"] = e.tags
	return data
}

func (e *Event) HasTag(tag string) bool {
	if e.empty {
		return false
	}
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
	if e.empty {
		return
	}
	if e.HasTag(tag) {
		return
	}
	e.tags = append(e.tags, tag)
}
func (e *Event) RemoveTag(tag string) {
	if e.empty {
		return
	}
	index := -1
	for i, t := range e.tags {
		if t == tag {
			index = i
		}
	}
	e.tags = append(e.tags[0:index], e.tags[index:]...)
}
func (e *Event) SortTag(asc bool) {
	if e.empty {
		return
	}
	if asc {
		sort.Strings(e.tags)
	} else {
		reverse := sort.Reverse(sort.StringSlice(e.tags))
		sort.Sort(reverse)
	}
}

func (e *Event) HasField(field string) bool {
	if e.empty {
		return false
	}
	if field == "" {
		return true
	}
	_, ok := e.fields[field]
	return ok
}
func (e *Event) GetField(field string) interface{} {
	if e.empty {
		return nil
	}
	if v, ok := e.fields[field]; ok {
		return v
	}
	return nil
}
func (e *Event) AddField(field string, value interface{}) {
	if e.empty {
		return
	}
	if e.HasField(field) {
		return
	}
	e.fields[field] = value
}
func (e *Event) SetField(field string, value interface{}) {
	if e.empty {
		return
	}
	e.fields[field] = value
}
func (e *Event) RemoveField(field string) interface{} {
	if e.empty {
		return nil
	}
	value := e.GetField(field)
	delete(e.fields, field)
	return value
}

package core

import "github.com/tk103331/logpipe/config"

var actions = make(map[string]ActionBuilder)

func init() {
	regAction(&AddTagActionBuilder{})
	regAction(&AddFieldActionBuilder{})
	regAction(&SetFieldActionBuilder{})
}

func regAction(builder ActionBuilder) {
	actions[builder.Kind()] = builder
}

func BuildAction(conf config.ActionConf) Action {
	if builder, ok := actions[conf.Kind()]; ok {
		return builder.Build(conf.Spec())
	}
	return nil
}

func BuildActions(confs []config.ActionConf) Actions {
	actions := make([]Action, len(confs))
	for i, conf := range confs {
		actions[i] = BuildAction(conf)
	}
	return actions
}

type Action interface {
	Exec(event *Event)
}

type Actions []Action

func (a Actions) Exec(event *Event) {
	actions := ([]Action)(a)
	for _, act := range actions {
		act.Exec(event)
	}
}

type ActionBuilder interface {
	Kind() string
	Build(spec config.Value) Action
}

type AddTagAction struct {
	tag string
}

func (a *AddTagAction) Exec(event *Event) {
	event.AddTag(a.tag)
}

type AddTagActionBuilder struct {
}

func (b *AddTagActionBuilder) Kind() string {
	return "add_tag"
}

func (b *AddTagActionBuilder) Build(spec config.Value) Action {
	tag := spec.GetString("tag")
	return &AddTagAction{tag: tag}
}

type AddFieldAction struct {
	field string
	value interface{}
}

func (a *AddFieldAction) Exec(event *Event) {
	event.AddField(a.field, a.value)
}

type AddFieldActionBuilder struct {
}

func (b *AddFieldActionBuilder) Kind() string {
	return "add_field"
}

func (b *AddFieldActionBuilder) Build(spec config.Value) Action {
	field := spec.GetString("field")
	var value interface{}
	spec.Get("value").Parse(&value)
	return &AddFieldAction{field: field, value: value}
}

type SetFieldAction struct {
	field string
	value interface{}
}

func (a *SetFieldAction) Exec(event *Event) {
	event.SetField(a.field, a.value)
}

type SetFieldActionBuilder struct {
}

func (b *SetFieldActionBuilder) Kind() string {
	return "set_field"
}

func (b *SetFieldActionBuilder) Build(spec config.Value) Action {
	field := spec.GetString("field")
	var value interface{}
	spec.Get("value").Parse(&value)
	return &SetFieldAction{field: field, value: value}
}

package mutate

import "github.com/logpipe/logpipe/core"

type AddFieldOp struct {
	Field string
	Value interface{}
}

func (o *AddFieldOp) Exec(event *core.Event) error {
	event.AddField(o.Field, o.Value)
	return nil
}

type SetFieldOp struct {
	Field string
	Value interface{}
}

func (o *SetFieldOp) Exec(event *core.Event) error {
	event.SetField(o.Field, o.Value)
	return nil
}

type RemoveFieldOp struct {
	Field string
}

func (o *RemoveFieldOp) Exec(event *core.Event) error {
	event.RemoveField(o.Field)
	return nil
}

type RenameFieldOp struct {
	Field string
}

func (o *RenameFieldOp) Exec(event *core.Event) error {
	value := event.GetField(o.Field)
	event.RemoveField(o.Field)
	event.AddField(o.Field, value)
	return nil
}

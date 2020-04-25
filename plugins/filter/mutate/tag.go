package mutate

import (
	"github.com/logpipe/logpipe/core"
)

type AddTagOp struct {
	Tag string
}

func (o *AddTagOp) Exec(event *core.Event) error {
	event.AddTag(o.Tag)
	return nil
}

type RemoveTagOp struct {
	Tag string
}

func (o *RemoveTagOp) Exec(event *core.Event) error {
	event.RemoveTag(o.Tag)
	return nil
}

type ReplaceTagOp struct {
	OldTag string
	NewTag string
}

func (o *ReplaceTagOp) Exec(event *core.Event) error {
	event.RemoveTag(o.OldTag)
	event.AddTag(o.NewTag)
	return nil
}

type SortTagOp struct {
	Asc bool
}

func (o *SortTagOp) Exec(event *core.Event) error {
	event.SortTag(o.Asc)
	return nil
}

package engine

import (
	"github.com/logpipe/logpipe/core"
)

type inputNode struct {
	name   string
	ctx    *core.Context
	input  core.Input
	action core.Actions
}

type filterNode struct {
	name   string
	ctx    *core.Context
	cond   core.Conds
	filter core.Filter
	action core.Actions
}

type outputNode struct {
	name   string
	ctx    *core.Context
	cond   core.Conds
	output core.Output
}

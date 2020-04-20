package engine

import (
	"github.com/tk103331/logpipe/core"
)

type InputNode struct {
	input  core.Input
	action core.Actions
}

type FilterNode struct {
	cond   core.Conds
	filter core.Filter
	action core.Actions
}

type OutputNode struct {
	cond   core.Conds
	output core.Output
}

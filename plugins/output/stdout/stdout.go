package stdout

import (
	"fmt"
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin"
)

func init() {
	plugin.RegOutput(&StdoutOutputBuilder{})
}

type StdoutOutput struct {
	core.BaseOutput
	spec StdoutOutputSpec
}

func (s *StdoutOutput) Output(event core.Event) error {
	if s.Codec() != nil {
		data, _ := s.Codec().Encode(event)
		fmt.Println(data)
	} else {
		fmt.Println(event)
	}
	return nil
}

type StdoutOutputSpec struct {
}

type StdoutOutputBuilder struct {
}

func (b *StdoutOutputBuilder) Kind() string {
	return "stdout"
}

func (b *StdoutOutputBuilder) Build(name string, spec config.Value) core.Output {
	var outputSpec StdoutOutputSpec
	spec.Parse(&outputSpec)
	return &StdoutOutput{spec: outputSpec}
}

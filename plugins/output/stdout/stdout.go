package stdout

import (
	"fmt"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/plugin"
)

func init() {
	plugin.RegOutput(&StdoutOutputBuilder{})
}

type StdoutOutput struct {
	core.BaseOutput
	name string
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
	return &StdoutOutput{name: name, spec: outputSpec}
}

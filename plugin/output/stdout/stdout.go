package stdout

import (
	"fmt"
	"github.com/tk103331/logpipe/core"
)

const OUTPUT_NAME = "stdout"

func init() {
	core.RegOutput(OUTPUT_NAME, &StdoutOutputBuilder{})
}

type StdoutOutput struct {
	core.BaseOutput
}

func (s *StdoutOutput) Output(event core.Event) error {
	if s.Codec != nil {
		data, _ := s.Codec.Encode(event)
		fmt.Println(data)
	} else {
		fmt.Println(event)
	}
	return nil
}

type StdoutOutputConf struct {
	core.BaseOutputConf
	Value string
}

func (c *StdoutOutputConf) GetKind() string {
	return OUTPUT_NAME
}

func (c StdoutOutputConf) Load(value *core.Value) error {
	return value.Parse(c)
}

type StdoutOutputBuilder struct {
}

func (s *StdoutOutputBuilder) NewConf() core.OutputConf {
	return &StdoutOutputConf{}
}

func (s *StdoutOutputBuilder) Build(conf core.OutputConf) core.Output {
	return &StdoutOutput{}
}

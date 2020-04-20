package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/plugin"
	"os"
)

func init() {
	plugin.RegInput(&StdinInputBuilder{})
}

type StdinInputSpec struct {
	Value1 bool
	Value2 string
	Value3 int
}

type StdinInput struct {
	name    string
	king    string
	codec   core.Decoder
	spec    StdinInputSpec
	stopped bool
}

func (s *StdinInput) Start(consumer func(event core.Event)) error {
	go s.run(consumer)
	return nil
}

func (s *StdinInput) Stop() error {
	s.stopped = true
	return nil
}

func (s *StdinInput) run(consumer func(event core.Event)) {
	reader := bufio.NewReader(os.Stdin)
	for !s.stopped {
		bytes, _, _ := reader.ReadLine()
		str := string(bytes)
		if s.codec != nil {
			event, err := s.codec.Decode(str)
			if err != nil {
				consumer(core.NewEmptyEvent())
			} else {
				consumer(event)
			}
		} else {
			event := core.NewEvent(str)
			consumer(event)
		}
	}
}

type StdinInputBuilder struct {
}

func (b *StdinInputBuilder) Kind() string {
	return "stdin"
}

func (b *StdinInputBuilder) Build(name string, codec core.Codec, spec config.Value) core.Input {
	var inputSpec StdinInputSpec
	spec.Parse(&inputSpec)
	return &StdinInput{name: name, codec: codec, spec: inputSpec}
}

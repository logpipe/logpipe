package stdin

import (
	"bufio"
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/core"
	"github.com/logpipe/logpipe/plugin"
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
	core.BaseInput
	spec StdinInputSpec
	stop chan struct{}
}

func (s *StdinInput) Start(consumer func(event core.Event)) error {
	s.Context().Info("starting...")
	go s.run(consumer)
	return nil
}

func (s *StdinInput) Stop() error {
	s.Context().Info("stopping...")
	s.stop <- struct{}{}
	return nil
}

func (s *StdinInput) run(consumer func(event core.Event)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-s.stop:
			break
		default:
		}
		bytes, _, _ := reader.ReadLine()
		str := string(bytes)
		if s.Codec() != nil {
			event, err := s.Codec().Decode(str)
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

func (b *StdinInputBuilder) Build(name string, spec config.Value) core.Input {
	var inputSpec StdinInputSpec
	spec.Parse(&inputSpec)
	return &StdinInput{spec: inputSpec}
}

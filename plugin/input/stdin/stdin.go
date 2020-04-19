package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"os"
)

func init() {
	core.RegInput("stdin", func(conf config.InputConf) core.Input {
		var spec StdinInputSpec
		conf.Spec().Parse(&spec)
		return &StdinInput{}
	})
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

func (s *StdinInput) Start(ctx core.Context) error {
	go s.run(ctx)
	return nil
}

func (s *StdinInput) Stop() error {
	s.stopped = true
	return nil
}

func (s *StdinInput) run(ctx core.Context) {
	reader := bufio.NewReader(os.Stdin)
	for !s.stopped {
		bytes, _, _ := reader.ReadLine()
		str := string(bytes)
		if s.codec != nil {
			event, _ := s.codec.Decode(str)
			ctx.Accept(event)
		} else {
			event := core.NewEvent("stdin", "localhost", string(bytes))
			ctx.Accept(event)
		}
	}
}

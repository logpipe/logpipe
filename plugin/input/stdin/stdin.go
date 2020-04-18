package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/engine"
	"os"
)

func init() {
	engine.RegInput("stdin", func(ctx core.Context) core.Input {
		value := config.Value{}
		conf := StdinInputConf{}
		conf.Load(value)
		return &StdinInput{conf: conf}
	})
}

type StdinInputConf struct {
	config.BaseConf
	Value1 bool
	Value2 string
	Value3 int
}

type StdinInput struct {
	core.BaseInput
	conf    StdinInputConf
	stopped bool
}

func (s *StdinInput) Start(ctx core.Context) error {
	go s.run(ctx)
	return nil
}

func (s *StdinInput) Stop(ctx core.Context) error {
	s.stopped = true
	return nil
}

func (s *StdinInput) run(ctx core.Context) {
	reader := bufio.NewReader(os.Stdin)
	for !s.stopped {
		bytes, _, _ := reader.ReadLine()
		str := string(bytes)
		if s.Codec != nil {
			event, _ := s.Codec.Decode(str)
			ctx.Accept(event)
		} else {
			event := core.NewEvent("stdin", "localhost", string(bytes))
			ctx.Accept(event)
		}
	}
}

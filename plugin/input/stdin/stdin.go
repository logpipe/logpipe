package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/core"
	"os"
)

func init() {
	core.RegInput("stdin", func(conf core.InputConf) core.Input {
		inputConf := StdinInputConf{}
		inputConf.Load(conf.Value())
		return &StdinInput{conf: inputConf}
	})
}

type StdinInputConf struct {
	core.BaseConf
	Value1 bool
	Value2 string
	Value3 int
}

func (c *StdinInputConf) Load(value *core.Value) error {
	return c.BaseConf.Load(value)
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

func (s *StdinInput) Stop() error {
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

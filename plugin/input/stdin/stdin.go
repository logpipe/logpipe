package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/core"
	"os"
)

const INPUT_NAME = "stdin"

func init() {
	core.RegInput(INPUT_NAME, &StdinInputBuilder{})
}

type StdinInputBuilder struct {
}

func (s *StdinInputBuilder) NewConf() core.InputConf {
	return &StdinInputConf{}
}

func (s *StdinInputBuilder) Build(conf core.InputConf) core.Input {
	return &StdinInput{conf: StdinInputConf{}}
}

type StdinInputConf struct {
	core.BaseInputConf
	Value1 bool
	Value2 string
	Value3 int
}

func (c *StdinInputConf) GetKind() string {
	return INPUT_NAME
}

func (c *StdinInputConf) Load(value *core.Value) error {
	return value.Parse(c)
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

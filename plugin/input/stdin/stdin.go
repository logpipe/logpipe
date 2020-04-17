package stdin

import (
	"bufio"
	"github.com/tk103331/logpipe/core"
	"os"
)

type StdinInput struct {
	core.BaseInput
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

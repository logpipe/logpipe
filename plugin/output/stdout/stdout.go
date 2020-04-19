package stdout

import (
	"fmt"
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
)

func init() {
	core.RegOutput("stdout", func(conf config.OutputConf) core.Output {
		return &StdoutOutput{}
	})
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

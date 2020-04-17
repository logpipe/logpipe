package file

import (
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/engine"
	"os"
)

func init() {
	engine.RegInput("file", func(ctx core.Context) core.Input {
		return &FileInput{}
	})
}

type FileInput struct {
	Path string
	file *os.File
}

func (i *FileInput) Start(ctx core.Context) error {
	file, err := os.OpenFile(i.Path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	i.file = file
	go func() {
		for {
			data := make([]byte, 1024)
			n, e := i.file.Read(data)
			if e == nil {
				event := core.NewEvent("file", "localhost", string(data[0:n]))
				ctx.Accept(event)
			}
		}
	}()
	return nil
}

func (i *FileInput) Stop(ctx core.Context) error {
	err := i.file.Close()
	return err
}

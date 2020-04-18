package file

import (
	"github.com/tk103331/logpipe/core"
	"os"
)

func init() {
	core.RegInput("file", func(conf core.InputConf) core.Input {
		return &FileInput{}
	})
}

type FileInput struct {
	core.BaseInput
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
				var source interface{} = string(data[0:n])
				if i.Codec != nil {
					source, e = i.Codec.Decode(source)
				}
				event := core.NewEvent("file", "localhost", source)
				ctx.Accept(event)
			}
		}
	}()
	return nil
}

func (i *FileInput) Stop() error {
	err := i.file.Close()
	return err
}

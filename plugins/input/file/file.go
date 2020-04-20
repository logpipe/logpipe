package file

import (
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/core"
	"github.com/tk103331/logpipe/plugin"
	"os"
)

func init() {
	plugin.RegInput(&FileInputBuilder{})
}

type FileInput struct {
	core.BaseInput
	Path string
	file *os.File
}

func (i *FileInput) Start(consumer func(event core.Event)) error {
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
				consumer(event)
			}
		}
	}()
	return nil
}

func (i *FileInput) Stop() error {
	err := i.file.Close()
	return err
}

type FileInputBuilder struct {
}

func (b *FileInputBuilder) Kind() string {
	return "file"
}

func (b *FileInputBuilder) Build(name string, codec core.Codec, spec config.Value) core.Input {
	return &FileInput{}
}

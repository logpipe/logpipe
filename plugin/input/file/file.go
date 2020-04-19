package file

import (
	"github.com/tk103331/logpipe/core"
	"os"
)

const INPUT_NAME = "file"

func init() {
	core.RegInput(INPUT_NAME, &FileInputBuilder{})
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

type FileInputConf struct {
	core.BaseInputConf
	Path string
}

type FileInputBuilder struct {
}

func (b *FileInputBuilder) NewConf() core.InputConf {
	return &FileInputConf{}
}

func (b *FileInputBuilder) Build(conf core.InputConf) core.Input {
	inputConf := conf.(*FileInputConf)
	return &FileInput{Path: inputConf.Path}
}

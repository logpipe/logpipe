package file

import (
	"bufio"
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
	codec core.Codec
	path  string
	delim byte
	file  *os.File
	stop  chan struct{}
}

func (i *FileInput) Start(consumer func(event core.Event)) error {
	file, err := os.OpenFile(i.path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	i.file = file
	go func() {
		reader := bufio.NewReader(i.file)
		for {
			select {
			case <-i.stop:
				break
			default:
			}
			line, e := reader.ReadString(i.delim)
			if e == nil {
				var source interface{} = line
				if i.codec != nil {
					event, e := i.codec.Decode(source)
					if e == nil {
						consumer(event)
					}
				} else {
					event := core.NewEvent(source)
					consumer(event)
				}
			}
		}
	}()
	return nil
}

func (i *FileInput) Stop() error {
	i.stop <- struct{}{}
	err := i.file.Close()
	return err
}

type FileInputBuilder struct {
}

func (b *FileInputBuilder) Kind() string {
	return "file"
}

func (b *FileInputBuilder) Build(name string, codec core.Codec, spec config.Value) core.Input {
	path := spec.GetString("path")
	delimValue := spec.Get("delim")
	var delim byte = '\n'
	if !delimValue.IsEmpty() {
		delim = byte(delimValue.Int())
	}
	return &FileInput{path: path, delim: delim, codec: codec}
}

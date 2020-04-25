package file

import (
	"errors"
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/plugin"
	"os"
)
import "github.com/logpipe/logpipe/core"

func init() {
	plugin.RegOutput(&FileOutputBuilder{})
}

type FileOutput struct {
	core.BaseOutput
	path  string
	delim byte
	file  *os.File
}

func (o *FileOutput) Start() error {
	stat, err := os.Stat(o.path)
	var file *os.File
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(o.path)
	} else if err != nil {
		return err
	} else if stat.IsDir() {
		return errors.New("the output file path is dir")
	} else {
		file, err = os.Open(o.path)
	}
	if err != nil {
		return err
	}
	o.file = file
	return nil
}

func (o *FileOutput) Stop() error {
	return o.file.Close()
}

func (o *FileOutput) Output(event core.Event) error {
	var err error
	data := event.Source()
	if o.Codec() != nil {
		data, err = o.Codec().Encode(event)
		if err != nil {
			return err
		}
	}
	if str, ok := data.(string); ok {
		_, err := o.file.Write([]byte(str))
		o.file.Write([]byte{o.delim})
		return err
	}
	return nil
}

type FileOutputBuilder struct {
}

func (f *FileOutputBuilder) Kind() string {
	return "file"
}

func (f *FileOutputBuilder) Build(name string, spec config.Value) core.Output {
	path := spec.GetString("path")
	delimValue := spec.Get("delim")
	var delim byte = '\n'
	if !delimValue.IsEmpty() {
		delim = byte(delimValue.Int())
	}
	return &FileOutput{path: path, delim: delim}
}

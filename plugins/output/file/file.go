package file

import "os"
import "github.com/tk103331/logpipe/core"

type FileOutput struct {
	core.BaseOutput
	Path string
	file *os.File
}

func (o *FileOutput) Start() error {
	file, err := os.Open(o.Path)
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
	data, err := o.Codec.Encode(event)
	if err != nil {
		return err
	}
	if str, ok := data.(string); ok {
		_, err := o.file.Write([]byte(str))
		return err
	}
	return nil
}

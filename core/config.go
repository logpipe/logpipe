package core

import (
	"fmt"
	"go.uber.org/config"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var appConf = AppConf{}
var pipeConf = make(map[string]PipeConf)

type AppConf struct {
	Path string
}

type Conf interface {
	Load(value *Value) error
}

type BaseConf struct {
	value *Value
}

func (c *BaseConf) Load(value *Value) error {
	c.value = value
	return nil
}

func (c *BaseConf) Value() *Value {
	return c.value
}

func LoadConf() error {

	if err := loadAppConf(); err != nil {
		return err
	}
	if err := loadPipeConf(); err != nil {
		return err
	}
	return nil
}
func loadAppConf() error {
	appConfPath := "conf/logpipe.yaml"

	stat, err := os.Stat(appConfPath)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		appConfPath = filepath.Join(appConfPath, "logpipe.yaml")
	}
	appConfPath, err = filepath.Abs(appConfPath)
	yaml, err := config.NewYAML(config.File(appConfPath))
	if err != nil {
		return err
	}
	pipePath := yaml.Get("path").String()
	if !filepath.IsAbs(pipePath) {
		pipePath = filepath.Join(filepath.Dir(appConfPath), pipePath)
	}
	appConf.Path = pipePath
	return nil
}

func loadPipeConf() error {
	path := appConf.Path

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fi := range dir {
		name := fi.Name()
		if !fi.IsDir() && (strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")) {
			absPath := filepath.Join(path, name)
			log.Printf("loading pipe conf: " + absPath)

			yaml, err := config.NewYAML(config.File(absPath), config.Permissive())
			if err != nil {
				fmt.Println(err)
				continue
			}
			value := yaml.Get("")
			conf := PipeConf{}
			err = conf.Load(&Value{value: value})
			if err != nil {
				fmt.Println(err)
				continue
			}
			pipeConf[absPath] = conf
		}
	}
	return nil
}

func GetPipeConf() map[string]PipeConf {
	return pipeConf
}

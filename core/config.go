package core

import (
	"gopkg.in/yaml.v3"
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

type KindConf interface {
	GetKind() string
}

type BaseKindConf struct {
	Kind string
}

func (c *BaseKindConf) GetKind() string {
	return c.Kind
}

type NameConf interface {
	GetName() string
}

type BaseNameConf struct {
	Name string
}

func (c *BaseNameConf) GetName() string {
	return c.Name
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
	if err != nil {
		return err
	}
	confFile, err := os.Open(appConfPath)
	if err != nil {
		return err
	}
	defer confFile.Close()
	decoder := yaml.NewDecoder(confFile)
	err = decoder.Decode(&appConf)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(appConf.Path) {
		appConf.Path = filepath.Join(filepath.Dir(appConfPath), appConf.Path)
	}
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

			err := readPipeConf(absPath)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func readPipeConf(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var node yaml.Node
	decoder.Decode(&node)
	if node.Kind == yaml.DocumentNode {
		node = *node.Content[0]
	}
	conf := PipeConf{}
	err = conf.Load(&Value{node: &node})
	if err != nil {
		return err
	}
	pipeConf[path] = conf
	return nil
}

func GetPipeConf() map[string]PipeConf {
	return pipeConf
}

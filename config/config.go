package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var appConf = AppConf{}

func init() {
	appConf.Pipes = make(map[string]PipeConf)
}

type AppConf struct {
	Path  string
	Pipes map[string]PipeConf
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
		if !fi.IsDir() && (strings.HasSuffix(name, ".yaml")) {
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
	node.Decode(&conf)
	if err != nil {
		return err
	}
	appConf.Pipes[path] = conf
	return nil
}

func GetAppConf() AppConf {
	return appConf
}

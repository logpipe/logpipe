package config

import (
	"github.com/logpipe/logpipe/log"
	"io/ioutil"
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
	Log   LogConf
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
	value, err := NewValue(confFile)
	if err != nil {
		return err
	}
	err = value.Parse(&appConf)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(appConf.Path) {
		appConf.Path = filepath.Join(filepath.Dir(appConfPath), appConf.Path)
	}
	err = initLog()
	if err != nil {
		return err
	}
	return nil
}

func initLog() error {
	if appConf.Log.Level == "" {
		appConf.Log.Level = DEFAULT_LOG_LEVEL
	}
	if appConf.Log.Path == "" {
		appConf.Log.Path = DEFAULT_LOG_PATH
	}
	stat, err := os.Stat(appConf.Log.Path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		appConf.Log.Path = filepath.Join(appConf.Log.Path, DEFAULT_LOG_NAME)
	}
	if !filepath.IsAbs(appConf.Log.Path) {
		appConf.Log.Path, err = filepath.Abs(appConf.Log.Path)
	}

	log.InitAppLogger(appConf.Log.Path, appConf.Log.Level)
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
			log.Info("loading example conf: " + absPath)

			err := readPipeConf(absPath)
			if err != nil {
				log.Error("loading %s error: %s", absPath, err)
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
	value, err := NewValue(file)
	if err != nil {
		return err
	}
	conf := PipeConf{}
	err = value.Parse(&conf)
	if err != nil {
		return err
	}
	appConf.Pipes[path] = conf
	return nil
}

func GetAppConf() AppConf {
	return appConf
}

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
	Pipe struct {
		Path string
	}
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
	log.Info("config file load finished")
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
	value, err := readValue(appConfPath)
	if err != nil {
		return err
	}
	err = value.Parse(&appConf)
	if err != nil {
		return err
	}
	if !filepath.IsAbs(appConf.Pipe.Path) {
		appConf.Pipe.Path = filepath.Join(filepath.Dir(appConfPath), appConf.Pipe.Path)
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
	log.Info("logging [%v] into %v", appConf.Log.Level, appConf.Log.Path)
	return nil
}

func loadPipeConf() error {
	path := appConf.Pipe.Path
	log.Info("scanning pipe conf file in %v", appConf.Pipe.Path)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fi := range dir {
		name := fi.Name()
		if !fi.IsDir() && (strings.HasSuffix(name, ".yaml")) {
			absPath := filepath.Join(path, name)
			log.Info("loading pipe conf: %v", absPath)

			err := readPipeConf(absPath)
			if err != nil {
				log.Error("loading %v error: %v", absPath, err.Error())
			}
		}
	}
	return nil
}

func readPipeConf(path string) error {
	value, err := readValue(path)
	if err != nil {
		return err
	}
	conf := PipeConf{file: path}
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

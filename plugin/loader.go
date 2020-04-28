// +build linux
package plugin

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

const plugin_dir = "plugins/"
const plugin_ext = ".so"

func LoadPlugins() error {
	appConf := config.GetAppConf()
	if appConf.Plugins == "" {
		appConf.Plugins = plugin_dir
	}
	stat, err := os.Stat(appConf.Plugins)
	if err == os.ErrNotExist {
		log.Error("plugins dir not exist: %v", appConf.Plugins)
		return err
	} else if err != nil {
		log.Error("scan plugins dir [%v] error: %v", appConf.Plugins, err.Error())
	}
	if !filepath.IsAbs(appConf.Plugins) {
		appConf.Plugins, err = filepath.Abs(appConf.Plugins)
		if err != nil {
			return err
		}
	}
	if stat.IsDir() {
		dir, err := ioutil.ReadDir(appConf.Plugins)
		if err != nil {
			return err
		}
		for _, f := range dir {
			if strings.HasSuffix(f.Name(), plugin_ext) {
				load(filepath.Join(stat.Name(), f.Name()))
			}
		}
	} else {
		load(appConf.Plugins)
	}
	return nil
}

func load(path string) {
	_, err := plugin.Open(path)
	if err != nil {
		log.Error("loading plugin [%v] error: %v", path, err.Error())
	}
}

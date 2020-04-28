// +build linux

package main

import (
	"github.com/logpipe/logpipe/engine"
	"github.com/logpipe/logpipe/log"
	"github.com/logpipe/logpipe/plugin"
	"github.com/logpipe/logpipe/config"
)

func main() {

	if err := config.LoadConf(); err != nil {
		log.Fatal("loading config file error: %v", err.Error())
	}

	if err := plugin.LoadPlugins(); err != nil {
		log.Fatal("loading plugins error: %v", err.Error())
	}

	if err := engine.Init(); err != nil {
		log.Fatal("initialize logpipe engine error: %v", err.Error())
	}
	if err := engine.Start(); err != nil {
		log.Fatal("starting logpipe engine error: %v", err.Error())
	}

	engine.Wait()
}

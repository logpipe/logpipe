package main

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/engine"
	_ "github.com/logpipe/logpipe/plugins"
	"log"
)

func main() {
	err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	if err := engine.Init(); err != nil {
		log.Fatal(err)
	}
	if err := engine.Start(); err != nil {
		log.Fatal(err)
	}

	engine.Wait()
}

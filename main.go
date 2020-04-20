package main

import (
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/engine"
	_ "github.com/tk103331/logpipe/plugins"
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

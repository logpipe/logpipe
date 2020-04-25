package main

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/engine"
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
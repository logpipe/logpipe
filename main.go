package main

import (
	"github.com/tk103331/logpipe/config"
	"github.com/tk103331/logpipe/engine"
	"log"
)

func main() {
	if err := config.LoadConf(); err != nil {
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

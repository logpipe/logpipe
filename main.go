package main

import (
	"github.com/tk103331/logpipe/core"
	_ "github.com/tk103331/logpipe/plugin"
	"log"
)

func main() {
	if err := core.LoadConf(); err != nil {
		log.Fatal(err)
	}

	if err := core.Init(); err != nil {
		log.Fatal(err)
	}
	if err := core.Start(); err != nil {
		log.Fatal(err)
	}

	core.Wait()
}

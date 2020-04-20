package engine

import (
	"github.com/tk103331/logpipe/config"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var pipes = make(map[string]*Pipe)
var done = make(chan int)

func Init() error {
	pipeConf := config.GetAppConf().Pipes
	for name, conf := range pipeConf {
		pipe := &Pipe{}
		pipe.Init(conf)
		pipes[name] = pipe
	}
	return nil
}

func Start() error {
	wg := sync.WaitGroup{}
	wg.Add(len(pipes))
	for name, pipe := range pipes {
		go func(name string, p *Pipe) {
			log.Println("starting pipe: " + name)
			p.Start()
			wg.Done()
		}(name, pipe)
	}
	wg.Wait()
	monitor()
	return nil
}
func Stop() {
	wg := sync.WaitGroup{}
	wg.Add(len(pipes))
	for name, pipe := range pipes {
		go func(name string, p *Pipe) {
			log.Println("stopping pipe: " + name)
			p.Stop()
			wg.Done()
		}(name, pipe)
	}
	wg.Wait()
	done <- 0
}

func Wait() {
	<-done
}

func monitor() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-signals
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			Stop()
		}
	}
}

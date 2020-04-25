package engine

import (
	"github.com/logpipe/logpipe/config"
	"github.com/logpipe/logpipe/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var pipes = make(map[string]*pipe)
var done = make(chan struct{})

func Init() error {
	pipeConf := config.GetAppConf().Pipes
	for name, conf := range pipeConf {
		pipe := &pipe{}
		pipe.Init(conf)
		pipes[name] = pipe
	}
	return nil
}

func Start() error {
	wg := sync.WaitGroup{}
	wg.Add(len(pipes))
	for name, p := range pipes {
		go func(name string, p *pipe) {
			log.Info("starting pipe [%s]", name)
			p.Start()
			wg.Done()
		}(name, p)
	}
	wg.Wait()
	go monitor()
	return nil
}
func Stop() {
	wg := sync.WaitGroup{}
	wg.Add(len(pipes))
	for name, p := range pipes {
		go func(name string, p *pipe) {
			log.Info("stopping pipe [%s]", name)
			p.Stop()
			wg.Done()
		}(name, p)
	}
	wg.Wait()
	done <- struct{}{}
}

func Wait() {
	<-done
}

func monitor() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	for {
		sig := <-signals
		if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGKILL {
			Stop()
		}
	}
}

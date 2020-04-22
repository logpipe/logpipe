package core

import (
	"fmt"
	"log"
	"time"
)

type Context struct {
	pipe   string
	name   string
	kind   string
	plugin string
	logger *log.Logger
}

func NewContext(pipe string, name string, kind string, logger *log.Logger) *Context {
	plugin := fmt.Sprintf("[%s-%s-%s]", pipe, kind, name)
	return &Context{pipe: pipe, name: name, kind: kind, plugin: plugin, logger: logger}
}

func (c *Context) log(level string, format string, values ...interface{}) {
	timestamp := time.Now().String()
	prefix := fmt.Sprintf("[%s] [%s] [%s] ", timestamp, level, c.plugin)
	message := fmt.Sprintf(format, values)
	c.logger.Println(prefix + message)
}

func (c *Context) Info(format string, values ...interface{}) {
	c.log("[INFO]", format, values)
}

func (c *Context) Warn(format string, values ...interface{}) {
	c.log("[WARN]", format, values)
}

func (c *Context) Error(format string, values ...interface{}) {
	c.log("[ERROR]", format, values)
}

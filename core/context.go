package core

import (
	"fmt"
)

type ContextContainer interface {
	Context() *Context
	SetContext(ctx *Context)
}

type Context struct {
	pipe   string
	name   string
	kind   string
	plugin string
	logger *Logger
}

func NewContext(pipe string, name string, kind string, logger *Logger) *Context {
	plugin := fmt.Sprintf(" [%s-%s-%s] ", pipe, kind, name)
	return &Context{pipe: pipe, name: name, kind: kind, plugin: plugin, logger: logger}
}

func (c *Context) Debug(format string, values ...interface{}) {
	c.logger.Debug(c.plugin+format, values...)
}

func (c *Context) Info(format string, values ...interface{}) {
	c.logger.Info(c.plugin+format, values...)
}

func (c *Context) Warn(format string, values ...interface{}) {
	c.logger.Warn(c.plugin+format, values...)
}

func (c *Context) Error(format string, values ...interface{}) {
	c.logger.Error(c.plugin+format, values...)
}

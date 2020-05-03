package core

import (
	"fmt"
	"github.com/logpipe/logpipe/log"
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
	logger *log.Logger
	vars   map[string]interface{}
}

func NewContext(pipe string, name string, kind string, logger *log.Logger, vars map[string]interface{}) *Context {
	plugin := fmt.Sprintf(" [%s-%s-%s] ", pipe, kind, name)
	return &Context{pipe: pipe, name: name, kind: kind, plugin: plugin, logger: logger, vars: vars}
}

func (c *Context) GetVar(name string) interface{} {
	return c.vars[name]
}

func (c *Context) SetVar(name string, value interface{}) {
	c.vars[name] = value
}

func (c *Context) HasVar(name string) bool {
	_, ok := c.vars[name]
	return ok
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

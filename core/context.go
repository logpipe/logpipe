package core

type Context struct {
	pipe *Pipe
}

func (c *Context) Accept(event Event) {
	go c.pipe.Input(event)
}

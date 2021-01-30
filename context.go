package cli

import "io"

var ctx *Context

type Context struct {
	Output io.Writer
	cmd    *Command
	subcmd *Subcommand
}

func (ctx *Context) Subcommand() string {
	return ctx.subcmd.Name
}

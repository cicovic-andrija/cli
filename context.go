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

func (ctx *Context) StringOption(name string) (s string) {
	strOption := ctx.cmd.findOption(name)
	if strOption == nil {
		return ""
	}
	if !ctx.subcmd.isvalid(name) {
		return ""
	}
	return strOption.Value().(string)
}

func (ctx *Context) BoolOption(name string) bool {
	return false
}

func (ctx *Context) IntOption(name string) int {
	return 0
}

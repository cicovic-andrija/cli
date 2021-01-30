package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"
)

const (
	stateInit = iota
	stateCmd
	stateSubcmd
)

var (
	state int
	cmd   *Command
)

func Execute(c *Command) {
	panicf := func(msg string) {
		panic("cli.Execute: " + msg)
	}

	if c == nil {
		panicf("nil argument")
	}

	// Initialize common variables.
	state = stateInit
	ctx = &Context{Output: os.Stdout}
	cmd = c
	err := cmd.init()
	if err != nil {
		panicf(err.Error()) // cmd configured incorrectly
	}

	// Parse templates.

	cmdTmpl = template.Must(
		template.New(commandHelpTemplateName).
			Parse(commandHelpTemplate),
	)

	subcmdTmplFuncMap := template.FuncMap{
		"CmdName":    _cmdName,
		"FindOption": _findOption,
	}

	subcmdTmpl = template.Must(
		template.New(subcommandHelpTemplateName).
			Funcs(subcmdTmplFuncMap).
			Parse(subcommandHelpTemplate),
	)

	state = stateCmd
	ctx.cmd = cmd

	// Top level command call, but only -h/-help/--help is supported, "thanks" to the flag package.
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		flag.Usage = func() { usageShort() }
		flag.Bool(HelpOptionName, false, "")      // need these two to override the default -h/-help/--help
		flag.Bool(HelpOptionShortName, false, "") // behavior, it doesn't matter if it's set
		flag.Parse()

		printHelp(cmdTmpl, cmd) // in any case, print usage message
		os.Exit(0)
	}

	// Subcommand call, possibly with options and positional arguments.

	subcmd := cmd.findSubcommand(os.Args[1])
	if subcmd == nil {
		failf("subcommand not defined: %q\n", os.Args[1])
	}

	state = stateSubcmd
	ctx.subcmd = subcmd

	// Custom flag set is needed here.
	flagSet := flag.NewFlagSet(subcmd.Name, flag.ExitOnError)
	flagSet.Usage = func() { usageShort() }

	err = registeroptions(subcmd, flagSet)
	if err != nil {
		panicf(err.Error()) // cmd configured incorrectly
	}

	// Parse command line arguments for the subcommand.
	flagSet.Parse(os.Args[2:])

	// The -h/-help/--help option handling is built-in.
	if helpOption.set {
		printHelp(subcmdTmpl, subcmd)
		os.Exit(0)
	}

	if subcmd.Handler == nil {
		// cmd configured incorrectly
		panicf(fmt.Sprintf("no handler definition for subcommand %q", subcmd.Name))
	}

	// Finally, execute the subcommand handler.
	subcmd.Handler(ctx)
}

func failf(format string, a ...interface{}) {
	fmt.Fprintf(ctx.Output, format, a...)
	usageShort()
	os.Exit(1)
}

func _findOption(name string) interface{} {
	return cmd.findOption(name)
}

func _cmdName() string {
	return cmd.Name
}

func printHelp(tmpl *template.Template, data interface{}) {
	w := tabwriter.NewWriter(ctx.Output, 8, 4, 2, ' ', 0)
	tmpl.Execute(w, data)
	w.Flush()
}

func usageShort() {
	param2 := ""
	if state >= stateSubcmd {
		param2 = " " + ctx.subcmd.Name
	}
	fmt.Fprintf(ctx.Output, "Run '%s%s --help' to get help\n", cmd.Name, param2)
}

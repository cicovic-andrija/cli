package cli

import "text/template"

var (
	cmdTmpl    *template.Template
	subcmdTmpl *template.Template
)

const (
	commandHelpTemplateName    = "command"
	subcommandHelpTemplateName = "subcommand"

	commandHelpTemplate = `{{.Description}}

USAGE
	{{.Name}} SUBCOMMAND [ARGUMENTS]

SUBCOMMAND{{range .Subcommands}}
	{{.Name}}	{{.Description}}	{{end}}

ARGUMENTS
	Run '{{.Name}} SUBCOMMAND --help' for individual subcommand usage

VERSION
	{{.Name}} {{.Version}}

AUTHOR
	{{.Author}}
`

	subcommandHelpTemplate = `NAME
	{{CmdName}} {{.Name}}

DESCRIPTION
	{{.Description}}

USAGE
	{{CmdName}} {{.Name}}{{if .ValidOptions}} [OPTIONS]{{end}} {{.PosArgsUsage}}{{if .ValidOptions}}

OPTIONS{{range .ValidOptions}}
	{{with $opt := FindOption .}}--{{$opt.Name}}{{if $opt.ShortName}},-{{$opt.ShortName}}{{end}}	[{{if $.Requires $opt.Name}}required{{else}}optional{{end}}]	{{$opt.Description}} (default: {{$opt.Default}})	{{end}}{{end}}{{end}}{{if .Examples}}

EXAMPLES{{range .Examples}}
	$ {{CmdName}} {{$.Name}} {{.Arguments}}
	~ {{.Explanation}}
	{{end}}{{end}}
`
)

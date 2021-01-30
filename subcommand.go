package cli

import "fmt"

const (
	NoArguments = ""
)

type (
	Command struct {
		Name        string
		Description string
		Author      string
		Version     string
		Subcommands []Subcommand
		Options     []Option
		optionMap   map[string]int
	}

	Subcommand struct {
		Name            string
		Description     string
		PosArgsUsage    string
		ValidOptions    []string
		RequiredOptions []string
		Examples        []Example
		Handler         HandlerFunc
	}

	HandlerFunc func(*Context)

	Example struct {
		Arguments   string
		Explanation string
	}
)

func (c *Command) init() error {
	c.optionMap = make(map[string]int)
	for i, o := range c.Options {
		c.optionMap[o.Name()] = i
	}

	// Option -h/-help/--help is predefined.
	if _, found := c.optionMap[helpOption.name]; found {
		return fmt.Errorf("duplicate option definition, option %q is predefined", helpOption.name)
	}
	c.Options = append(c.Options, helpOption)
	c.optionMap[helpOption.name] = len(c.Options) - 1

	return nil
}

func (c *Command) findOption(name string) Option {
	i, ok := c.optionMap[name]
	if !ok {
		return nil
	}
	return c.Options[i]
}

func (c *Command) findSubcommand(name string) *Subcommand {
	for i, s := range c.Subcommands {
		if s.Name == name {
			return &c.Subcommands[i]
		}
	}
	return nil
}

func (s *Subcommand) isvalid(option string) bool {
	return find(s.ValidOptions, option) != notFound
}

func (s *Subcommand) Requires(option string) bool {
	return find(s.RequiredOptions, option) != notFound
}

package cli

import (
	"flag"
	"fmt"
)

const (
	HelpOptionName      = "help"
	HelpOptionShortName = "h"
)

type (
	Option interface {
		Name() string
		ShortName() string
		Description() string
		Default() string
		Value() interface{}
		register(*flag.FlagSet)
	}

	optionBase struct {
		name        string // Must not be empty
		shortName   string // Must be empty string or have a len 1
		description string // Can be empty
	}
)

var (
	helpOption = &boolOption{
		optionBase: optionBase{
			name:        HelpOptionName,
			shortName:   HelpOptionShortName,
			description: "Print help",
		},
		set: false,
	}
)

func (o optionBase) Name() string {
	return o.name
}

func (o optionBase) ShortName() string {
	return o.shortName
}

func (o optionBase) Description() string {
	return o.description
}

func registeroptions(subcmd *Subcommand, flagSet *flag.FlagSet) error {
	// Option -h/-help/--help is predefined and not required.
	subcmd.ValidOptions = append(subcmd.ValidOptions, helpOption.name)

	for _, name := range subcmd.ValidOptions {
		option := cmd.findOption(name)
		if option == nil {
			return fmt.Errorf("no definition for --%s", name)
		}
		option.register(flagSet)
	}

	return nil
}

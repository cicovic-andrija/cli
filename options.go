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

	stringOption struct {
		optionBase
		defaultValue string
		value        string
	}

	boolOption struct {
		optionBase
		set bool
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

func NewStringOption(name string, shortName string, description string, defaultValue string) Option {
	return &stringOption{
		optionBase: optionBase{
			name:        name,
			shortName:   shortName,
			description: description,
		},
		defaultValue: defaultValue,
		value:        defaultValue,
	}
}

func (s *stringOption) register(flagSet *flag.FlagSet) {
	flagSet.StringVar(&s.value, s.name, s.value, s.description)
	if s.shortName != "" {
		flagSet.StringVar(&s.value, s.shortName, s.value, s.description)
	}
}

func (s *stringOption) Default() string {
	return fmt.Sprintf("%q", s.defaultValue)
}

func (s *stringOption) Value() interface{} {
	return s.value
}

func NewBoolOption(name string, shortName string, description string) Option {
	return &boolOption{
		optionBase: optionBase{
			name:        name,
			shortName:   shortName,
			description: description,
		},
		set: false,
	}
}

func (b *boolOption) register(flagSet *flag.FlagSet) {
	flagSet.BoolVar(&b.set, b.name, false, b.description)
	if b.shortName != "" {
		flagSet.BoolVar(&b.set, b.shortName, false, b.description)
	}
}

func (b *boolOption) Default() string {
	return "false"
}

func (b *boolOption) Value() interface{} {
	return b.set
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

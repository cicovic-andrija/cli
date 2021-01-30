package cli

import (
	"flag"
	"fmt"
)

type stringOption struct {
	optionBase
	defaultValue string
	value        string
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

package cli

import "flag"

type boolOption struct {
	optionBase
	set bool
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

package flagutil

import (
	"log/slog"

	"github.com/spf13/pflag"
)

func MustGetString(flagSet *pflag.FlagSet, name string) string {
	v, err := flagSet.GetString(name)
	if err != nil {
		slog.Error(err.Error())
		return ""
	}

	return v
}

func MustGetBool(flagSet *pflag.FlagSet, name string) bool {
	v, err := flagSet.GetBool(name)
	if err != nil {
		slog.Error(err.Error())
		return false
	}

	return v
}

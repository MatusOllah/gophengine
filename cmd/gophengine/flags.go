package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var (
	helpFlag          bool
	extractAssetsFlag bool
	configFlag        string
	progressFlag      string
	logLevelFlag      string
	justInitFlag      bool
	portableFlag      bool
	forceLocaleFlag   string
)

var flagSet *pflag.FlagSet

func initFlags() error {
	flagSet = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// help flag
	flagSet.BoolVarP(&helpFlag, "help", "h", false, "Shows this help message")

	flagSet.BoolVar(&extractAssetsFlag, "extract-assets", false, "Extract embedded assets")
	flagSet.StringVar(&configFlag, "config", "", "Path to config.gecfg config file")
	flagSet.StringVar(&progressFlag, "progress", "", "Path to progress.gecfg progress file")
	flagSet.StringVar(&logLevelFlag, "log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	flagSet.BoolVar(&justInitFlag, "just-init", false, "Initialize game and exit")
	flagSet.BoolVar(&portableFlag, "portable", false, "Save everything in the current directory (aka portable mode)")
	flagSet.StringVar(&forceLocaleFlag, "force-locale", "", "Force a specific locale (used for testing locales)")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if helpFlag {
		fmt.Printf("GophEngine is a Go implementation of Friday Night Funkin' with improvements.\n\n")
		fmt.Printf("Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		fmt.Print(flagSet.FlagUsages())
		os.Exit(0)
	}

	return nil
}

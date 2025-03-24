package main

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

var flagSet *pflag.FlagSet = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

var (
	helpFlag = flagSet.BoolP("help", "h", false, "Shows this help message")

	extractAssetsFlag = flagSet.String("extract-assets", "", "Extract embedded assets")
	configFlag        = flagSet.String("config", "", "Path to config.gecfg config file")
	progressFlag      = flagSet.String("progress", "", "Path to progress.gecfg progress file")
	logLevelFlag      = flagSet.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	justInitFlag      = flagSet.Bool("just-init", false, "Initialize game and exit")
	portableFlag      = flagSet.Bool("portable", false, "Save everything in the current directory (aka portable mode)")
	forceLocaleFlag   = flagSet.String("force-locale", "", "Force a specific locale (used for testing locales)")
	cpuProfileFlag    = flagSet.String("cpu-profile", "", "Write CPU profile to file")
	memProfileFlag    = flagSet.String("mem-profile", "", "Write memory profile to file")
	httpProfileFlag   = flagSet.Bool("http-profile", false, "Serve profiling data via HTTP server on port :6060 (see https://pkg.go.dev/net/http/pprof)")
)

func initFlags() error {
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if *helpFlag {
		fmt.Printf("GophEngine is a Go port of Friday Night Funkin' with improvements.\n\n")
		fmt.Printf("Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		fmt.Print(flagSet.FlagUsages())
		os.Exit(0)
	}

	return nil
}

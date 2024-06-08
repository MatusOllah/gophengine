package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/spf13/pflag"
)

var flagSet *pflag.FlagSet

func initFlags() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	flagSet = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// help flag
	flagSet.BoolP("help", "h", false, "Shows this help message")

	flagSet.Bool("extract-assets", false, "Extract embedded assets")
	flagSet.String("config", filepath.Join(configDir, "GophEngine/config.gecfg"), "Path to config.gecfg config file")
	flagSet.String("progress", filepath.Join(configDir, "GophEngine/progress.gecfg"), "Path to progress.gecfg progress file")
	flagSet.BoolP("gui", "g", true, "Enable GUI & dialogs")
	flagSet.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	flagSet.Bool("just-init", false, "Initialize game and exit")
	flagSet.Bool("portable", false, "Save everything in the current directory (aka portable mode)")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if flagutil.MustGetBool(flagSet, "help") {
		fmt.Printf("GophEngine is a Go implementation of Friday Night Funkin' with improvments.\n\n")
		fmt.Printf("Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		fmt.Print(flagSet.FlagUsages())
		os.Exit(0)
	}

	return nil
}

package gophengine

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MatusOllah/gophengine/internal/flagutil"
	"github.com/spf13/pflag"
)

var FlagSet *pflag.FlagSet

func InitFlags() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	flagSet := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	// help flag
	flagSet.BoolP("help", "h", false, "Shows this help message")

	flagSet.Bool("extract-assets", false, "Extract embedded assets")
	flagSet.String("config", filepath.Join(configDir, "GophEngine/config.gecfg"), "Path to config.gecfg config file")
	flagSet.String("progress", filepath.Join(configDir, "GophEngine/progress.gecfg"), "Path to progress.gecfg progress file")
	flagSet.Bool("vsync", false, "Enable VSync")
	flagSet.BoolP("gui", "g", true, "Enable GUI & dialogs")
	flagSet.String("log-level", "info", "Log level (\"debug\", \"info\", \"warn\", \"error\")")
	flagSet.Bool("just-init", false, "Initialize game and exit")

	// config flags
	flagSet.Bool("config-load-defaults", false, "Wipe config and load defaults")

	flagSet.StringSlice("config-string", []string{}, "Override a string value in config")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	if flagutil.MustGetBool(flagSet, "help") {
		fmt.Printf("GophEngine is a Go implementation of Friday Night Funkin' with improvments.\n\n")
		fmt.Printf("Usage: %s [OPTIONS]\n\nOptions:\n", os.Args[0])
		fmt.Print(flagSet.FlagUsages())
		fmt.Printf("\nConfig Override Usage: --config-[TYPE]=\"[KEY]:[VALUE]\"\n")
		os.Exit(0)
	}

	FlagSet = flagSet

	return nil
}

package main

import (
	"fmt"
	"os"

	"github.com/joaogiacometti/aurgo/aur"
	"github.com/joaogiacometti/aurgo/helpers"
)

func main() {
	config, args := helpers.ParseFlags()

	if err := run(config, args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(config helpers.FlagOptions, args []string) error {
	if config.ShowHelp {
		helpers.PrintUsage()
		return nil
	}

	if (config.Install && config.Remove) || (config.Install && config.Update) || (config.Remove && config.Update) {
		helpers.PrintUsage()
		return fmt.Errorf("cannot use -S, -R, and -U together")
	}

	if config.Update {
		if len(args) > 0 {
			return fmt.Errorf("-U does not take package arguments")
		}
		return aur.UpdateAll()
	}

	if len(args) == 0 {
		helpers.PrintUsage()
		return fmt.Errorf("no package name provided")
	}

	switch {
	case config.Search:
		for _, pkg := range args {
			if err := aur.SearchPackage(pkg); err != nil {
				return fmt.Errorf("search failed for %s: %w", pkg, err)
			}
		}
	case config.Install:
		for _, pkg := range args {
			if err := aur.InstallPackage(pkg); err != nil {
				return fmt.Errorf("install failed for %s: %w", pkg, err)
			}
		}
	case config.Remove:
		for _, pkg := range args {
			if err := aur.RemovePackage(pkg); err != nil {
				return fmt.Errorf("remove failed for %s: %w", pkg, err)
			}
		}
	default:
		helpers.PrintUsage()
		return fmt.Errorf("no operation specified")
	}

	return nil
}

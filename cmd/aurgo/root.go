package aurgo

import (
	"fmt"
	"os"

	"github.com/joaogiacometti/aurgo/internal/aur"
	"github.com/joaogiacometti/aurgo/internal/helpers"
)

func Execute() {
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

	if config.Install && config.Remove {
		helpers.PrintUsage()
		return fmt.Errorf("-S and -R cannot be used together")
	}

	if len(args) == 0 {
		helpers.PrintUsage()
		return fmt.Errorf("no package name provided")
	}

	for _, pkg := range args {
		switch {
		case config.Search:
			if err := aur.SearchPackage(pkg); err != nil {
				return fmt.Errorf("search failed for %s: %w", pkg, err)
			}
		case config.Install:
			if err := aur.InstallPackage(pkg); err != nil {
				return fmt.Errorf("install failed for %s: %w", pkg, err)
			}
		case config.Remove:
			if err := aur.RemovePackage(pkg); err != nil {
				return fmt.Errorf("remove failed for %s: %w", pkg, err)
			}
		default:
			helpers.PrintUsage()
			return fmt.Errorf("no operation specified, use -Ss, -S, or -R")
		}
	}

	return nil
}

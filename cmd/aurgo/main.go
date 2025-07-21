package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joaogiacometti/aurgo/internal/aur"
)

type Config struct {
	Install  bool
	ShowHelp bool
	Remove   bool
}

func main() {
	var config Config

	flag.BoolVar(&config.Install, "S", false, "Install package(s)")
	flag.BoolVar(&config.ShowHelp, "h", false, "Show help")
	flag.BoolVar(&config.Remove, "R", false, "Remove package(s)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <package-name>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -h             Show this help\n\n")
		fmt.Fprintf(os.Stderr, "  -S             Install package(s)\n")
		fmt.Fprintf(os.Stderr, "  -R             Remove package(s)\n")
		fmt.Fprintf(os.Stderr, "Examples:\n")
		fmt.Fprintf(os.Stderr, "  %s -h                 # Show help\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -S package-name     # Install package\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -R package-name     # Remove package\n", os.Args[0])
	}

	flag.Parse()
	args := flag.Args()

	if config.ShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	if config.Install && config.Remove {
		fmt.Println("Error: -S and -R cannot be used together")
		flag.Usage()
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println("Error: No package name provided")
		flag.Usage()
		os.Exit(1)
	}

	if config.Install {
		packageName := args[0]

		if err := aur.InstallPackage(packageName); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if config.Remove {
		packageName := args[0]

		if err := aur.RemovePackage(packageName); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		os.Exit(0)
	}
}

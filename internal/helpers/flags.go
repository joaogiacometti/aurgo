package helpers

import (
	"flag"
	"fmt"
	"os"
)

type FlagOptions struct {
	ShowHelp bool
	Search   bool
	Install  bool
	Remove   bool
	Update   bool
}

func ParseFlags() (FlagOptions, []string) {
	var config FlagOptions

	flag.BoolVar(&config.ShowHelp, "h", false, "Show help")
	flag.BoolVar(&config.Search, "Ss", false, "Search for package(s)")
	flag.BoolVar(&config.Install, "S", false, "Install package(s)")
	flag.BoolVar(&config.Remove, "R", false, "Remove package(s)")
	flag.BoolVar(&config.Update, "U", false, "Update package(s)")

	flag.Usage = PrintUsage
	flag.Parse()

	return config, flag.Args()
}

func PrintUsage() {
	fmt.Fprintf(os.Stderr, "aurgo - AUR helper written in Go\n")
	fmt.Fprintf(os.Stderr, "Usage: aurgo [options] <package-name...>\n\n")
	fmt.Fprintf(os.Stderr, "Options:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  aurgo -h                 # Show help\n")
	fmt.Fprintf(os.Stderr, "  aurgo -Ss package        # Search\n")
	fmt.Fprintf(os.Stderr, "  aurgo -S package         # Install multiple packages\n")
	fmt.Fprintf(os.Stderr, "  aurgo -R package         # Remove a package\n")
	fmt.Fprintf(os.Stderr, "  aurgo -U package         # Update a package\n")
}

package aur

import (
	"fmt"
)

func SearchPackage(query string) error {
	aurResponse, err := SearchPackages(query)
	if err != nil {
		return fmt.Errorf("searching packages: %w", err)
	}

	for _, pkg := range aurResponse.Results {
		fmt.Printf("Package: %s\n", pkg.Name)
		fmt.Printf("Description: %s\n", pkg.Description)
		fmt.Printf("Version: %s\n", pkg.Version)
		fmt.Printf("Maintainer: %s\n", pkg.Maintainer)
		fmt.Printf("Upstream URL: %s\n", pkg.UpstreamURL)
		fmt.Printf("Votes: %d\n", pkg.NumVotes)
		fmt.Println()
	}

	return nil
}

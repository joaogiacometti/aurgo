package aur

import (
	"fmt"

	"github.com/joaogiacometti/aurgo/helpers"
)

func UpdateAll() error {
	versions, err := helpers.ReadVersions()
	if err != nil {
		return fmt.Errorf("failed to read installed packages: %w", err)
	}

	updated := false

	for name, currVersion := range versions {
		result, err := FindPackage(name)
		if err != nil {
			fmt.Printf("✗ Failed to fetch info for %s: %v\n", name, err)
			continue
		}

		if result.Version != currVersion {
			fmt.Printf("\nUpdating %s from %s to %s\n\n", name, currVersion, result.Version)

			if err := InstallPackage(name); err != nil {
				fmt.Printf("✗ Failed to update %s: %v\n", name, err)
				continue
			}

			versions[name] = result.Version
			updated = true
		}
	}

	if updated {
		helpers.WriteVersions(versions)
	} else {
		fmt.Println("No packages needed updating.")
	}

	return nil
}

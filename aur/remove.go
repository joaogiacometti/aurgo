package aur

import (
	"fmt"
	"os"

	"github.com/joaogiacometti/aurgo/helpers"
)

func RemovePackage(pkgName string) error {
	if err := helpers.RemoveWithPacman(pkgName); err != nil {
		return fmt.Errorf("pacman failed to remove package '%s': %w", pkgName, err)
	}

	if err := helpers.RemoveVersion(pkgName); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not update installed list: %v\n", err)
	}

	fmt.Printf("Successfully removed package '%s'\n", pkgName)
	return nil
}

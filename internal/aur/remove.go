package aur

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/helpers"
)

func RemovePackage(pkgName string) error {
	packagePath := filepath.Join(config.DataDir, pkgName)

	ok, err := helpers.IsDir(packagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("package '%s' not found in cache", pkgName)
		}
		return fmt.Errorf("failed to stat package path '%s': %w", packagePath, err)
	}
	if !ok {
		return fmt.Errorf("'%s' is not a directory", packagePath)
	}

	if err := helpers.RemoveWithPacman(pkgName); err != nil {
		return fmt.Errorf("pacman failed to remove package '%s': %w", pkgName, err)
	}

	if err := helpers.RemoveVersion(pkgName); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not update installed list: %v\n", err)
	}

	fmt.Printf("Successfully removed package '%s'\n", pkgName)
	return nil
}

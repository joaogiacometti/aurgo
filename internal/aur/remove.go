package aur

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/helpers"
)

func RemovePackage(pkgName string) error {
	packagePath := filepath.Join(config.CacheDir, pkgName)

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

	if err := os.RemoveAll(packagePath); err != nil {
		return fmt.Errorf("failed to remove package directory '%s': %w", packagePath, err)
	}

	fmt.Printf("Successfully removed package '%s'\n", pkgName)
	return nil
}

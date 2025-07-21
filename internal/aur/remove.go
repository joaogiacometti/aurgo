package aur

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joaogiacometti/aurgo/internal/config"
)

func RemovePackage(pkgName string) error {
	packagePath := filepath.Join(config.CacheDir, pkgName)

	info, err := os.Stat(packagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("package '%s' not found in cache", pkgName)
		}
		return fmt.Errorf("failed to stat package path '%s': %w", packagePath, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("'%s' is not a directory", packagePath)
	}

	cmd := exec.Command("sudo", "pacman", "-Rns", pkgName)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("pacman failed to remove package '%s': %w", pkgName, err)
	}

	if err := os.RemoveAll(packagePath); err != nil {
		return fmt.Errorf("failed to remove package directory '%s': %w", packagePath, err)
	}

	fmt.Printf("Successfully removed package '%s'\n", pkgName)
	return nil
}

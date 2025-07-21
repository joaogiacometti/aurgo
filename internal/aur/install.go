package aur

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/helpers"
)

func InstallPackage(pkgName string) error {
	pkg, err := FindPackage(pkgName)
	if err != nil {
		return fmt.Errorf("searching package: %w", err)
	}

	fmt.Printf("Package: %s\nVersion: %s\nDescription: %s\n",
		pkg.Name, pkg.Version, pkg.Description)

	if !helpers.AskConfirmation(fmt.Sprintf("Install %s? (y/n) ", pkg.Name)) {
		return nil
	}

	if err := clonePackage(pkg.Name); err != nil {
		return fmt.Errorf("cloning package: %w", err)
	}

	if err := makePackage(pkg.Name); err != nil {
		return fmt.Errorf("building package: %w", err)
	}

	if err := helpers.AddInstalled(pkgName, pkg.Version); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not track installed package: %v\n", err)
	}

	fmt.Printf("Successfully installed %s\n", pkg.Name)
	return nil
}

func clonePackage(pkgName string) error {
	if err := helpers.EnsureDir(config.DataDir); err != nil {
		return err
	}

	pkgPath := filepath.Join(config.DataDir, pkgName)

	if _, err := os.Stat(pkgPath); err == nil {
		fmt.Println("Package already cloned at:", pkgPath)
		return nil
	}

	url := config.AURUrl + pkgName + ".git"
	fmt.Println("Cloning", url)

	return helpers.CloneRepo(url, pkgPath)
}

func makePackage(pkgName string) error {
	dir := filepath.Join(config.DataDir, pkgName)

	fmt.Println("Installing package", pkgName)

	return helpers.BuildPackage(dir)
}

package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/helpers"
	"github.com/joaogiacometti/aurgo/internal/models"
)

type AURResponse struct {
	Results []models.AurPackage `json:"results"`
}

func InstallPackage(pkgName string) error {
	pkg, err := SearchPackage(pkgName)
	if err != nil {
		return fmt.Errorf("searching package: %w", err)
	}

	fmt.Printf("Package: %s\nVersion: %s\nDescription: %s\n",
		pkg.Name, pkg.Version, pkg.Description)

	if !helpers.AskConfirmation(fmt.Sprintf("Install %s? (y/n) ", pkg.Name)) {
		return nil
	}

	if err := ClonePackage(pkg.Name); err != nil {
		return fmt.Errorf("cloning package: %w", err)
	}

	if err := MakePackage(pkg.Name); err != nil {
		return fmt.Errorf("building package: %w", err)
	}

	fmt.Printf("Successfully installed %s\n", pkg.Name)
	return nil
}

func SearchPackage(pkgName string) (*models.AurPackage, error) {
	url := config.SearchURL + pkgName

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query AUR: %w", err)
	}
	defer resp.Body.Close()

	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(aurResp.Results) == 0 {
		return nil, fmt.Errorf("package '%s' not found", pkgName)
	}

	return &aurResp.Results[0], nil
}

func ClonePackage(pkgName string) error {
	if err := os.MkdirAll(config.CacheDir, 0755); err != nil {
		return err
	}

	pkgPath := filepath.Join(config.CacheDir, pkgName)

	if _, err := os.Stat(pkgPath); err == nil {
		fmt.Println("Package already cloned at:", pkgPath)
		return nil
	}

	url := config.AurUrl + pkgName + ".git"
	fmt.Println("Cloning", url)

	cmd := exec.Command("git", "clone", "--depth=1", url, pkgPath)

	return cmd.Run()
}

func MakePackage(pkgName string) error {
	dir := filepath.Join(config.CacheDir, pkgName)

	cmd := exec.Command("makepkg", "-si", "--noconfirm")
	cmd.Dir = dir

	fmt.Println("Installing package", pkgName)
	return cmd.Run()
}

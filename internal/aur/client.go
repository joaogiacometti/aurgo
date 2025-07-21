package aur

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/types"
)

var httpClient = &http.Client{
	Timeout: time.Duration(config.HTTPTimeout) * time.Second,
}

func SearchPackages(query string) (types.AURResponse, error) {
	url := config.SearchURL + query

	resp, err := httpClient.Get(url)
	if err != nil {
		return types.AURResponse{}, fmt.Errorf("failed to query AUR: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.AURResponse{}, fmt.Errorf("AUR query failed with status: %d", resp.StatusCode)
	}

	var aurResponse types.AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResponse); err != nil {
		return types.AURResponse{}, fmt.Errorf("failed to decode AUR response: %w", err)
	}

	return aurResponse, nil
}

func GetPackageInfo(pkgName string) (*types.AURPackage, error) {
	url := config.InfoURL + pkgName

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query AUR: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AUR query failed with status: %d", resp.StatusCode)
	}

	var aurResp types.AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(aurResp.Results) == 0 {
		return nil, fmt.Errorf("package '%s' not found", pkgName)
	}

	return &aurResp.Results[0], nil
}

func FindPackage(pkgName string) (*types.AURPackage, error) {
	pkg, err := GetPackageInfo(pkgName)
	if err == nil {
		return pkg, nil
	}

	packages, err := SearchPackages(pkgName)
	if err != nil {
		return nil, err
	}

	if len(packages.Results) == 0 {
		return nil, fmt.Errorf("package '%s' not found", pkgName)
	}

	for _, pkg := range packages.Results {
		if pkg.Name == pkgName {
			return &pkg, nil
		}
	}

	return &packages.Results[0], nil
}

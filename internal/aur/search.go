package aur

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaogiacometti/aurgo/internal/config"
	"github.com/joaogiacometti/aurgo/internal/types"
)

func SearchPackage(query string) error {
	url := config.SearchURL + query

	resp, err := getClient().Get(url)
	if err != nil {
		return fmt.Errorf("failed to query AUR: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("AUR query failed with status: %d", resp.StatusCode)
	}

	var aurResponse types.AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResponse); err != nil {
		return fmt.Errorf("failed to decode AUR response: %w", err)
	}

	if len(aurResponse.Results) == 0 {
		fmt.Println("No packages found")
		return nil
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

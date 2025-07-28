package helpers

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/joaogiacometti/aurgo/config"
)

var mu sync.Mutex

type VersionsMap map[string]string

func ReadVersions() (VersionsMap, error) {
	mu.Lock()
	defer mu.Unlock()

	if err := EnsureDir(config.DataDir); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(config.VersionsFile)
	if os.IsNotExist(err) {
		return VersionsMap{}, nil
	}
	if err != nil {
		return nil, err
	}

	var installed VersionsMap
	err = json.Unmarshal(data, &installed)
	if err != nil {
		return nil, err
	}

	return installed, nil
}

func WriteVersions(versions VersionsMap) error {
	mu.Lock()
	defer mu.Unlock()

	if err := EnsureDir(config.DataDir); err != nil {
		return err
	}

	data, err := json.MarshalIndent(versions, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(config.VersionsFile, data, 0644)
}

func AddVersion(pkg, version string) error {
	versions, err := ReadVersions()
	if err != nil {
		return err
	}
	versions[pkg] = version
	return WriteVersions(versions)
}

func RemoveVersion(pkg string) error {
	versions, err := ReadVersions()
	if err != nil {
		return err
	}
	delete(versions, pkg)
	return WriteVersions(versions)
}

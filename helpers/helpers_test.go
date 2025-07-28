package helpers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joaogiacometti/aurgo/config"
)

func TestAddVersion(t *testing.T) {
	cleanup := setupTempConfig(t)
	defer cleanup()

	err := AddVersion("testpkg", "1.0.0")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	versions, err := ReadVersions()
	if err != nil {
		t.Fatalf("Expected no error reading versions, got %v", err)
	}

	if versions["testpkg"] != "1.0.0" {
		t.Errorf("Expected version 1.0.0 for testpkg, got %s", versions["testpkg"])
	}
}

func TestRemoveVersion(t *testing.T) {
	cleanup := setupTempConfig(t)
	defer cleanup()

	err := AddVersion("testpkg", "1.0.0")
	if err != nil {
		t.Fatalf("Expected no error adding version, got %v", err)
	}

	err = RemoveVersion("testpkg")
	if err != nil {
		t.Fatalf("Expected no error removing version, got %v", err)
	}

	versions, err := ReadVersions()
	if err != nil {
		t.Fatalf("Expected no error reading versions, got %v", err)
	}

	if _, exists := versions["testpkg"]; exists {
		t.Error("Expected testpkg to be removed, but it still exists")
	}
}

func TestWriteAndReadVersions(t *testing.T) {
	cleanup := setupTempConfig(t)
	defer cleanup()

	versions := VersionsMap{"testpkg": "1.0.0"}
	err := WriteVersions(versions)
	if err != nil {
		t.Fatalf("Expected no error writing versions, got %v", err)
	}

	readVersions, err := ReadVersions()
	if err != nil {
		t.Fatalf("Expected no error reading versions, got %v", err)
	}

	if readVersions["testpkg"] != "1.0.0" {
		t.Errorf("Expected version 1.0.0 for testpkg, got %s", readVersions["testpkg"])
	}
}

func setupTempConfig(t *testing.T) func() {
	t.Helper()

	tempDir := t.TempDir()
	config.DataDir = tempDir
	config.VersionsFile = filepath.Join(tempDir, "versions.json")

	return func() {
		os.RemoveAll(tempDir)
	}
}

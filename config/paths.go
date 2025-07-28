package config

import (
	"os"
	"path/filepath"
)

var (
	HomeDir, _   = os.UserHomeDir()
	DataDir      = filepath.Join(HomeDir, ".local", "share", "aurgo")
	VersionsFile = filepath.Join(DataDir, "versions.json")
	InfoURL      = "https://aur.archlinux.org/rpc/?v=5&type=info&arg="
	SearchURL    = "https://aur.archlinux.org/rpc/?v=5&type=search&arg="
	AURUrl       = "https://aur.archlinux.org/"
)

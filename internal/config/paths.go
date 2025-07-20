package config

import (
	"os"
	"path/filepath"
)

var (
	CacheDir     = filepath.Join(os.Getenv("HOME"), ".cache", "aurgo")
	DataDir      = filepath.Join(os.Getenv("HOME"), ".local", "share", "aurgo")
	InstalledTxt = filepath.Join(DataDir, "installed.txt")
	SearchURL    = "https://aur.archlinux.org/rpc/?v=5&type=info&arg="
	AurUrl       = "https://aur.archlinux.org/"
)

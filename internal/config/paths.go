package config

import (
	"os"
	"path/filepath"
)

var (
	CacheDir  = filepath.Join(os.Getenv("HOME"), ".cache", "aurgo")
	SearchURL = "https://aur.archlinux.org/rpc/?v=5&type=info&arg="
	AurUrl    = "https://aur.archlinux.org/"
)

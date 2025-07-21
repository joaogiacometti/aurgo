package config

import (
	"os"
	"path/filepath"
)

var (
	CacheDir  = filepath.Join(os.Getenv("HOME"), ".cache", "aurgo")
	InfoURL   = "https://aur.archlinux.org/rpc/?v=5&type=info&arg="
	SearchURL = "https://aur.archlinux.org/rpc/?v=5&type=search&arg="
	AURUrl    = "https://aur.archlinux.org/"
)

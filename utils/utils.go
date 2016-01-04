package utils

import (
	"path/filepath"
)

const (
	baseDefaultDir = "C:\\meld"
	steamCmdURL    = "https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip"
	oxideURL       = "https://github.com/OxideMod/Snapshots/raw/master/Oxide-Rust.zip"
)

var (
	SteamDefaultDir = filepath.Join(baseDefaultDir, "steam")
	RustDefaultDir  = filepath.Join(baseDefaultDir, "rust")
)

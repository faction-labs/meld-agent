package utils

import (
	"path/filepath"
	"runtime"
)

const (
	baseDefaultDirWin   = "C:\\meld"
	baseDefaultDirLinux = "/opt/meld"
	steamCmdURL         = "https://steamcdn-a.akamaihd.net/client/installer/steamcmd.zip"
	oxideURL            = "https://github.com/OxideMod/Snapshots/raw/master/Oxide-Rust.zip"
)

func GetSteamDir() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(baseDefaultDirWin, "steam")
	case "linux":
		return filepath.Join(baseDefaultDirLinux, "steam")
	default:
		return ""
	}
}

func GetRustDir() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(baseDefaultDirWin, "rust")
	case "linux":
		return filepath.Join(baseDefaultDirLinux, "rust")
	default:
		return ""
	}
}

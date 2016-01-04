package utils

import (
	"os"
	"path/filepath"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

const (
	steamCmdName = "steamcmd.exe"
)

func getSteamCmdName() string {
	switch runtime.GOOS {
	case "windows":
		return "steamcmd.exe"
	case "linux":
		return "steamcmd"
	default:
		return ""
	}
}

func GetSteamCmdPath(destPath string) (string, error) {
	if destPath == "" {
		destPath = GetSteamDir()
	}

	p := filepath.Join(destPath, getSteamCmdName())

	log.Debugf("checking for steam cmd: path=%s", p)

	if _, err := os.Stat(p); err != nil {
		return "", err
	}

	return p, nil
}

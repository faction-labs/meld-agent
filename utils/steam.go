package utils

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
)

const (
	steamCmdName = "steamcmd.exe"
	rustCmdName  = "RustDedicated.exe"
)

func GetSteamCmdPath(destPath string) (string, error) {
	if destPath == "" {
		destPath = SteamDefaultDir
	}

	p := filepath.Join(destPath, steamCmdName)

	log.Debugf("checking for steam cmd: path=%s", p)

	if _, err := os.Stat(p); err != nil {
		return "", err
	}

	return p, nil
}

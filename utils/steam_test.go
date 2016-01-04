package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInstallSteamCmd(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "meld-steam")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	if err := InstallSteamCmd(tempDir); err != nil {
		t.Error(err)
		return
	}

	destPath := filepath.Join(tempDir, steamCmdName)
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Errorf("unable to find downloaded binary at %s", destPath)
		return
	} else if err != nil {
		t.Error(err)
		return
	}
}

func TestInstallRust(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "meld-steam")
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	if err := InstallSteamCmd(tempDir); err != nil {
		t.Error(err)
		return
	}

	destPath := filepath.Join(tempDir, steamCmdName)
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Errorf("unable to find downloaded binary at %s", destPath)
		return
	} else if err != nil {
		t.Error(err)
		return
	}

	// install rust
	if err := InstallRust(destPath, tempDir, false); err != nil {
		t.Error(err)
		return
	}
}

package utils

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

// unzip file into the dest
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, f.Mode())
		} else {
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				log.Fatal(err)
				return err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// InstallSteamCmd installs the Steam CMD utility at the destination path
func InstallSteamCmd(destPath string) error {
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return err
	}

	log.Infof("downloading steamcmd: url=%s", steamCmdURL)

	destFile := filepath.Join(destPath, "steamcmd.zip")
	d, err := os.Create(destFile)
	if err != nil {
		return err
	}

	resp, err := http.Get(steamCmdURL)
	if err != nil {
		return err
	}

	n, err := io.Copy(d, resp.Body)
	if err != nil {
		return err
	}

	d.Close()

	log.Debugf("downloaded steamcmd: bytes=%d", n)

	defer os.Remove(destFile)

	// extract
	log.Debugf("extracting steamcmd: path=%s", destFile)
	if err := unzip(destFile, destPath); err != nil {
		return err
	}

	log.Infof("installed steam cmd: path=%s", destPath)

	return nil
}

func InstallRust(steamCmdPath, destPath string, update bool) error {
	log.Infof("installing/updating rust: path=%s", destPath)

	rustInstallArgs := []string{
		"+login",
		"anonymous",
		"+force_install_dir",
		destPath,
		"+app_update",
		"258550",
		"-beta",
		"experimental",
		"validate",
		"+quit",
	}

	cmd := exec.Command(steamCmdPath, rustInstallArgs...)
	if err := cmd.Start(); err != nil {
		return err
	}

	log.Info("waiting for install to finish")
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			// HACK: for some reason the steam installer exits
			// with a status code of 7.  le sigh.
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Printf("Exit Status: %d", status.ExitStatus())
				if status.ExitStatus() != 7 {
					return err
				}
			}
		} else {
			return err
		}
	}

	log.Info("install/update complete")

	return nil
}

func InstallOxideMod(rustPath string) error {
	destFile := filepath.Join(rustPath, "oxide.zip")
	d, err := os.Create(destFile)
	if err != nil {
		return err
	}

	log.Infof("downloading latest oxide snapshot: url=%s", oxideURL)

	resp, err := http.Get(oxideURL)
	if err != nil {
		return err
	}

	n, err := io.Copy(d, resp.Body)
	if err != nil {
		return err
	}

	d.Close()

	log.Debugf("downloaded oxide snapshot: bytes=%d", n)

	defer os.Remove(destFile)

	// extract
	log.Debugf("extracting oxide: path=%s", destFile)
	if err := unzip(destFile, rustPath); err != nil {
		return err
	}

	log.Infof("installed oxide mod: path=%s", rustPath)

	return nil
}

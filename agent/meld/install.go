package meld

import (
	"fmt"
	"os"

	"github.com/factionlabs/meld-agent/utils"
)

type InstallArgs struct {
	SteamPath string
	RustPath  string
}

func (m *Meld) InstallSteam(args *InstallArgs, success *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if args.SteamPath == "" {
		args.SteamPath = utils.GetSteamDir()
	}

	// check for existing
	_, err := utils.GetSteamCmdPath(args.SteamPath)
	if err != nil {
		if os.IsNotExist(err) {
			// install
			if err := utils.InstallSteamCmd(args.SteamPath); err != nil {
				*success = false
				return err
			}

			if _, err := utils.GetSteamCmdPath(args.SteamPath); err != nil {
				*success = false
				return err
			}
		} else {
			*success = false
			return err
		}
	}

	*success = true
	return nil
}

func (m *Meld) InstallRust(args *InstallArgs, success *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if args.SteamPath == "" {
		args.SteamPath = utils.GetSteamDir()
	}

	if args.RustPath == "" {
		args.RustPath = utils.GetRustDir()
	}

	// check for existing
	steamCmdPath, err := utils.GetSteamCmdPath(args.SteamPath)
	if err != nil {
		*success = false
		return fmt.Errorf("unable to find steam; perhaps try installing; err=%s", err)
	}

	if err := utils.InstallRust(steamCmdPath, args.RustPath, true); err != nil {
		*success = false
		return err
	}

	return nil
}

func (m *Meld) InstallOxide(args *InstallArgs, success *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if args.RustPath == "" {
		args.RustPath = utils.GetRustDir()
	}

	if err := utils.InstallOxideMod(args.RustPath); err != nil {
		*success = false
		return fmt.Errorf("error installing oxide: %s", err)
	}

	*success = true
	return nil
}

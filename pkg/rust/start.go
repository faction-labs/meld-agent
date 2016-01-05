package rust

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld-agent/utils"
)

func getWindowsRunScript(args *StartArgs) (string, error) {
	startTemplate := `
@echo off
@title = Meld Rust Server
cls
:start

{{.BinPath}} -batchmode -nographics +server.ip {{.ServerIP}} +rcon.ip {{.RconIP}} +server.port {{.ServerPort}} +rcon.port {{.RconPort}} +rcon.password "{{.RconPassword}}" +server.maxplayers {{.MaxPlayers}} +server.hostname "{{.Hostname}}" +server.identity "{{.Identity}}" +server.level "{{.Level}}" +server.seed {{.Seed}} +server.worldsize {{.WorldSize}} +server.saveinterval {{.SaveInterval}} +server.globalchat true +server.description "{{.Description}}" +server.url "{{.URL}}"

@exit
`
	t := template.Must(template.New("startTemplateWin").Parse(startTemplate))
	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, args); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func getLinuxRunScript(args *StartArgs) (string, error) {
	// TODO: finish
	return "echo not implemented", nil
}

func getRunScript(args *StartArgs) (string, error) {
	switch runtime.GOOS {
	case "windows":
		return getWindowsRunScript(args)
	case "linux":
		return getLinuxRunScript(args)
	default:
		return "", fmt.Errorf("unknown platform")
	}
}
func getRunScriptPath() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(utils.GetRustDir(), "run-meld.bat"), nil
	case "linux":
		return filepath.Join(utils.GetRustDir(), "run-meld.sh"), nil
	default:
		return "", fmt.Errorf("unknown platform")
	}
}

// Start runs a new Rust Server
func (r *RustServer) Start(args *StartArgs) (int, error) {
	log.Debugf("starting rust server: args=%v", args)

	binPath := RustDedicatedPath()
	scriptPath, err := getRunScriptPath()
	if err != nil {
		return -1, err
	}

	args.BinPath = binPath

	log.Debugf("creating run script: path=%s", scriptPath)
	runScript, err := getRunScript(args)
	if err != nil {
		return -1, err
	}

	// remove existing
	if s, _ := os.Stat(scriptPath); s != nil {
		os.Remove(scriptPath)
	}

	if s, _ := os.Stat(rustServerPidPath); s != nil {
		os.Remove(rustServerPidPath)
	}

	// create
	f, err := os.Create(scriptPath)
	if err != nil {
		return -1, err
	}

	f.Write([]byte(runScript))
	f.Close()

	if err := os.Chmod(scriptPath, 0755); err != nil {
		return -1, err
	}

	sh := ""
	cArgs := []string{}

	switch runtime.GOOS {
	case "windows":
		sh = "cmd.exe"
		cArgs = []string{
			"/C",
			"start",
			scriptPath,
			">meld-server.log",
			"2>&1",
		}
	case "linux":
		sh = "bash"
		cArgs = []string{
			"-s",
			scriptPath,
			">meld-server.log",
			"2>&1",
		}
	}

	c := exec.Command(sh, cArgs...)
	c.Dir = utils.GetRustDir()

	log.Debugf("starting rust server: cmd=%s args=%v", c.Path, c.Args)

	if err := c.Start(); err != nil {
		return -1, err
	}

	// wait slightly for process to start
	time.Sleep(time.Millisecond * 500)

	pid := c.Process.Pid

	// write pid file
	pf, err := os.Create(rustServerPidPath)
	if err != nil {
		return -1, err
	}

	pf.Write([]byte(fmt.Sprintf("%d", pid)))
	pf.Close()

	log.Debugf("rust server started: pid=%d", pid)

	return pid, nil
}

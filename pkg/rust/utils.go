package rust

import (
	"path/filepath"
	"runtime"

	"github.com/factionlabs/meld-agent/utils"
)

func RustDedicatedPath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(utils.GetRustDir(), rustDedicatedWin)
	case "linux":
		return filepath.Join(utils.GetRustDir(), rustDedicatedLinux)
	default:
		return ""
	}

}

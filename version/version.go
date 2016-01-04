package version

import (
	"fmt"
)

var (
	Name = "meld-agent"

	Version = "0.1.0"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

func FullName() string {
	return Name
}

func FullVersion() string {
	return fmt.Sprintf("%s (sha: %s)", Version, GitCommit)
}

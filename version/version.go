package version

import (
	"bytes"
	"fmt"
)

// GitCommit that was compiled. This will be filled in by the compiler.
var GitCommit string

// GitDescribe that was compiled. This will be filled in by the compiler.
var GitDescribe string

// Version is the main version number that is being run at the moment.
const Version = "0.1.0"

// VersionPrerelease is a pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
const VersionPrerelease = "dev"

// Info contains all version info.
type Info struct {
	Revision          string
	Version           string
	VersionPrerelease string
}

// GetVersion gets the version.
func GetVersion() *Info {
	ver := Version
	rel := VersionPrerelease
	if GitDescribe != "" {
		ver = GitDescribe
	}
	if GitDescribe == "" && rel == "" && VersionPrerelease != "" {
		rel = "dev"
	}

	return &Info{
		Revision:          GitCommit,
		Version:           ver,
		VersionPrerelease: rel,
	}
}

func (c *Info) String() string {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "Stackctl v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	return versionString.String()
}

package build

import (
	"fmt"
	"runtime"
)

var (
	Version = "dev"
	Build   = "dev"
	Date    = ""
)

func LogVersionInfo() {
	fmt.Printf(`
************************************************
Version: %v
Build: %v
Date: %v
OS: %s
Arch: %s
Go: %s
************************************************
`,
		Version, Build, Date, runtime.GOOS, runtime.GOARCH, runtime.Version())
}

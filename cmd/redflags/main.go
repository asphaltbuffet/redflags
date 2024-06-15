package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/asphaltbuffet/redflags"
)

var (
	version  = "dev" // set by build script
	revision = "n/a" // set by build script
)

func main() {
	// override the builtin -V flag
	flag.Var(versionFlag{}, "V", "print version and exit")
	singlechecker.Main(redflags.New(nil))
}

type versionFlag struct{}

func (versionFlag) String() string   { return "" }
func (versionFlag) IsBoolFlag() bool { return true }
func (versionFlag) Set(string) error {
	fmt.Printf("%s (%s)\n", version, revision)
	os.Exit(0)
	return nil
}

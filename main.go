package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin"
)

var (
	command string

	app     = kingpin.New(Name, Description)
	verbose = app.Flag("verbose", "Verbose mode.").Bool()
	version = app.Flag("version", "Show version and terminate").Short('v').Bool()
)

func init() {

	app.HelpFlag.Short('h')
	command = kingpin.MustParse(app.Parse(os.Args[1:]))

	if *version {
		fmt.Printf("%s version %s build %s (%s), build on %s\n", Name, BuildVersion, BuildHash, runtime.GOARCH, BuildDate)
		os.Exit(0)
	}

}
func main() {}

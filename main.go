package main

import (
	"fmt"
	"os"
	"runtime"

	"git.matoski.com/ilijamt/proxy-checker/job"

	"github.com/alecthomas/kingpin"
)

var (
	command string

	app       = kingpin.New(Name, Description)
	version   = app.Flag("version", "Show version and terminate").Short('v').Bool()
	queueSize = app.Flag("queue", "How many request to process at one time").Default("25").Int()

	check         = app.Command("check", "Check the single proxy")
	checkHostPort = check.Arg("host-port", "The hostname of the proxy with the proxy, add the schema to the URL like https://proxy.com:9000").Required().String()
	checkUsername = check.Arg("username", "The username to use for the proxy, can be empty if no authentication is required").String()
	checkPassword = check.Arg("password", "The password to use for the proxy, can be empty if no authentication is required").String()

	queue *job.Queue
)

func init() {

	app.HelpFlag.Short('h')
	command = kingpin.MustParse(app.Parse(os.Args[1:]))

	if *version {
		fmt.Printf("%s version %s build %s (%s), build on %s\n", Name, BuildVersion, BuildHash, runtime.GOARCH, BuildDate)
		os.Exit(0)
	}

	queue = job.NewQueue(*queueSize)

}

func main() {

	go queue.Run()

	switch command {
	case "check":

		queue.AddJob(job.Detail{
			Host:     *checkHostPort,
			Username: *checkUsername,
			Password: *checkPassword,
		})

	}

	queue.Wait()

}

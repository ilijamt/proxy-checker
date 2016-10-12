package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin"

	"git.matoski.com/ilijamt/proxy-checker/job"
)

var (
	command string

	app        = kingpin.New(Name, Description)
	version    = app.Flag("version", "Show version and terminate").Short('v').Bool()
	queueSize  = app.Flag("queue", "How many request to process at one time").Default("25").Int()
	failedOnly = app.Flag("failed-only", "Show only failed proxies").Bool()

	check         = app.Command("check", "Check the single proxy")
	checkHostPort = check.Arg("host-port", "The hostname of the proxy with the proxy, add the schema to the URL like https://proxy.com:9000").Required().String()
	checkUsername = check.Arg("username", "The username to use for the proxy, can be empty if no authentication is required").String()
	checkPassword = check.Arg("password", "The password to use for the proxy, can be empty if no authentication is required").String()

	file     = app.Command("csv-file", "Check all the proxies in the file specified")
	fileName = file.Arg("name", "The file name to load").Required().File()

	queue *job.Queue
)

func init() {

	app.HelpFlag.Short('h')
	command = kingpin.MustParse(app.Parse(os.Args[1:]))

	if *version {
		fmt.Printf("%s version %s build %s (%s), build on %s\n", Name, BuildVersion, BuildHash, runtime.GOARCH, BuildDate)
		os.Exit(0)
	}

	queue = job.NewQueue(*queueSize, *failedOnly)

}

func main() {

	defer queue.Wait()
	go queue.Run()

	switch command {
	case "check":

		queue.WgIncr()
		go queue.AddJob(job.Detail{
			Host:     *checkHostPort,
			Username: *checkUsername,
			Password: *checkPassword,
		})

	case "csv-file":
		f, e := os.Open((*fileName).Name())

		if e != nil {
			fmt.Errorf("%v", e)
			return
		}

		r := csv.NewReader(bufio.NewReader(f))
		result, _ := r.ReadAll()

		var host string
		var username string
		var password string

		for i := range result {

			host = result[i][0]
			username = result[i][1]
			password = result[i][2]

			queue.WgIncr()
			go queue.AddJob(job.Detail{
				Host:     host,
				Username: username,
				Password: password,
			})

		}

	}

}

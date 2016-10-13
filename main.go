package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"

	"github.com/alecthomas/kingpin"

	"github.com/ilijamt/proxy-checker/job"
)

var (
	command string

	app        = kingpin.New(Name, Description)
	version    = app.Command("version", "Show version and terminate").Action(ShowVersion)
	queueSize  = app.Flag("queue", "How many request to process at one time").Default("25").Int()
	failedOnly = app.Flag("failed-only", "Show only failed proxies").Bool()

	check         = app.Command("check", "Check the single proxy")
	checkHostPort = check.Arg("host-port", "The hostname of the proxy with the proxy, add the schema to the URL like https://proxy.com:9000").Required().String()
	checkUsername = check.Arg("username", "The username to use for the proxy, can be empty if no authentication is required").String()
	checkPassword = check.Arg("password", "The password to use for the proxy, can be empty if no authentication is required").String()

	file     = app.Command("csv-file", "Check all the proxies in the file specified")
	fileName = file.Arg("file", "The file name to load").Required().File()

	queue *job.Queue
)

func ShowVersion(c *kingpin.ParseContext) error {
	fmt.Printf("%s version %s build %s (%s), built on %s, by %s\n", Name, BuildVersion, BuildHash, runtime.GOARCH, BuildDate, Maintainer)
	os.Exit(0)
	return nil
}

func init() {

	app.HelpFlag.Short('h')
	command = kingpin.MustParse(app.Parse(os.Args[1:]))

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
			_ = fmt.Errorf("%v", e)
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

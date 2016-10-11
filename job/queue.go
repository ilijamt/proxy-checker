package job

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Queue struct {
	jobs chan Detail
	quit chan bool

	queryURL *url.URL

	failedOnly bool

	wg sync.WaitGroup
}

func (q *Queue) Wait() {
	q.wg.Wait()
}

func NewQueue(queueSize int, failedOnly bool) *Queue {
	queue := &Queue{}
	queue.jobs = make(chan Detail, queueSize)
	queue.quit = make(chan bool, 1)
	queue.queryURL, _ = url.Parse("https://api.ipify.org")
	queue.failedOnly = failedOnly
	return queue
}

func (q *Queue) Run() {
	var quit bool = false

	for {
		select {

		// we got a new job to process
		case j := <-q.jobs:
			go q.IsValidProxy(j)

		// we got the signal to quit
		case <-q.quit:
			quit = true
		}

		// quit if we have finished
		if quit {
			break
		}
	}
}

func (q *Queue) Stop() {
	q.quit <- true
}

func (q *Queue) IsValidProxy(job Detail) bool {
	defer q.wg.Done()

	var err error
	var validStatus bool = false
	var validResponse = false

	var proxyURL, _ = job.ToURL()

	defer func() {
		var rsp = "invalid"
		if validStatus && validResponse {
			rsp = "valid"
		}
		if q.failedOnly {
			if !(validStatus && validResponse) {
				fmt.Printf("%s is %s\n", proxyURL.String(), rsp)
			}
		} else {
			fmt.Printf("%s is %s\n", proxyURL.String(), rsp)
		}

	}()

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, _ := http.NewRequest("GET", q.queryURL.String(), nil)
	response, err := client.Do(request)

	if err != nil {
		return validStatus && validResponse
	}

	body, _ := ioutil.ReadAll(response.Body)

	validStatus = response.Status == "200 OK"
	validResponse = strings.Contains(proxyURL.String(), string(body))

	return validStatus && validResponse
}

func (q *Queue) WgIncr() {
	q.wg.Add(1)
}

func (q *Queue) AddJob(job Detail) {
	q.jobs <- job
}

package job

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	"context"
	"io/ioutil"
)

type Queue struct {
	jobs chan Detail
	quit chan bool

	queryURL *url.URL

	failedOnly bool

	wg   sync.WaitGroup
}

func (q *Queue) Wait() {
	q.wg.Wait()
}

func NewQueue(queueSize int, host string, failedOnly bool) *Queue {
	queue := &Queue{}
	queue.jobs = make(chan Detail, queueSize)
	queue.quit = make(chan bool, 1)
	queue.queryURL, _ = url.Parse(host)
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
	var response *http.Response
	var proxyURL, _ = job.ToURL()

	defer func() {
		var rsp = "invalid"
		if validStatus {
			rsp = "valid"
		}

		var status string
		var contentLength int64

		if response == nil {
			status = "UNKNOWN"
			contentLength = 0
		} else {
			status = response.Status
			contentLength = response.ContentLength
			if contentLength <= 0 {
				body, _ := ioutil.ReadAll(response.Body)
				l := len(body)
				contentLength = int64(l)
			}
		}

		if q.failedOnly {
			if !validStatus {
				fmt.Printf("%s is %s (Status Code: %s, Content Length: %d)\n", proxyURL.String(), rsp, status, contentLength)
			}
		} else {
			fmt.Printf("%s is %s (Status Code: %s, Content Length: %d)\n", proxyURL.String(), rsp, status, contentLength)
			if validStatus {
				ioutil.ReadAll(response.Body)
			}
		}

	}()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	req, _ := http.NewRequest("GET", q.queryURL.String(), nil)
	req = req.WithContext(ctx)
	response, err = client.Do(req)

	validStatus = err == nil
	if response != nil {
		validStatus = err == nil && response.Status == "200 OK"
	}

	return validStatus
}

func (q *Queue) WgIncr() {
	q.wg.Add(1)
}

func (q *Queue) AddJob(job Detail) {
	q.jobs <- job
}

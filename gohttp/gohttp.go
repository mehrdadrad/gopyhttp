package gohttp

import (
	"net/http"
	"sync"
)

const MAXThread = 100

type task interface {
	exec()
	output()
}

type taskManager interface {
	get() chan task
}

func run(tm taskManager, out chan task) {
	tasks := make(chan task)
	var wg sync.WaitGroup

	go func() {
		for t := range tm.get() {
			tasks <- t
		}

		close(tasks)
	}()

	for i := 0; i < MAXThread; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				t.exec()
				out <- t
			}
		}()
	}

	wg.Wait()
	close(out)
}

func Run(requests []Request) {
	out := make(chan task, 1000)
	f := new(taskFactory)
	f.requests = requests
	run(f, out)

	for task := range out {
		task.output()
	}
}

func TestMe() {
	println("hello")
}

type Request struct {
	URL    string
	method string
}

type taskFactory struct {
	requests []Request
}

func (h Request) exec() {
	req(h.URL)
}

func (h Request) output() {
}

func (t *taskFactory) get() chan task {
	tCh := make(chan task)

	go func() {
		for _, r := range t.requests {
			tCh <- r
		}

		close(tCh)
	}()

	return tCh
}

func req(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	println(resp.StatusCode, url)
}

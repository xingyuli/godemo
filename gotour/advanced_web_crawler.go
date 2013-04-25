package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var fmt = log.New(os.Stdout, "", 0)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	m := map[string]bool{url: true}
	var mx sync.Mutex
	var wg sync.WaitGroup
	var _crawl_ func(string, int)
	_crawl_ = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("found: %s %q\n", url, body)
		mx.Lock()
		for _, u := range urls {
			if !m[u] {
				m[u] = true
				wg.Add(1)
				go _crawl_(u, depth-1)
			}
		}
		mx.Unlock()
	}

	wg.Add(1)
	_crawl_(url, depth)
	wg.Wait()
}

func main() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		Crawl("http://golang.org/", 4, fetcher)
	}
	fmt.Println("ellapse:", time.Now().Sub(start))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type UrlNotFoundError struct {
	url string
}

func (err UrlNotFoundError) Error() string {
	return "not found: " + err.url
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, UrlNotFoundError{url}
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

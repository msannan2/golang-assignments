package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	v    map[string]bool
	mt   sync.Mutex
	wait sync.WaitGroup
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func (m *SafeMap) checker(u string) bool {
	m.mt.Lock()
	defer m.mt.Unlock()
	_, ok := m.v[u]
	if !ok {
		m.v[u] = true
		return false
	}
	return true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (m *SafeMap) Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 || m.checker(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		m.wait.Add(1)
		go m.Crawl(u, depth-1, fetcher)
		m.wait.Done()
	}

	return
}

func main() {
	temp := SafeMap{v: make(map[string]bool)}
	temp.Crawl("http://golang.org/", 4, fetcher)
	fmt.Println(temp.v)
	temp.wait.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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

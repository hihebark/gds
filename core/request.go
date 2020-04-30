package core

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

// Request struct
type Request struct {
	wg        sync.WaitGroup
	client    *http.Client
	finish    chan bool
	options   Options
	req       *http.Request
	responses []Response
	sync.RWMutex
}

// NewRequest function
func NewRequest(options Options) *Request {
	return &Request{
		wg:        sync.WaitGroup{},
		client:    &http.Client{},
		finish:    make(chan bool, 2),
		options:   options,
		responses: []Response{},
	}
}

// Run run the brutforce
func (r *Request) Run() {
	transport := &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	if r.options.Proxy != "" {
		urlProxy, err := url.Parse(r.options.Proxy)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		transport.Proxy = http.ProxyURL(urlProxy)
	}
	if r.options.Tor {
		torurl, err := url.Parse("socks5://127.0.0.1:9050")
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		dialer, err := proxy.FromURL(torurl, proxy.Direct)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		transport.Dial = dialer.Dial
	}
	r.client = &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}
	url, err := url.Parse(r.options.URL)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	r.req = &http.Request{
		URL:    url,
		Method: "GET",
		Header: map[string][]string{
			"Cookie":          {r.options.Cookie},
			"User-Agent":      {r.options.UserAgent},
			"Accept-Encoding": {"identity", ""},
			"Connection":      {"close"},
		},
		Close: true,
	}
	fmt.Printf("Starting....\n")
	r.start()
}

func (r *Request) start() {
	r.wg.Add(1)
	lines := make(chan string)
	go readFile(r.options.Wordlist, lines)
	go func() {
		for {
			line, ok := <-lines
			if !ok {
				break
			}
			if line != "" {
				go r.dial(line)
			}
		}
		r.wg.Done()
	}()
	r.wg.Wait()
}

func (r *Request) dial(path string) {
	r.wg.Add(1)
	r.Lock()
	r.req.URL.Path = path
	res, err := r.client.Do(r.req)
	if err != nil {
		fmt.Printf("Request:dial%v\n", err)
	}
	response := Response{
		Timestamp: time.Now().Unix(),
		Link:      r.req.URL.String(),
		Status:    res.StatusCode,
		Length:    byteConverter(res.ContentLength),
	}
	res.Body.Close()
	if r.options.Output != "" {
		r.responses = append(r.responses, response)
	}
	if response.Status >= 200 && response.Status <= 500 {
		fmt.Printf("[%d] %d - %10s - %s\n", response.Timestamp, response.Status, response.Length, response.Link)
	}
	r.Unlock()
	r.wg.Done()
}

// Healthz function
func Healthz(host string) (int, error) {
	resp, err := http.Get(host)
	return resp.StatusCode, err
}

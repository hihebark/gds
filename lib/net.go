package lib

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

type work struct {
	wg sync.WaitGroup
	sync.RWMutex
	threads int
	client  http.Client
	datas   DataSlice
	path    chan string
	done    chan bool
}

func newWork(thread int, c http.Client) *work {
	if thread <= 0 {
		thread = 2 * runtime.NumCPU()
	}
	return &work{
		wg:      sync.WaitGroup{},
		threads: thread,
		datas:   DataSlice{},
		client:  c,
		path:    make(chan string),
		done:    make(chan bool, 2),
	}
}

//Data json format
type Data struct {
	ID         int    `json:"id"`
	URL        string `json:"url"`
	Status     int    `json:"status"`
	Length     string `json:"length"`
	Screenshot string `json:"screenshot"`
}

//DataSlice json format
type DataSlice struct {
	Data []Data `json:"host"`
}

//Options for Request data.
type Options struct {
	Host       string
	Proxyfile  string
	Wordlist   string
	UserAgent  string
	Cookie     string
	Ex         []string
	Proxy      string
	Tor        bool
	ResultFile string
	IsUp       bool
	Thread     int
}

//StartWork start brutforcing...
func StartWork(o Options) {

	transport := &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}
	if o.Proxy != "" {
		urlProxy, err := url.Parse(o.Proxy)
		Printerr(err, "StartWork:url.Parse")
		transport.Proxy = http.ProxyURL(urlProxy)
	}
	if o.Tor {
		transport.Dial = throwTor().Dial
	}
	wordlist := readFromFile(o.Wordlist)
	if len(wordlist) == 0 {
		Bad("The file is empty!")
		os.Exit(1)
	}
	tnow := time.Now().Format("15:04:05")
	Info(fmt.Sprintf("Wordlist size: %d / Extensions:%s / Starting time: %s\n",
					len(wordlist), o.Ex, tnow))
	client := &http.Client{
						Transport:     transport,
						Timeout:       10 * time.Second,
					}
	work := newWork(o.Thread, *client)
	u, _ := url.Parse(o.Host)
	req := &http.Request{
		URL:    u,
		Method: "GET",
		Header: map[string][]string{
			"Cookie":          {o.Cookie},
			"User-Agent":      {o.UserAgent},
			"Accept-Encoding": {"identity", ""},
		},
		Close: true,
	}
	go work.producer(wordlist, o.Ex)
	for i := 0; i <= work.threads; i++ {
		go work.consumer(req)
	}
	work.wg.Wait()
	<- work.done
}

func (w *work) producer(wl []string, ext []string) {
	for _, path := range wl {
		w.wg.Add(1)
		w.path <- path
		//fmt.Printf("%s\n", path)
		if string(path[len(path)-1]) != "/" && len(ext) >= 1 && ext[0] != "" {
			w.wg.Add(1)
			for _, e := range ext {
				w.path <- path + "." + e
			}
		}
	}
	w.done <- true
}

func (w *work) consumer(r *http.Request) {
	for p := range w.path {
		w.Lock()
		r.URL.Path = p
		resp, err := w.client.Do(r)
		if err != nil {
			fmt.Printf("net:consumer-%d: path: %s error: %v\n", p, err)
			//continue
		}
		//showOutput(resp.StatusCode, byteConverter(resp.ContentLength), r.URL.String())
		fmt.Printf("%d - %10s -\t%s\n",
				resp.StatusCode, byteConverter(resp.ContentLength), resp.Request.URL.String())
		w.Unlock()
		w.wg.Done()
	}
	w.done <- true
	//return
}

//CheckConnectivity check if the provided host is up or not.
func CheckConnectivity(host string) (int, error) {
	resp, err := http.Get(host)
	return resp.StatusCode, err
}

func byteConverter(length int64) string {
	mbyte := []string{"bytes", "KB", "MB", "GB", "TB"}
	if length == -1 {
		return "0 byte"
	}
	for _, x := range mbyte {
		if length < 1024.0 {
			return fmt.Sprintf("%3.1d %s", length, x)
		}
		length = length / 1024.0
	}
	return "Error"
}

func throwTor() proxy.Dialer {
	torurl, err := url.Parse("socks5://127.0.0.1:9050")
	Printerr(err, "ThrowTor:url.Parse")
	dialer, err := proxy.FromURL(torurl, proxy.Direct)
	Printerr(err, "ThrowTor:proxy.FromURL")
	return dialer
}

func showOutput(status int, length string, url string) {
	os.Stdout.Sync()
	switch {
	case status >= 100 && status <= 102:
		Say(LIGHTCYAN, fmt.Sprintf("%d - %10s - %s", status, length, url))
	case status >= 200 && status <= 226:
		Say(LIGHTGREEN, fmt.Sprintf("%d - %10s - %s", status, length, url))
	case status >= 300 && status <= 308:
		Say(LIGHTBLUE, fmt.Sprintf("%d - %10s - %s", status, length, url))
	case status >= 400 && status <= 451:
		//fmt.Printf("Testing: %s\r", url)
		Say(LIGHTRED, fmt.Sprintf("%d - %-10s - %s", status, length, url))
	case status >= 500 && status <= 512:
		Say(YELLOW, fmt.Sprintf("%d - %10s - %s", status, length, url))
	}
}

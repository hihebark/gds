package lib

import (
	_"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

type Work struct{
	wg sync.WaitGroup
	sync.RWMutex
	threads int
	client http.Client
	datas DataSlice
	path chan string
	done chan bool
}

func NewWork (thread int, c http.Client) *Work {
	if thread <= 0 {
		thread = 2*runtime.NumCPU()
	}
	return &Work{
		wg:       sync.WaitGroup{},
		threads:  thread,
		datas:    DataSlice{},
		client:   c,
		path:     make(chan string),
		done:     make(chan bool),
	}
}

//WebServer json format
type Data struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Status int    `json:"status"`
	Length string `json:"length"`
	Screenshot string `json:"screenshot"`
}

//WebServerslice json format
type DataSlice struct {
	Data []Data `json:"host"`
}

//NetRequest for Request data.
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
	IsUp     bool
}

func StartWork(o Options) {
	
	//path := make(chan string)
	//done := make(chan bool)
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	if o.Proxy != "" {
		urlProxy, err := url.Parse(o.Proxy)
		Printerr(err, "Fuxe:url.Parse")
		transport.Proxy = http.ProxyURL(urlProxy)
	}
	if o.Tor {
		transport.Dial = ThrowTor().Dial
	}
	wordlist := ReadFromFile(o.Wordlist)
	if len(wordlist) == 0 {
		Bad("The file is empty!")
		os.Exit(1)
	}
	Info(fmt.Sprintf("Wordlist size: %d / Extensions:%s\n", len(wordlist), o.Ex))
	client := &http.Client{Transport: transport}
	work := NewWork(0, *client)
	u, _ := url.Parse(o.Host)
	req := &http.Request{
		URL:    u,
		Method: "GET",
		Header: map[string][]string{
			"Cookie":          {o.Cookie},
			"User-Agent":      {o.UserAgent},
			"Accept-Encoding": {"identity", ""},
		},
	}
	go work.producer(wordlist, o.Ex)
	for i := 0; i <= work.threads; i++ {
		go work.consumer(req)
	}
	work.wg.Wait()
	<- work.done
}

func (w *Work)producer(wl []string, ext []string) {
	for _, path := range wl {
		w.path <- path
		if string(path[len(path)-1]) != "/" && len(ext) >= 1 && ext[0] != "" {
			for _, e := range ext {
				w.path <- path + "." +e
			}
		}
	}
	w.done <- true
}

func (w *Work)consumer(r *http.Request) {
	
	for p := range w.path {
		w.Lock()
		w.wg.Add(1)
		r.URL.Path = p
		resp, err := w.client.Do(r)
		if err != nil{
			fmt.Printf("error: %s - %v\n", p, err)
			continue
		}
		fmt.Printf("%d - %10s - \t%s\n", resp.StatusCode, ByteConverter(resp.ContentLength), r.URL.String())
		w.Unlock()
	}
	w.wg.Done()
	return
}

// StartSearch to brute force the sitweb
// param NetRequest (struct)
//func StartSearch(netreq NetRequest){

//	transport := &http.Transport{
//		MaxIdleConns:       10,
//		IdleConnTimeout:    30 * time.Second,
//		DisableCompression: true,
//	}
//	if netreq.Proxy != "" {
//		urlProxy, err := url.Parse(netreq.Proxy)
//		Printerr(err, "Fuxe:url.Parse")
//		transport.Proxy = http.ProxyURL(urlProxy)
//	}
//	if netreq.Tor {
//		transport.Dial = ThrowTor().Dial
//	}
//	allPath := ReadFromFile(netreq.Wordlist)
//	pathLength := len(allPath)
//	if pathLength == 0 {
//		Bad("the file is empty!")
//		os.Exit(1)
//	}
//	Info(fmt.Sprintf("Wordlist size: %d / Extensions:%s\n", pathLength, netreq.Ex))
//	
//	client := &http.Client{Transport: transport}
//	req := &http.Request{
//		Method: "GET",
//		Header: map[string][]string{
//			"Cookie":     {netreq.Cookie},
//			"User-Agent": {netreq.UserAgent},
//		},
//	}
//	for i := 0; i < pathLength; i++ {
//		t <- 0
//		wg.Add(1)
//		req.URL, _ = url.Parse(netreq.Host + allPath[i])
//		go doRequest(&wg, req, *client, i, netreq.Ex)
//	}
//	wg.Wait()
//	jsonF, _ := json.Marshal(webserver)
//	timenow := time.Now().Format("2006-01-02-15-04-05")
//	filePath := "data/results/" + netreq.ResultFile + strings.Split(netreq.ResultFile, "/")[0] + "-" + timenow + ".json"
//	WriteToFile(filePath, fmt.Sprintf("%+v\n", string(jsonF)))

//}

//// DoRequest to make request
//// param *http.Request, http.Client
//func doRequest(wg *sync.WaitGroup, req *http.Request, client http.Client, i int, ex []string) {
//	defer wg.Done()
//	response, err := client.Get(req.URL.String())
//	Printerr(err, fmt.Sprintf("DoRequest: %s", req.URL))
//	wb := WebServer{
//		ID:     i,
//		URL:    req.URL.String(),
//		Status: response.StatusCode,
//		Length: ByteConverter(response.ContentLength),
//	}
//	webserver.WebServers = append(webserver.WebServers, wb)
//	ShowOutput(response.StatusCode, ByteConverter(response.ContentLength), req.URL.String())	
//	<-t
//}

//CheckConnectivity check if the provided host is up or not.
func CheckConnectivity(host string) int {

	resp, err := http.Get(host)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
	return resp.StatusCode

}

// ByteConverter convert length to bytes, KB, MB, GB, TB.
// param  int64
// return string
func ByteConverter(length int64) string {
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
	return ""
}

//GetBody fetch the body
func GetBody(req *http.Request) (string, error) {

	client := &http.Client{}
	response, err := client.Get(req.URL.String())
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

//ThrowTor activate the the app to go throw Tor
func ThrowTor() proxy.Dialer {
	torurl, err := url.Parse("socks5://127.0.0.1:9050")
	Printerr(err, "ThrowTor:url.Parse")
	dialer, err := proxy.FromURL(torurl, proxy.Direct)
	Printerr(err, "ThrowTor:proxy.FromURL")
	return dialer
}

//ShowOutput print pretty output from a request
func ShowOutput(status int, length string, url string) {
	switch {
	case status >= 100 && status <= 102:
		Say(LIGHTCYAN, fmt.Sprintf("%d - %-10s - %s", status, length, url))
	case status >= 200 && status <= 226:
		Say(LIGHTGREEN, fmt.Sprintf("%d - %-10s - %s", status, length, url))
	case status >= 300 && status <= 308:
		Say(LIGHTBLUE, fmt.Sprintf("%d - %-10s - %s", status, length, url))
	case status >= 400 && status <= 451:
		//os.Stdout.Sync()
		//fmt.Printf("Testing: %s\r", url)
		//Say(LIGHTRED, fmt.Sprintf("%d - %-10s\t - %s", status, length, url))
	case status >= 500 && status <= 512:
		Say(YELLOW, fmt.Sprintf("%d - %-10s - %s", status, length, url))
	}
}

//ReturnURL to return HTTP.URL
func ReturnURL(host string) (*url.URL, error) {
	murl, err := url.ParseRequestURI(host)
	return murl, err
}

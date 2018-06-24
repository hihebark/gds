package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	_"runtime"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var (
	wg sync.WaitGroup
	mutex       = &sync.Mutex{}
	webserver   WebServerslice
	t 			= make(chan int, 8)
)

//NetRequest for Request data.
type NetRequest struct {
	Host       string
	Proxyfile  string
	Wordlist   string
	UserAgent  string
	Cookie     string
	Ex         []string
	Proxy      string
	Tor        bool
	ResultFile string
}

//WebServer json format
type WebServer struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	Status int    `json:"status"`
	Length string `json:"length"`
}

//WebServerslice json format
type WebServerslice struct {
	WebServers []WebServer `json:"host"`
}

// StartSearch to brute force the sitweb
// param NetRequest (struct)
func StartSearch(netreq NetRequest){

	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	if netreq.Proxy != "" {
		urlProxy, err := url.Parse(netreq.Proxy)
		Printerr(err, "Fuxe:url.Parse")
		transport.Proxy = http.ProxyURL(urlProxy)
	}
	if netreq.Tor {
		transport.Dial = ThrowTor().Dial
	}
	allPath := ReadFromFile(netreq.Wordlist)
	pathLength := len(allPath)
	if pathLength == 0 {
		Bad("the file is empty!")
		os.Exit(1)
	}
	Info(fmt.Sprintf("Wordlist size: %d / Extensions:%s\n", pathLength, netreq.Ex))
	
	client := &http.Client{Transport: transport}
	req := &http.Request{
		Method: "GET",
		Header: map[string][]string{
			"Cookie":     {netreq.Cookie},
			"User-Agent": {netreq.UserAgent},
		},
	}
	for i := 0; i < pathLength; i++ {
		t <- 0
		wg.Add(1)
		req.URL, _ = url.Parse(netreq.Host + allPath[i])
		go doRequest(&wg, req, *client, i, netreq.Ex)
	}
	wg.Wait()
	jsonF, _ := json.Marshal(webserver)
	timenow := time.Now().Format("2006-01-02-15-04-05")
	filePath := "data/results/" + netreq.ResultFile + strings.Split(netreq.ResultFile, "/")[0] + "-" + timenow + ".json"
	WriteToFile(filePath, fmt.Sprintf("%+v\n", string(jsonF)))

}

// DoRequest to make request
// param *http.Request, http.Client
func doRequest(wg *sync.WaitGroup, req *http.Request, client http.Client, i int, ex []string) {
	defer wg.Done()
	response, err := client.Get(req.URL.String())
	Printerr(err, fmt.Sprintf("DoRequest: %s", req.URL))
	wb := WebServer{
		ID:     i,
		URL:    req.URL.String(),
		Status: response.StatusCode,
		Length: ByteConverter(response.ContentLength),
	}
	webserver.WebServers = append(webserver.WebServers, wb)
	ShowOutput(response.StatusCode, ByteConverter(response.ContentLength), req.URL.String())	
	<-t
}

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

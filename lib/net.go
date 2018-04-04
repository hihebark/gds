package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var (
	waitRequest sync.WaitGroup
	mutex       = &sync.Mutex{}
)

//NetRequest for Request data.
type NetRequest struct {
	Host      string
	Proxyfile string
	Wordlist  string
	UserAgent string
	Cookie    string
	Ex        []string
	Proxy     string
	Tor       bool
}

//CheckConnectivty check if the provided host is up or not.
func CheckConnectivty(host string) int {

	resp, err := http.Get(host)
	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
	return resp.StatusCode

}

//DoRequest to make request and return status, content-length (string, int, int64)
func DoRequest(req *http.Request, client http.Client) {

	response, err := client.Get(req.URL.String())
	Printerr(err, fmt.Sprintf("MakeRequest: %s", req.URL))
	ShowOutput(response.StatusCode, int64(GetLenBody(req)), req.URL.String())
	//return req.URL.String(), response.StatusCode, response.ContentLength

}

//ByteConverter convert length to bytes, KB, MB, GB, TB.
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

//Fuxe to brute force the host
func Fuxe(netreq NetRequest) {

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
	if len(allPath) == 0 {
		Bad("the file is empty!")
		os.Exit(1)
	}
	Info(fmt.Sprintf("Wordlist size: %s / Extensions:%s\n", CountLine(netreq.Wordlist), netreq.Ex))
	waitRequest.Add(len(allPath))
	murl, _ := url.ParseRequestURI(netreq.Host)
	client := &http.Client{Transport: transport}
	req := &http.Request{
		Method: "GET",
		Header: map[string][]string{
			"Cookie":     {netreq.Cookie},
			"User-Agent": {netreq.UserAgent},
		},
	}
	for i := 0; i < len(allPath); i++ {

		go func(i int) {

			defer waitRequest.Done()
			mutex.Lock()
			murl.Path = allPath[i]
			req.URL = murl
			DoRequest(req, *client)
			mutex.Unlock()
			if !strings.HasSuffix(req.URL.String(), "/") && (len(netreq.Ex) >= 1 && netreq.Ex[0] != "") {
				for _, ext := range netreq.Ex {
					mutex.Lock()
					req, _ := http.NewRequest("GET", req.URL.String()+"."+ext, nil)
					DoRequest(req, *client)
					mutex.Unlock()
				}
			}

		}(i)
	}
	waitRequest.Wait()
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

//GetLenBody get the length of the body
func GetLenBody(req *http.Request) int {
	data, err := GetBody(req)
	if err != nil {
		Printerr(err, "GetLenBody:")
	}
	return len(data)
}

//ThrowTor activate the the app to go throw Tor
func ThrowTor() proxy.Dialer {
	torurl, err := url.Parse("socks5://127.0.0.1:9050")
	Printerr(err, "ThrowTor:url.Parse")
	dialer, err := proxy.FromURL(torurl, proxy.Direct)
	Printerr(err, "ThrowTor:proxy.FromURL")
	return dialer
}

//ShowOutput
func ShowOutput(status int, length int64, url string) {
	switch {
	case status >= 200 && status < 299:
		Say(LIGHTGREEN, fmt.Sprintf("%d - %s\t - \t%s", status, ByteConverter(length), url))
	case status >= 300 && status < 399:
		Say(LIGHTBLUE, fmt.Sprintf("%d - %s\t - \t%s", status, ByteConverter(length), url))
	case status >= 400 && status < 500:
		Say(LIGHTRED, fmt.Sprintf("%d - %s\t - \t%s", status, ByteConverter(length), url))
	}
}

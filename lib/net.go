package core

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

//DoRequest to make request and return status, content-length
func DoRequest(req *http.Request, client http.Client) (int, int64) {

	response, err := client.Do(req)
	Printerr(err, fmt.Sprintf("MakeRequest: %s", req.URL))
	return response.StatusCode, response.ContentLength

}

//MakeRequest to make request and return status, content-length
//func MakeRequest(host string, req *http.Request, client http.Client) (int, int64) {

////	if netreq.Cookie != "" {
////		req.Header.Set("Cookie", netreq.Cookie)
////	}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Fatalln("MakeRequest: ", err, host)
//		os.Exit(0)
//	}
//	return resp.StatusCode, resp.ContentLength

//}

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
			murl.Path = allPath[i]
			urlpath := murl.String()
			req.URL = murl
			status, length := DoRequest(req, *client)
			ShowOutput(status, length, urlpath)
			if !strings.HasSuffix(urlpath, "/") && len(netreq.Ex) != 1 {
				for _, ext := range netreq.Ex {
					req, _ := http.NewRequest("GET", urlpath+"."+ext, nil)
					status, length := DoRequest(req, *client)
					ShowOutput(status, length, urlpath+"."+ext)
				}
			}
		}(i)
	}
	waitRequest.Wait()
}

//GetBody fetch the body
func GetBody(netreq NetRequest) {

	client := &http.Client{}
	url, _ := url.Parse(netreq.Host)
	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Set("Cookie", netreq.Cookie)
	request.Header.Set("User-Agent", netreq.UserAgent)
	Printerr(err, "GetBody:http.NewRequest")
	response, err := client.Do(request)
	Printerr(err, "GetBody:client.Do")
	data, err := ioutil.ReadAll(response.Body)
	Printerr(err, "GetBody:ioutil.ReadAll")
	fmt.Println(string(data))
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

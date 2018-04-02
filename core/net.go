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

var waitg sync.WaitGroup

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

//MakeRequest to make request and return status, content-length
func MakeRequest(host string, req *http.Request, client http.Client, netreq NetRequest) (int, int64) {

	if netreq.Cookie != "" {
		req.Header.Set("Cookie", netreq.Cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("MakeRequest: ", err, host)
		os.Exit(0)
	}
	return resp.StatusCode, resp.ContentLength

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
	Info(fmt.Sprintf("File count: %d", len(allPath)))
	waitg.Add(len(allPath))
	murl, _ := url.ParseRequestURI(netreq.Host)
	client := &http.Client{Transport: transport}
	for i := 0; i < len(allPath); i++ {

		go func(i int) {

			defer waitg.Done()
			murl.Path = allPath[i]
			urlpath := murl.String()
			req, _ := http.NewRequest("GET", urlpath, nil)
			req.Header.Set("User-Agent", netreq.UserAgent)
			status, length := MakeRequest(urlpath, req, *client, netreq)

			switch {
			case status >= 200 && status < 299:
				Say(GREEN, fmt.Sprintf("Status: %d - %s\t\t%s",
					status, ByteConverter(length), urlpath))
			case status >= 300 && status < 399:
				Say(LIGHTRED, fmt.Sprintf("Status: %d - %s\t\t%s",
					status, ByteConverter(length), urlpath))
			case status >= 400 && status < 500:
				Say(ORANGE, fmt.Sprintf("Status: %d - %s\t\t%s",
					status, ByteConverter(length), urlpath))
			}
			if !strings.HasSuffix(urlpath, "/") && len(netreq.Ex) != 0 {
				for _, ext := range netreq.Ex {
					req, _ := http.NewRequest("GET", urlpath+"."+ext, nil)
					mstatus, mlength := MakeRequest(urlpath+"."+ext, req, *client, netreq)
					switch {
					case mstatus >= 200 && mstatus < 299:
						Say(GREEN, fmt.Sprintf("# Status: %d - %s\t\t%s",
							mstatus, ByteConverter(mlength), urlpath))
					case mstatus >= 300 && mstatus < 399:
						Say(LIGHTRED, fmt.Sprintf("# Status: %d - %s\t\t%s",
							mstatus, ByteConverter(mlength), urlpath))
					case mstatus >= 400 && mstatus < 500:
						Say(ORANGE, fmt.Sprintf("# Status: %d - %s\t\t%s",
							mstatus, ByteConverter(mlength), urlpath))
					}
				}
			}
		}(i)
	}
	waitg.Wait()
}

//GetBody fetch the body
func GetBody(netreq NetRequest) {

	//fixedURL, err := url.Parse(netreq.Proxy)
	//Printerr(err, "GetBody:url.Parse")
//	client := &http.Client{
//		Transport: &http.Transport{
//			Proxy: http.ProxyURL(fixedURL),
//		},
//	}
	client := &http.Client{}
	url, _ := url.Parse(netreq.Host)
	request, err := http.NewRequest("GET", url.String(), nil)
	request.Header.Set("Cookie", netreq.Cookie)
	request.Header.Set("User-Agent", netreq)
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

//Printerr print error message
func Printerr(err error, fromwhere string) {
	if err != nil {
		Bad(fmt.Sprintf("%s : %v", fromwhere, err))
	}
}

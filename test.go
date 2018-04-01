package main

import (
	"fmt"
	"github.com/hihebark/godirsearch/core"
	"net/http"
	"net/url"
	"sync"
)

//SayTest say test
func SayTest() {
	fmt.Println("Just for Testing...")
	line := core.ReadFromFile("test.txt")
	var waitg sync.WaitGroup
	host := "http://ouedkniss.com/"
	waitg.Add(len(line))
	murl, _ := url.ParseRequestURI(host)
	client := &http.Client{}
	for i := 0; i < len(line); i++ {
		go func(i int) {
			defer waitg.Done()
			murl.Path = line[i]
			urlpath := murl.String()
			req, _ := http.NewRequest("GET", urlpath, nil)
			status, length := core.MakeRequest(urlpath, req, *client)
			fmt.Printf("status: %d\t-\tlength: %d\n", status, length)
		}(i)
	}
	waitg.Wait()
	/*Proxy test
	  req := core.NetRequest{
	              Host:"http://httpbin.org/get",
	              Proxy:"http://171.97.67.88:3128",
	      }
	  core.GetBody(req)*/
	/*Tor
	  transport := &http.Transport{}
	  transport.Dial = core.ThrowTor().Dial
	*/
	/*excute shell script
	  s, err := core.Execute("/bin/bash", []string{"core/grepproxylist.sh"})
	  core.Printerr(err, "test:core.Execute")
	  fmt.Println(strings.Split(s, "\n"))
	*/
}

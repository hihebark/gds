package main

import (
	"fmt"
	"sync"
	"github.com/hihebark/godirsearch/core"
)
//SayTest say test
func main() {
	fmt.Println("Just for Testing...")
	line := core.ReadFromFile("test.txt")
	var waitg sync.WaitGroup
	waitg.Add(len(line))
	for i := 0; i < len(line); i++ {
		go func(i int) {
			defer waitg.Done()
			fmt.Printf("%d path: %s\n", i, line[i])
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

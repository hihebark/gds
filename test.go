package main

import (
	"fmt"
	//"github.com/hihebark/godirsearch/core"
)
//SayTest say test
func SayTest() {
	fmt.Println("Just for Testing...")
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

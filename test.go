package main

import (
    //"fmt"
    "github.com/hihebark/godirsearch/core"
)

func main(){

    req := core.NetRequest{
                Host:"http://httpbin.org/get",
                Proxy:"http://171.97.67.88:3128",
        }
    core.GetBody(req)
}

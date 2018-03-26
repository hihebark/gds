package main

/**
* DONE regex for verifing url
* DONE check the connecivity to the url
* TODO set Cookies
* TODO generating json file of the result +
* TODO a log file for debbuging
* TODO use goroutine just for fun and for learning //for thread
**/

import (
        "os"
        "fmt"
        "flag"
        "bufio"
        "regexp"
        "strings"
        "net/url"
        "net/http"

        "github.com/hihebark/godirsearch/core"
)

var (
    tor         *bool
    host        *string
    proxy       *string
    thread      *int
    version     *string
    proxyfile   *string
    
    urlpath     string
)

//type Cons struct {
//    status      *int
//    length      *int64
//    hostpath    *string
//    time        *string
//}

func init(){
    tor         = flag.Bool("tor", false, "Brutforce using Tor")
    host        = flag.String("host", "", "Host to brutforce")
    proxy       = flag.String("proxy", "", "Use a proxy to brutforce")
    thread      = flag.Int("thread", 4, "Number of thread")
    version     = flag.String("version", "v ~0.0.1", "print version")
    proxyfile   = flag.String("proxyfile", "", "Use a proxy file")
    
}

func main() {
    fmt.Println("\tGoDirSearch -v ~0.0.2\n")
    flag.Parse()
    /***************************************************************************
    * Best regex `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`
    * http://www.golangprograms.com/golang-package-examples/regular-expression-to-extract-domain-from-url.html
    */
    re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
    status := core.CheckConnectivty(*host)
    if (re.MatchString(*host) && (status >= 200 && status < 300)){
        if (!strings.HasSuffix(*host, "/")){
            *host += "/"
        }
        fmt.Println("\033[92mConnection to the target Ok!\033[0m",status)
        file, err := os.Open(*proxyfile)
        if err != nil {
            fmt.Printf("error opening file: %v\n",err)
            os.Exit(1)
        }
        reader := bufio.NewReader(file)
        path, err := core.Readln(reader)
        client := &http.Client{}
        
        murl, err := url.ParseRequestURI(*host)
        if(err != nil){
            fmt.Println(err)
        }
        
        for err == nil {
        
            murl.Path = path
            urlpath = murl.String()
            req, _ := http.NewRequest("HEAD", urlpath, nil)
            req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
            mstatus, mlength := core.MakeRequest(*host+path, req, *client)
            fmt.Printf("Status: %d - %s\t\tPath: %s\n", mstatus, core.ByteConverter(mlength), *host+path)
            path,err = core.Readln(reader)
        
        }
    } else {
        fmt.Println("\033[91mHost not recheable status:\033[0m", status)
    }
    
}


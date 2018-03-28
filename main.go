package main

/**
* DONE regex for verifing url
* DONE check the connecivity to the url
* TODO set Cookies
* DONE set extension
* TODO generating json file of the result +
* TODO a log file for debbuging
* TODO use thread
* TODO why no an UI Just saying
**/

import (
        "os"
        "fmt"
        "flag"
//        "bufio"
        "regexp"
        "strings"
//        "net/url"
//        "net/http"
        "github.com/hihebark/godirsearch/core"
)

const version string = "0.0.2"

var (
    tor *bool
    thread *int
    host, proxy, cookie, wordlist, proxyfile, userAgent, ex *string
)

func init(){

    ex          = flag.String("ex", "", "separate with coma like php,txt ...")
    tor         = flag.Bool("tor", false, "Brutforce using Tor")
    host        = flag.String("host", "", "Host to brutforce")
    proxy       = flag.String("proxy", "", "Use a proxy to brutforce")
    thread      = flag.Int("thread", 4, "Number of thread")
    cookie      = flag.String("cookie", "", "cookie")
    wordlist    = flag.String("worlist", "test.txt", "wordlist to brutforce")
    proxyfile   = flag.String("proxyfile", "", "Use a proxy file")
    userAgent   = flag.String("useragent", "Golang_Spider_Bot/3.0", "userAgent")
    
}

func main() {
    
    fmt.Printf("\tGoDirSearch \033[92m~%s\n\033[0m", version)
    flag.Parse()
    if(*host == ""){
        fmt.Printf("No host argument found? add -host http://examples.com/ \n")
        os.Exit(0)
    }
    
    /***************************************************************************
    * Best regex `^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`
    * http://www.golangprograms.com/golang-package-examples/regular-expression-to-extract-domain-from-url.html
    ****************************************************************************/
    
    re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
    status := core.CheckConnectivty(*host)
    if (re.MatchString(*host) && (status >= 200 && status < 300)){
    
        if (!strings.HasSuffix(*host, "/")){
            *host += "/"
        }
        fmt.Println("\033[92mConnection to the target Ok!\033[0m",status)
        req := core.NetRequest{
                Host:*host,
                Proxyfile:*proxyfile,
                Wordlist:*wordlist,
                UserAgent:*userAgent,
                Cookie:*cookie,
                Ex:strings.Split(*ex, ","),
                Proxy:*proxy,
                Tor:*tor,
        }
        //core.Fuxe(req)
        core.GetBody(req)
    
    } else {
        fmt.Println("\033[91mHost not recheable status:\033[0m", status)
    }
    
}


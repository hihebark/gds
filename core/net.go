package core

import(
    "io/ioutil"
    "net/http"
    "net/url"
    "bufio"
    "strings"
    "time"
    "log"
    "fmt"
    "os"
    
    "golang.org/x/net/proxy"
)

var urlpath string

type NetRequest struct{
    Host        string
    Proxyfile   string
    Wordlist    string
    UserAgent   string
    Cookie      string
    Ex          []string
    Proxy       string
    Tor         bool
}

func CheckConnectivty(host string) (int){

    resp, err := http.Get(host)
    if (err != nil){
        log.Fatalln(err)
        os.Exit(0)
    }
    return resp.StatusCode

}

func MakeRequest(host string, req *http.Request, client http.Client) (int, int64){

    resp, err := client.Do(req)
    if (err != nil){
        log.Fatalln("MakeRequest: ",err, host)
        os.Exit(0)
    }
    return resp.StatusCode, resp.ContentLength

}

func ByteConverter(length int64) string{
    mbyte := []string{"bytes", "KB", "MB", "GB", "TB"}
    if (length == -1){
            return "0 byte"
    }
    for _, x := range mbyte{
        if (length < 1024.0){
            return fmt.Sprintf("%3.1d %s", length, x)
        }
        length = length / 1024.0
    }
    return ""
}

func Fuxe(netreq NetRequest) {

    transport := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    30 * time.Second,
        DisableCompression: true,
    }
    if (netreq.Proxy != ""){
        urlProxy, err := url.Parse(netreq.Proxy)
        Printerr(err, "Fuxe:url.Parse")
        transport.Proxy = http.ProxyURL(urlProxy)
    }
    if netreq.Tor {
        transport.Dial = ThrowTor().Dial
    }
    file, err := os.Open(netreq.Wordlist)
    Printerr(err, "Fuxe:os.Open")
    murl, err := url.ParseRequestURI(netreq.Host)
    Printerr(err, "Fuxe:url.ParseRequestURI")
    reader := bufio.NewReader(file)
    path, err := Readln(reader)
    client := &http.Client{ Transport: transport }
    for err == nil {
    
        murl.Path = path
        urlpath = murl.String()
        req, _ := http.NewRequest("GET", urlpath, nil)
        req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
        mstatus, mlength := MakeRequest(urlpath, req, *client)
        fmt.Printf("Status: %d - %s\t\tPath: %s\n", mstatus, ByteConverter(mlength), urlpath)
        path, err = Readln(reader)
        if (!strings.HasSuffix(urlpath, "/") && len(netreq.Ex) != 0){
            for _, ext := range netreq.Ex {
                req, _ := http.NewRequest("GET", urlpath+"."+ext, nil)
                mstatus, mlength := MakeRequest(urlpath+"."+ext, req, *client)
                fmt.Printf("Status: %d - %s\t\tPath: %s\n", mstatus, ByteConverter(mlength), urlpath+"."+ext)
            }
        }
    
    }

}

func GetBody(netreq NetRequest){

    fixedURL, err := url.Parse(netreq.Proxy)
    Printerr(err, "GetBody:url.Parse")
    client := &http.Client{
        Transport:&http.Transport{
            Proxy:http.ProxyURL(fixedURL),
        },
    }
    url, _ := url.Parse(netreq.Host)
    request, err := http.NewRequest("GET", url.String(), nil)
    Printerr(err, "GetBody:http.NewRequest")
    response, err := client.Do(request)
    Printerr(err, "GetBody:client.Do")
    data, err := ioutil.ReadAll(response.Body)
    Printerr(err, "GetBody:ioutil.ReadAll")
    fmt.Println(string(data))
}

func ThrowTor() proxy.Dialer{
    torurl, err := url.Parse("socks5://127.0.0.1:9050")
    Printerr(err, "ThrowTor:url.Parse")
    dialer, err := proxy.FromURL(torurl, proxy.Direct)
    Printerr(err, "ThrowTor:proxy.FromURL")
    return dialer
}

func Printerr(err error, fromwhere string) {
    if err != nil {
        fmt.Printf("%s : %v", fromwhere, err)
    }
}


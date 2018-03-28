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
        if err != nil {
            fmt.Println("urlProxy url.Parse: ", err)
        }
        transport.Proxy = http.ProxyURL(urlProxy)
    }
    if netreq.Tor {
        transport.Dial = ThrowTor().Dial
    }
    file, err := os.Open(netreq.Wordlist)
    if err != nil {
        fmt.Printf("error opening file: %v\n",err)
        os.Exit(1)
    }
    murl, err := url.ParseRequestURI(netreq.Host)
    if(err != nil){
        fmt.Println("url.ParseRequestURI:", err)
    }
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
        if (!strings.HasSuffix(urlpath, "/") && len(netreq.Ex) !=0){
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
    if (err != nil){
        fmt.Println("proxy not in use ",err)
    }
    client := &http.Client{
        Transport:&http.Transport{
            Proxy:http.ProxyURL(fixedURL),
        },
    }
    url, _ := url.Parse(netreq.Host)
    
    request, err := http.NewRequest("GET", url.String(), nil)
    if err != nil {
        log.Println(err)
    }
    response, err := client.Do(request)
    if err != nil {
        log.Println(err)
    }
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
    }
    //printing the response
    fmt.Println(string(data))
}

func ThrowTor() proxy.Dialer{

    torurl, _ := url.Parse("socks5://127.0.0.1:9050")
    dialer, _ := proxy.FromURL(torurl, proxy.Direct)
    
    return dialer
}


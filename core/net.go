package core

import(
    "net/http"
    "net/url"
    "bufio"
    "log"
    "fmt"
    "os"
)

var urlpath string

type T struct {
    S  string
    is []int
}

type NetRequest struct{
    Host        string
    Proxyfile   string
    Wordlist    string
    UserAgent   string
    Cookie      string
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
    client := &http.Client{}
    for err == nil {
    
        murl.Path = path
        urlpath = murl.String()
        req, _ := http.NewRequest("HEAD", urlpath, nil)
        req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
        mstatus, mlength := MakeRequest(netreq.Host+path, req, *client)
        fmt.Printf("Status: %d - %s\t\tPath: %s\n", mstatus, ByteConverter(mlength), netreq.Host+path)
        path, err = Readln(reader)
    
    }
}



























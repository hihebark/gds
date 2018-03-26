package core

import(
    "net/http"
    "log"
    "fmt"
)

func CheckConnectivty(host string) (int){

    resp, err := http.Get(host)
    if (err != nil){
        log.Fatalln(err)
        return 0
    }
    return resp.StatusCode

}

func MakeRequest(host string, req *http.Request, client http.Client) (int, int64){

    resp, err := client.Do(req)
    if (err != nil){
        log.Fatalln("MakeRequest: ",err, host)
    }
    return resp.StatusCode, resp.ContentLength

}

func ByteConverter(length int64) string{
    mbyte := []string{"bytes", "KB", "MB", "GB", "TB"}
    for _, x := range mbyte{
        if (length == -1){
            return "0 bytes"
        }
        if (length < 1024.0){
            return fmt.Sprintf("%3.1d %s", length, x)
        }
        length = int64(length) / 1024.0
    }
    return ""
}

//func GetBodyLength(host string)(io.ReadCloser, error){

//    req, err := http.Get(host)
//    return ioutil.ReadAll(req.Body), err
//    
//}

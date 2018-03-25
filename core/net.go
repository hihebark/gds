package core

import(
    "net/http"
    //"io/ioutil"
    "log"
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
    return resp.StatusCode, req.ContentLength

}


//func GetBodyLength(host string)(io.ReadCloser, error){

//    req, err := http.Get(host)
//    return ioutil.ReadAll(req.Body), err
//    
//}

package lib

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"html/template"
)

// MyMux wtf is this <<<
type MyMux struct {
	
}

//ServeHTTP hundle results route
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/results" {
		ShowResult(w, r)
		return
	}
	http.Redirect(w, r, "/results", http.StatusFound)
	return
}

//ShowResult show the result
func ShowResult(w http.ResponseWriter, r *http.Request) {

	data := WebServerslice{}
	jsonfile, err := os.Open("data/results/www.exemple.com/www.exemple.com+2018-04-09-18-26-02.json")
	defer jsonfile.Close()
	Printerr(err, "gui:ShowResult:os.Open: ")
	jsonParser := json.NewDecoder(jsonfile)
	if err = jsonParser.Decode(&data); err != nil {
		Printerr(err, "gui:ShowResult:Parsing config file error")
	}
	
	htmlTemplate := template.New("Godirsearch")
//	htmlTemplate, err = htmlTemplate.ParseFiles("data/web/index.html")
	htmlTemplate, err = htmlTemplate.Parse(ReturnStringFile("data/web/index.html"))
	Printerr(err, "gui:ShowResult:htmlTemplate.ParseFiles:")
	//fmt.Printf("%+v\n",data.WebServers)
	htmlTemplate.Execute(w, data.WebServers)
	//fmt.Fprintf(w, fmt.Sprintf("%+v", data.WebServers))
	//fmt.Printf("%+d\n", r)
}

//StartListning start listning to the given port
func StartListning() {
	Info("Stating server on http://localhost:9011/results | http://[::1]:9011/results")
	mux := &MyMux{}
	err := http.ListenAndServe(":9011", mux)
	if err != nil {
		fmt.Printf("StartListning:error: %s\n", err)
	}
}

package lib

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
)

// ServeMux for concurrency
type ServeMux struct {
	mutex sync.RWMutex
}

//ServeHTTP hundle results route
func (mutex *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch{
		case r.URL.Path == "/":
			mutex.mutex.RLock()
			defer mutex.mutex.RUnlock()
			ShowResult(w, r)
			return
		case r.URL.Path == "/Logo.png":
			http.ServeFile(w, r, "data/web/Logo.png")
			return
		case r.URL.Path == "/results"://case len(strings.Split(r.URL.Path, "/")) == 2:
			ShowResultsFile(w, r, "data/results")
			return
		default:
			http.Redirect(w, r, "/", http.StatusFound)
			return
	}

}

//ShowResultsFile see if the directory is in the data/results and show all json files in it
func ShowResultsFile (w http.ResponseWriter, r *http.Request, folder string){

	if Existe(folder) {
		list := GetListFile(folder)
		var listfile []string
		for _, value := range list {
			if value == ""{
				break
			}
			f, err := os.Stat(folder+"/"+value)
			Printerr(err, "ShowResultsFile:os.Stat")
			if f.Mode().IsDir() {
				listfile = append(listfile, value)
			}
		}
		htmlTemplate := template.New("index.html")
		htmlTemplate, err := htmlTemplate.ParseFiles("data/web/index.html")
		Printerr(err, "gui:ShowResult:htmlTemplate.ParseFiles")
		htmlTemplate.Execute(w, listfile)
	}else{
		Bad("File don't existe.")
	}
	
}

//ShowResult show the result
func ShowResult(w http.ResponseWriter, r *http.Request) {

	data := WebServerslice{}
	jsonfile, err := os.Open("data/results/www.ouedkniss.com/www.ouedkniss.com+2018-04-09-18-26-02.json")
	defer jsonfile.Close()
	Printerr(err, "gui:ShowResult:os.Open")
	jsonParser := json.NewDecoder(jsonfile)
	if err = jsonParser.Decode(&data); err != nil {
		Printerr(err, "gui:ShowResult:Parsing config file error")
	}

	htmlTemplate := template.New("result.html")
	htmlTemplate, err = htmlTemplate.ParseFiles("data/web/result.html")
	Printerr(err, "gui:ShowResult:htmlTemplate.ParseFiles")
	htmlTemplate.Execute(w, data.WebServers)
}

//StartListning start listning to the given port
func StartListning() {
	Info("Stating server on http://localhost:9011/ | http://[::1]:9011/")
	mux := &ServeMux{}
	err := http.ListenAndServe(":9011", mux)
	if err != nil {
		fmt.Printf("StartListning:error: %s\n", err)
	}
}

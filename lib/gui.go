package lib

import (
	"fmt"
	"net/http"
)

//MyMux wtf is this <<<
type MyMux struct {
	
}

//ServeHTTP hundle route
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
	fmt.Fprintf(w, fmt.Sprintf("%s", ReturnStringFile("data/web/index.html")))
}

//StartListning start listning to the given port
func StartListning(mux *MyMux) {
	Info("Stating server on http://localhost:9011/results")
	Info("Stating server on http://[::1]:9011/results")
	err := http.ListenAndServe(":9011", mux)
	if err != nil {
		fmt.Printf("StartListning:error: %s\n", err)
	}
}

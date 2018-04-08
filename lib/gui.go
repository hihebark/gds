package lib

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}

//ServeHTTP hundle route
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/results" {
		ShowResult(w, r)
		return
	}else {
		http.Redirect(w, r, "/results", http.StatusFound)
	}
	http.NotFound(w, r)
	return
}

//ShowResult show the result
func ShowResult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//StartListning start listning to the given port
func StartListning(mux *MyMux) {
	Info("Stating server on http://localhost:9011")
	err := http.ListenAndServe(":9011", mux)
	if err != nil {
		fmt.Printf("StartListning:error: %s\n", err)
	}
}


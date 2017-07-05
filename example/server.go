package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan/formatter"
	"github.com/shirou/golightan/lexer"
)

type render struct {
	target string
	src    string
}

func Render(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only POST alllowed", http.StatusMethodNotAllowed)
		return
	}

	target := r.URL.Query().Get("target")
	if target == "" {
		target = "c"
	}
	body, _ := ioutil.ReadAll(r.Body)

	pf, err := lexer.Factory(target)
	if err != nil {
		http.Error(w, "undefined lexer", http.StatusBadRequest)
		return
	}

	input := antlr.NewInputStream(string(body))
	tokens, err := pf.Tokenize(input)
	if err != nil {
		http.Error(w, "tokenize failed", http.StatusBadRequest)
		return
	}

	ff, err := formatter.Factory("html", "style")
	if err != nil {
		http.Error(w, "factory failed", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	ff.FormatTokens(w, tokens)
}
func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/render", Render)
	http.HandleFunc("/", Index)

	log.Printf("starting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Articles []Article

type Article struct {
	SNo   int    `json:"Serial Number"`
	Title string `json:"Title"`
	PNo   int    `json:"Page Number"`
}

func main() {
	//defer muxhandleRequests()
	Articles = []Article{
		Article{SNo: 1, Title: "GoLang", PNo: 4},
		Article{SNo: 2, Title: "Python", PNo: 24},
		Article{SNo: 3, Title: "Scala", PNo: 44},
		Article{SNo: 4, Title: "Shell Script", PNo: 64},
		Article{SNo: 5, Title: "Java Script", PNo: 74},
	}
	//muxhandleRequests()
	handleRequests()
	//	muxhandleRequests()

	//r := mux.NewRouter()

}

func website(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Articles)
}

func empty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Dengey Bey ")
	//json.NewEncoder(w).Encode(Articles)
}

func special(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jaffa ")
	fmt.Fprintf(w, "BADKAV")
}

func handleRequests() {
	http.HandleFunc("/", empty)
	http.HandleFunc("/articles", website)
	http.HandleFunc("/special", special)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/*
func muxhandleRequests() {
	//	muxRouter := mux.NewRouter.StrictSlash(true)
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", empty)
	muxRouter.HandleFunc("/jaffa", special)
	muxRouter.HandleFunc("/special", website)
	log.Fatal(http.ListenAndServe(":8082", muxRouter))
}
*/

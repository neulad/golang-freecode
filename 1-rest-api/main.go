package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func main() {
	handleRequests()
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{{Title: "Test", Desc: "Test Desc", Content: "Some contant"}}
	fmt.Println("Endpoinnt Hit: All articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
	fmt.Println("Request  has been  made! ")
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Okay got yur response bummer")
	fmt.Printf("Articles post: %v", r.Body)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", postArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

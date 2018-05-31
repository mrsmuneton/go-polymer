package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {

	// Had to do this because returns svg as text/xml when running on AppEngine: http://goo.gl/hwZSp2
	mime.AddExtensionType(".svg", "image/svg+xml")

	r := mux.NewRouter()
	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/posts", Posts)
	sr.HandleFunc("/constitution", SendConstitution)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/{rest:.*}", handler)
	http.Handle("/", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("path:", r.URL.Path)
	http.ServeFile(w, r, "static/"+r.URL.Path)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

type Post struct {
	Uid      int    `json:"uid"`
	Text     string `json:"text"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Favorite bool   `json:"favorite"`
}

type Constitution struct {
	Uid    string `json:"uid"`
	Text   string `json:"text"`
	Number string `json:"number"`
	Part   string `json:"part"`
}

func Posts(w http.ResponseWriter, r *http.Request) {
	posts := []Post{}
	// you'd use a real database here
	file, err := ioutil.ReadFile("posts.json")
	if err != nil {
		log.Println("Error reading posts.json:", err)
		panic(err)
	}
	fmt.Printf("file: %s\n", string(file))
	err = json.Unmarshal(file, &posts)
	if err != nil {
		log.Println("Error unmarshalling posts.json:", err)
	}

	val, err := json.Marshal(posts)
	if err != nil {
		ReturnError(w, err)
		return
	}
	fmt.Fprint(w, string(val))
}

func SendConstitution(w http.ResponseWriter, r *http.Request) {
	println(" sending constitution ")
	constitution := []Constitution{}
	// you'd use a real database here
	file, err := ioutil.ReadFile("Constitution.json")
	if err != nil {
		log.Println("Error reading Constitution.json:", err)
		panic(err)
	}
	fmt.Printf("file: %s\n", string(file))
	err = json.Unmarshal(file, &constitution)
	if err != nil {
		log.Println("Error unmarshalling Constitution.json:", err)
	}

	val, err := json.Marshal(constitution)
	// fmt.Printf("val", var)
	if err != nil {
		ReturnError(w, err)
		return
	}
	fmt.Fprint(w, string(val))
}

func ReturnError(w http.ResponseWriter, err error) {
	fmt.Fprint(w, "{\"error\": \"%v\"}", err)
}

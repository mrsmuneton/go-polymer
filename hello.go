package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {

	// Had to do this because returns svg as text/xml when running on AppEngine: http://goo.gl/hwZSp2
	mime.AddExtensionType(".svg", "image/svg+xml")

	r := mux.NewRouter()
	sr := r.PathPrefix("/api").Subrouter()
	sr.HandleFunc("/posts", Posts)
	sr.HandleFunc("/constitution", SendConstitution)
	r.HandleFunc("/chain", Chain)
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

func chainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "chain")
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

type Block struct {
	Uid          string    `json:"number"`
	Timestamp    time.Time `json:"timestamp"`
	TextValue    string    `json:"textvalue"`
	PreviousHash string    `json:previoushash`
	Hash         string    `json:hash`
}

func Chain(w http.ResponseWriter, r *http.Request) {
	blocks := []Block{}
	file, err := ioutil.ReadFile("blocks.json")
	if err != nil {
		log.Println("Error reading blocks.json:", err)
		panic(err)
	}
	fmt.Printf("file: %s\n", string(file))
	err = json.Unmarshal(file, &blocks)
	if err != nil {
		log.Println("Error unmarshalling blocks.json:", err)
	}

	for i, _ := range blocks {
		blocks[i].Timestamp = time.Now()
		sha256hash(&blocks[i])
		fmt.Println(blocks[i])
	}

	val, err := json.Marshal(blocks)
	if err != nil {
		ReturnError(w, err)
		return
	}

	// fmt.Printf("value: %s\n", string(val))
	// fmt.Println(reflect.TypeOf(val.))

	fmt.Fprint(w, string(val))
}

func sha256hash(thisblock *Block) {
	h := sha256.New()
	h.Write([]byte(thisblock.TextValue))
	thisblock.Hash = hex.EncodeToString(h.Sum(nil))
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

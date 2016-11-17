package Func

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/albertrizcha/makegolangapi/Model"

	"time"

	"github.com/gorilla/mux"
)

var Articles = Model.Articles{
	Model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content", Due: time.Now()},
	Model.Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content", Due: time.Now()},
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home")
}
func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(Articles)
}

func ReturnOneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	result := []byte("null data")
	if key == "1" {
		result, _ = json.Marshal(Articles[0])
	} else if key == "2" {
		result, _ = json.Marshal(Articles[1])
	}

	fmt.Fprintln(w, "Key: "+key)
	fmt.Fprintln(w, string(result))
}

func ReturnOneArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	var1 := vars["var1"]
	var2 := vars["var2"]

	fmt.Println("Var 1: " + var1)
	fmt.Println("Var 2: " + var2)
	fmt.Fprintf(w, "Key: "+key)
}

func TestJSON(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Model.Article
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	t.Title = t.Title + " POST"
	t.Due = time.Now()

	p, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(p))
	json.NewEncoder(w).Encode(t)
}

func TestJson(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t Model.Test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(t.Test)
}

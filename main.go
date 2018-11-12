package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/felipedacs/yugo-api/yfl"
	"github.com/felipedacs/yugo-api/yutils"
)

const (
	pathPosts = "content" + string(filepath.Separator) + "post"
)

func main() {
	defer runServer()

	yfl.IniciaFoldersHugo(pathPosts)

	//d1 := []byte("hello\ngo\n")
	// criar
	//err := ioutil.WriteFile("teste/post/aaaa.txt", d1, os.ModePerm)
	//check(err)

}

func runServer() {
	//rotas
	http.HandleFunc("/api", handler)
	http.HandleFunc("/api/posts", posts)
	http.HandleFunc("/api/post/", post)

	http.ListenAndServe(":8083", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	fmt.Println("acesso /api")
	w.Write([]byte("/api"))
}

func posts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	fmt.Println("acesso /api/posts")
	posts := yfl.ListaPosts(pathPosts)

	js, err := json.Marshal(posts)
	yutils.Check(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	postNome := r.URL.Path[len("/api/post/"):]
	strPostNome := string(postNome)
	fmt.Println("acesso /api/post/" + strPostNome)

	post, err := yfl.LePost(pathPosts, strPostNome)
	yutils.Check(err)

	js, rr := json.Marshal(post)
	yutils.Check(rr)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/felipedacs/yugo-api/yfl"
	"github.com/felipedacs/yugo-api/ytr"
	"github.com/felipedacs/yugo-api/yutils"
)

const (
	pathPosts = "content" + string(filepath.Separator) + "post"
	port      = ":8083"
)

type configResult struct {
	Repo    string `json:"repo"`
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
}

func init() {
	// log.SetOutput(ioutil.Discard) // desabilitar logs
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}

func main() {
	// go ytr.IniciaHugo()
	defer runServer()
	yfl.IniciaFoldersHugo()
	yfl.IniciaArquivoConfig()
}

func runServer() {
	//rotas
	http.HandleFunc("/api", handler)
	http.HandleFunc("/api/posts", posts)
	http.HandleFunc("/api/post/", post)
	http.HandleFunc("/api/config", config)
	http.HandleFunc("/api/newpost", newpost)
	http.HandleFunc("/api/savepost", savepost)
	http.HandleFunc("/api/renamepost", renamepost)
	http.HandleFunc("/api/deletepost", deletepost)
	http.Handle("/", http.FileServer(http.Dir("./ui")))

	fmt.Println("Não feche essa janela!")
	fmt.Printf("Acesse no navegador localhost%v\n", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)

	log.Println(r.Method, " -> ", r.URL)
	w.Write([]byte("/api"))
}

func posts(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)

	log.Println(r.Method, " -> ", r.URL)
	posts := yfl.ListaPosts()

	js, err := json.Marshal(posts)
	yutils.Check(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func post(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)

	postNome := r.URL.Path[len("/api/post/"):]
	strPostNome := string(postNome)
	log.Println(r.Method, " -> ", r.URL, strPostNome)

	post, err := yfl.LePost(strPostNome)
	yutils.Check(err)

	js, rr := json.Marshal(post)
	yutils.Check(rr)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func config(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	log.Println(r.Method, " -> ", r.URL)

	if r.Method == "POST" {
		var cr configResult
		err := json.NewDecoder(r.Body).Decode(&cr)
		yutils.Check(err)

		yfl.AtualizaConfig(r.Body)
		ytr.PublicaESalva(cr.Usuario, cr.Senha, cr.Repo)
	} else if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		dat, err := ioutil.ReadFile("yugo.json")
		yutils.Check(err)
		fmt.Fprintf(w, string(dat))
	}
}

func newpost(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	log.Println(r.Method, " -> ", r.URL)
	if r.Method == "POST" {
		yfl.NewPost()
	}
}

func savepost(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	log.Println(r.Method, " -> ", r.URL)

	if r.Method == "POST" {
		setupHeaders(w)
		yfl.SavePost(r.Body)
	}
}

func renamepost(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	log.Println(r.Method, " -> ", r.URL)
	if r.Method == "POST" {
		yfl.RenamePost(r.Body)
	}
}

func deletepost(w http.ResponseWriter, r *http.Request) {
	setupHeaders(w)
	log.Println(r.Method, " -> ", r.URL)
	if r.Method == "POST" {
		yfl.DeletePost(r.Body)
	}
}

func setupHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") //cors
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
}

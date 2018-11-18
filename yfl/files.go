// yugo files
package yfl

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/felipedacs/yugo-api/yutils"
)

type Post struct {
	Nome     string `json:"nome"`
	NovoNome string `json:"novoNome"`
	Conteudo string `json:"conteudo"`
}

type configResult struct {
	Repo string `json:"repo"`
}

const (
	path = "content" + string(filepath.Separator) + "post"
)

func IniciaArquivoConfig() {
	f, err := os.Create("yugo.json")
	defer f.Close()
	yutils.Check(err)
}

func ListaPosts() []Post {
	var posts []Post
	files, err := ioutil.ReadDir(path)
	yutils.Check(err)

	for _, f := range files {
		tmp := Post{Nome: f.Name()}
		posts = append(posts, tmp)
	}
	return posts
}

func LePost(nomePost string) (Post, error) {
	post := Post{Nome: nomePost}
	posts := ListaPosts()
	for _, p := range posts {
		if post.Nome == strings.Replace(p.Nome, ".md", "", -1) {
			f, err := ioutil.ReadFile(path + string(filepath.Separator) + p.Nome) // just pass the file name
			yutils.Check(err)
			post.Conteudo = string(f)
			return post, nil
		}
	}
	return post, errors.New("erro")
}

func AtualizaConfig(body io.Reader) {
	var cr configResult
	err := json.NewDecoder(body).Decode(&cr)
	yutils.Check(err)

	crJSON, err := json.MarshalIndent(cr, "", "	")
	yutils.Check(err)

	file, err := os.Create("yugo.json")
	yutils.Check(err)
	defer file.Close()
	file.WriteString(string(crJSON))
}

func NewPost() {
	file, err := os.Create("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + "_newpost.md")
	yutils.Check(err)
	defer file.Close()
	file.WriteString("renomeie esse arquivo!")
}

func SavePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	file, err := os.Create("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + post.Nome + ".md")
	yutils.Check(err)
	defer file.Close()
	file.WriteString(post.Conteudo)
}

func RenamePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	os.Rename("content"+string(filepath.Separator)+"post"+string(filepath.Separator)+post.Nome+".md", "content"+string(filepath.Separator)+"post"+string(filepath.Separator)+post.NovoNome+".md")
}

func DeletePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	err = os.Remove("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + post.Nome + ".md")
	yutils.Check(err)
}

func get() {

}

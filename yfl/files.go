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

// Post estrutura do padrão de arquivo post
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

// IniciaArquivoConfig cria ou não yugo.json
func IniciaArquivoConfig() {
	filename := "yugo.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err := os.Create("yugo.json")
		defer f.Close()
		yutils.Check(err)
	}
}

// ListaPosts listagem dos titulos dos arquivos de post sem ".md"
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

// LePost leitura do arquivo para retornar a struct post
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

// AtualizaConfig recria yugo.json
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

// NewPost cria novo arquivo ".md" com título e conteúdo padrão
func NewPost() {
	file, err := os.Create("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + "_newpost.md")
	yutils.Check(err)
	defer file.Close()
	file.WriteString("renomeie esse arquivo!")
}

// SavePost cria novo arquivo de mesmo nome com conteúdo diferente, sobescrevendo o antigo
func SavePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	file, err := os.Create("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + post.Nome + ".md")
	yutils.Check(err)
	defer file.Close()
	file.WriteString(post.Conteudo)
}

// RenamePost renomeia novo post a partir do nome atual para o novo
func RenamePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	os.Rename("content"+string(filepath.Separator)+"post"+string(filepath.Separator)+post.Nome+".md", "content"+string(filepath.Separator)+"post"+string(filepath.Separator)+post.NovoNome+".md")
}

// DeletePost remove post
func DeletePost(body io.Reader) {
	var post Post
	err := json.NewDecoder(body).Decode(&post)
	yutils.Check(err)

	err = os.Remove("content" + string(filepath.Separator) + "post" + string(filepath.Separator) + post.Nome + ".md")
	yutils.Check(err)
}

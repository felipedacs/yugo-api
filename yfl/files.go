// yugo files
package yfl

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/felipedacs/yugo-api/yutils"
)

type Post struct {
	Nome     string
	Conteudo string
}

func ListaPosts(path string) []Post {
	var posts []Post
	files, err := ioutil.ReadDir(path)
	yutils.Check(err)

	for _, f := range files {
		tmp := Post{Nome: f.Name()}
		posts = append(posts, tmp)
	}

	return posts
}

func LePost(path string, nomePost string) (Post, error) {
	post := Post{Nome: nomePost}
	posts := ListaPosts(path)
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

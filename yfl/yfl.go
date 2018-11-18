// package yfl
//
// // go test -v ./yfl
//
// import (
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"testing"
//
// 	"github.com/felipedacs/yugo-api/yfl"
// )
//
// const (
// 	pathPosts  = ".." + string(filepath.Separator) + "content"
// 	yugoConfig = ".." + string(filepath.Separator) + "yugo.json"
// )
//
// func TestMain(m *testing.M) {
// 	setup()
// 	code := m.Run()
//
// 	// shutdown()
// 	os.Exit(code)
// }
//
// func TestIniciaArquivoConfig(t *testing.T) {
// 	// setup()
//
// 	yfl.IniciaArquivoConfig()
//
// 	encontrou := false
//
// 	files, err := ioutil.ReadDir(".." + string(filepath.Separator))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	for _, file := range files {
// 		if file.Name() == "yugo.json" {
// 			encontrou = true
// 		}
// 	}
//
// 	if !encontrou {
// 		t.Error("Arquivo não yugo.json não foi criado")
// 	}
// }
//
// func TestIniciaFoldersHugo(t *testing.T) {
// 	yfl.IniciaFoldersHugo()
// 	encontrou := 0
// 	filepath.Walk(pathPosts, func(path string, info os.FileInfo, err error) error {
// 		if path == ".."+string(filepath.Separator)+"content" || path == ".."+string(filepath.Separator)+"content"+string(filepath.Separator)+"post" {
// 			encontrou++
// 		}
// 		// fmt.Println(path)
// 		return nil
// 	})
//
// 	if encontrou != 2 {
// 		t.Error("Folders do hugo não foram criados")
// 	}
//
// }
//
// func setup() {
// 	err := os.RemoveAll(pathPosts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = os.Remove(yugoConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package yfl

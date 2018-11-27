package yfl

import (
	"os"
	"path/filepath"
)

// IniciaFoldersHugo os folders de post do hugo devem ser criados
func IniciaFoldersHugo() {
	path := "content" + string(filepath.Separator) + "post"
	os.MkdirAll(path, os.ModePerm)
}

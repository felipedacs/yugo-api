// yugo files
package yfl

import (
	"os"
	"path/filepath"
)

// os folders de post do hugo devem ser criados
func IniciaFoldersHugo() {
	var path string = "content" + string(filepath.Separator) + "post"
	os.MkdirAll(path, os.ModePerm)
}

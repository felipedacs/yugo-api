// yugo files
package yfl

import "os"

// os folders de post do hugo devem ser criados
func IniciaFoldersHugo(path string) {
	os.MkdirAll(path, os.ModePerm)
}

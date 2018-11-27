package ytr

import (
	"os/exec"

	"github.com/felipedacs/yugo-api/yutils"
)

// IniciaHugo inicializa server do Hugo
func IniciaHugo() {
	cmd := exec.Command("hugo", "server", "-D")
	_, err := cmd.CombinedOutput()
	yutils.Check(err)
}

// yugo terminal
package ytr

import (
	"fmt"
	"log"
	"os/exec"
)

func IniciaHugo() {
	cmd := exec.Command("hugo", "server", "-D")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

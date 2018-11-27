package ytr

import (
	"os/exec"

	"github.com/felipedacs/yugo-api/yutils"
)

// PublicaESalva hub de dois pushs
func PublicaESalva(user, senha, repo string) {
	cmd := exec.Command("hugo")
	_, err := cmd.CombinedOutput()
	yutils.Check(err)

	PushTo(".", user, senha, repo, "source")      //codigo fonte do blog
	PushTo("public", user, senha, repo, "master") //blog.github.io
}

// PushTo inicia git, adiciona, commita e pusha
func PushTo(dir, user, senha, url, branch string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	yutils.Check(err)

	cmd = exec.Command("git", "add", ".")
	_, err = cmd.CombinedOutput()
	yutils.Check(err)

	cmd = exec.Command("git", "commit", "-m", "foi")
	_, err = cmd.CombinedOutput()
	yutils.Check(err)

	remote := "https://" + user + ":" + senha + "@" + url
	branchremote := "master:" + branch
	cmd = exec.Command("git", "push", "-f", remote, branchremote)
	_, err = cmd.CombinedOutput()
	yutils.Check(err)
}

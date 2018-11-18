package ytr

import (
	"fmt"
	"log"
	"os/exec"
)

func PublicaESalva(user, senha, repo string) {
	fmt.Println("chegou aq")
	cmd := exec.Command("hugo")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	PushTo(".", user, senha, repo, "source")      //codigo fonte do blog
	PushTo("public", user, senha, repo, "master") //blog.github.io
}

func PushTo(dir, user, senha, url, branch string) {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	cmd = exec.Command("git", "add", ".")
	cmd.Dir = dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	cmd = exec.Command("git", "commit", "-m", "foi")
	cmd.Dir = dir
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	remote := "https://" + user + ":" + senha + "@" + url
	fmt.Println(remote)
	branchremote := "master:" + branch
	fmt.Println(branchremote)
	cmd = exec.Command("git", "push", "-f", remote, branchremote)
	cmd.Dir = dir
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

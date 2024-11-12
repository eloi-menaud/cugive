package main

import (
	"fmt"
	"os/exec"
	"bytes"
)

type Commit struct{
	short_id string
	message string
}

func GetCommits(target string, branch string) ([]Commit,error){

	// git fetch
	fmt.Printf("executing git fetch on %s\n", branch)

	cmd := exec.Command("git", "fetch","--depth=100000", "origin", branch)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetching commit history :\n-- stdout --\n%s\n-- stderr --\n%s\nerror : %s",
			stdout.String(), stderr.String(), err)
	}


	// get commit
	fmt.Printf("executing log %s\n", branch)

	cmd = exec.Command("git", "log", branch, "--", "target")

	stdout.Reset()
	stderr.Reset()

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get all commits :\n-- stdout --\n%s\n-- stderr --\n%s\nerror : %s",
			stdout.String(), stderr.String(), err)
	}

	fmt.Println(stdout.String())

	return nil,nil
}

func main(){
	_, err := GetCommits(".","main")
}
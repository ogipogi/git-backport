package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

const Version = "0.0.5"

type BackportOp struct {
	hash string
	branches []string
}

func main()  {
	args := os.Args
	var (
		cmdOut []byte
		err    error
		gitBranches []string
	)

	if len(args) < 2 {
		PrintManual()
		return
	}

	backportInfo := GetHashAndBranches(args[1])

	if len(backportInfo.branches) < 1 {
		fmt.Errorf("no git branches specified\n")
		PrintManual()
		return
	}

	cmdName := "git"
	cmdArgs := []string{"branch", "-a"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		PrintManual()
		os.Exit(1)
	}

	branches := strings.Split(strings.TrimSpace(string(cmdOut)), "\n")
	for _, element := range branches {
		branch := strings.Replace(element, "*", "", 1)
		branch = strings.Replace(branch, "remotes/origin/", "", 1)
		gitBranches = append(gitBranches, strings.TrimSpace(branch))
	}

	for _, branch := range backportInfo.branches {
		exists := BranchInBranchesSlice(strings.TrimSpace(branch), gitBranches)
		fmt.Println(branch, exists)
		if !exists {
			fmt.Fprintln(os.Stderr, "Error: could not find branch with name", branch)
			PrintManual()
			return
		}
	}
}

func PrintManual() {
	fmt.Printf("\ngit backport :: v%s\n\n", Version)

	fmt.Printf("HOW TO >>>>>\n")
	fmt.Println("$ git backport commit_hash:branch_name")
	fmt.Println("$ git backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Printf("<<<<<\n\n")
}

func GetHashAndBranches(input string) BackportOp {
	command := strings.Split(input, ":")
	hash := command[0]
	branches := strings.Split(command[1], ",")

	return BackportOp{hash, branches}
}

func BranchInBranchesSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

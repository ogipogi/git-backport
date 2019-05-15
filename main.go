package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

const Version = "0.0.8"

type BackportOperation struct {
	hash string
	branches []string
}

func main()  {
	args := os.Args

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

	branches := GetBranches()
	parsedBranches := ParseBranches(branches)
	CheckIfBranchesExist(backportInfo.branches, parsedBranches)
}

func PrintManual() {
	fmt.Printf("\ngit backport :: v%s\n\n", Version)

	fmt.Printf("HOW TO >>>>>\n")
	fmt.Println("$ git backport commit_hash:branch_name")
	fmt.Println("$ git backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Printf("<<<<<\n\n")
}

func GetBranches() []string {
	var (
		cmdOut []byte
		err error
	)

	cmdOp := "git"
	cmdArgs := []string{"branch", "-a"}

	if cmdOut, err = exec.Command(cmdOp, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		PrintManual()
		os.Exit(1)
	}

	return strings.Split(strings.TrimSpace(string(cmdOut)), "\n")
}

func ParseBranches(branches []string) []string {
	var gitBranches []string

	for _, element := range branches {
		branch := strings.Replace(element, "*", "", 1)
		branch = strings.Replace(branch, "remotes/origin/", "", 1)
		gitBranches = append(gitBranches, strings.TrimSpace(branch))
	}

	return gitBranches
}

func CheckIfBranchesExist(branches []string, gitBranches []string) {
	for _, branch := range branches {
		exists := BranchInBranchesSlice(strings.TrimSpace(branch), gitBranches)
		fmt.Println(branch, exists)
		if !exists {
			fmt.Fprintln(os.Stderr, "Error: could not find branch with name", branch)
			PrintManual()
			os.Exit(1)
		}
	}
}

func GetHashAndBranches(input string) BackportOperation {
	command := strings.Split(input, ":")
	hash := command[0]
	branches := strings.Split(command[1], ",")

	return BackportOperation{hash, branches}
}

func BranchInBranchesSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

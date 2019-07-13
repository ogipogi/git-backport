package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const Version = "0.0.14"

type BackportOperation struct {
	hash     string
	branches []string
}

func main() {
	args := os.Args

	if len(args) < 2 {
		PrintManual()
		os.Exit(0)
	}

	backportInfo := GetHashAndBranches(args[1])

	if len(backportInfo.branches) < 1 {
		fmt.Println(fmt.Errorf("no git branches specified\n"))
		PrintManual()
		return
	}

	branches := GetBranches()

	// check if the branch actually exists
	parsedBranches := ParseBranches(branches)
	CheckIfBranchesExist(backportInfo.branches, parsedBranches)

	// if all ok, then backport
	Backport(backportInfo.hash, backportInfo.branches)
}

func PrintManual() {
	fmt.Printf("\ngit Backport :: v%s\n\n", Version)

	fmt.Printf("HOW TO >>>>>\n")
	fmt.Println("$ git Backport commit_hash:branch_name")
	fmt.Println("$ git Backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Printf("<<<<<\n\n")
}

func GetHashAndBranches(input string) BackportOperation {
	command := strings.Split(input, ":")
	hash := command[0]
	branches := strings.Split(command[1], ",")

	return BackportOperation{hash, branches}
}

func GetBranches() []string {
	var (
		cmdOut []byte
		err    error
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

func Backport(commitHash string, branches []string) {

	cmdOp := "git"

	for _, branch := range branches {
		checkoutArgs := []string{"checkout", branch}

		if _, err := exec.Command(cmdOp, checkoutArgs...).Output(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cherryPickArgs := []string{"cherry-pick", commitHash}
		if _, err := exec.Command(cmdOp, cherryPickArgs...).Output(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
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

func BranchInBranchesSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func CheckIfBranchesExist(branches []string, gitBranches []string) {
	for _, branch := range branches {
		exists := BranchInBranchesSlice(strings.TrimSpace(branch), gitBranches)
		if !exists {
			fmt.Fprintln(os.Stderr, "Error: could not find branch with name", branch)
			PrintManual()
			os.Exit(1)
		}
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
	"os/exec"
)

const Version = "0.0.4"

type BackportOp struct {
	hash string
	branches []string
}

func main()  {
	args := os.Args
	var (
		cmdOut []byte
		err    error
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
		os.Exit(1)
	}

	fmt.Println(string(cmdOut))
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

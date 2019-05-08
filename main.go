package main

import (
	"fmt"
	"os"
	"strings"
)

const Version = "0.0.3"

type backportop struct {
	hash string
	branches []string
}

func main()  {
	args := os.Args

	if len(args) < 2 {
		PrintInfo()
	}

	GetHashAndBranches(args[1])
}

func PrintInfo() {
	fmt.Printf("\ngit backport :: v%s\n\n", Version)

	fmt.Printf("HOW TO >>>>>\n")
	fmt.Println("$ git backport commit_hash:branch_name")
	fmt.Println("$ git backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Printf("<<<<<\n\n")
}

func GetHashAndBranches(input string) backportop {
	command := strings.Split(input, ":")
	hash := command[0]
	branches := strings.Split(command[1], ",")

	return backportop{hash, branches}
}

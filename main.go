package main

import (
	"fmt"
	"os"
)

const Version = "0.0.2"

func PrintInfo() {
	fmt.Printf("\ngit backport :: v%s\n\n", Version)

	fmt.Printf("HOW TO >>>>>\n")
	fmt.Println("$ git backport commit_hash:branch_name")
	fmt.Println("$ git backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Printf("<<<<<\n\n")
}

func main()  {
	args := os.Args

	if len(args) < 2 {
		PrintInfo()
	}
}

package main

import "fmt"

const Version = "0.0.1"

func PrintInfo() {
	fmt.Printf("git backport :: v%s\n\n", Version)

	fmt.Println("HOW TO >>>>>")
	fmt.Println("$ git backport commit_hash:branch_name")
	fmt.Println("$ git backport commit_hash:branch_name1,branch_name2,branch_name3")
	fmt.Println("<<<<<")
}

func main()  {
	PrintInfo()
}

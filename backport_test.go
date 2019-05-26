package main

import "testing"

func TestHashParser(t *testing.T) {
	cmd := "4c3c5c7132eb5de5de1d552feb29dfca18a852d4:branch1"
	op := GetHashAndBranches(cmd)

	if op.hash != "4c3c5c7132eb5de5de1d552feb29dfca18a852d4" {
		t.Errorf("Commit hash was incorrect, got: %s", op.hash)
	}
}

func TestBranchParserWithOneBranch(t *testing.T) {
	cmd := "4c3c5c7132eb5de5de1d552feb29dfca18a852d4:branch1"
	op := GetHashAndBranches(cmd)

	if len(op.branches) != 1 {
		t.Errorf("Should have been 1 branch, got: %d", len(op.branches))
	}
}

func TestBranchParserWithThreeBranches(t *testing.T) {
	cmd := "4c3c5c7132eb5de5de1d552feb29dfca18a852d4:branch1,branch2,branch3"
	op := GetHashAndBranches(cmd)

	if len(op.branches) != 3 {
		t.Errorf("Should have been 3 branch, got: %d", len(op.branches))
	}
}

func TestBranchInListOfBranches(t *testing.T) {
	targetBranch := "branch2"
	branches := []string{"branch1", "branch2", "branch2"}
	op := BranchInBranchesSlice(targetBranch, branches)

	if op != true {
		t.Errorf("Branch should have been found in the list of branches")
	}
}

func TestBranchNotInListOfBranches(t *testing.T) {
	targetBranch := "branch4"
	branches := []string{"branch1", "branch2", "branch2"}
	op := BranchInBranchesSlice(targetBranch, branches)

	if op == true {
		t.Errorf("Branch should have not been found in the list of branches")
	}
}

func TestThatChecksIfGitBranchExists(t *testing.T) {
	targetBranches := []string{"branch1", "branch2"}
	gitBranches := []string{"branch1", "branch2", "branch2", "surpriseBranch"}
	CheckIfBranchesExist(targetBranches, gitBranches)
}

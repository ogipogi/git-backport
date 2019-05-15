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

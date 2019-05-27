#!/usr/bin/env bats

@test "when no args are supplied the usage manual is printed" {
  result="$(git backport)"
  [ "$result" -eq "

    git Backport :: v0.0.14

    HOW TO >>>>>
    $ git Backport commit_hash:branch_name
    $ git Backport commit_hash:branch_name1,branch_name2,branch_name3
    <<<<<

  "]
}
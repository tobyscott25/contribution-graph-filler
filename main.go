package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/tobyscott25/contribution-graph-filler/dates"
	"github.com/tobyscott25/contribution-graph-filler/files"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go </path/to/repo>")
		return
	}

	dummyCommitRepoPath := os.Args[1]
	// dummyCommitRepoPath := "../dummy-commit-repo"
	dummyCommitFileName := "dummy-commits.txt"
	dummyCommitFilePath := dummyCommitRepoPath + "/" + dummyCommitFileName

	joinDate, err := dates.AskUserForDate()
	if err != nil {
		fmt.Println("Error taking user input:", err)
		return
	}
	fmt.Println("Join date:", joinDate)

	daysBetween := dates.DaysBetween(joinDate, time.Now())
	fmt.Println("Will create " + strconv.Itoa(daysBetween) + " dummy commits.")

	repository, err := git.PlainOpen(dummyCommitRepoPath)
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	workTree, err := repository.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree from repository:", err)
		return
	}

	files.EditDummyCommitFile(dummyCommitFilePath, 1)

	_, err = workTree.Add(dummyCommitFileName)
	if err != nil {
		fmt.Println("Error staging file:", err)
		return
	}

	// We can verify the current status of the worktree using the method Status.
	status, err := workTree.Status()
	if err != nil {
		fmt.Println("Error getting the working tree status:", err)
		return
	}
	fmt.Println("Git status:")
	fmt.Println(status)

	globalGitConfig, _ := config.LoadConfig(1)

	commit, err := workTree.Commit("Synthetic Commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  globalGitConfig.User.Name,
			Email: globalGitConfig.User.Email,
			When:  time.Now().AddDate(0, 0, -1),
		},
	})
	if err != nil {
		fmt.Println("Error creating commit:", err)
		return
	}

	commitObject, err := repository.CommitObject(commit)
	if err != nil {
		fmt.Println("Error getting commit:", err)
		return
	}
	fmt.Println(commitObject)
}

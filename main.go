package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/tobyscott25/contribution-graph-filler/helper"
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

	startDate, err := helper.AskUserForDate()
	if err != nil {
		fmt.Println("Error taking user input:", err)
		return
	}

	endDate := time.Now()

	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {

		formattedDate := helper.HumanReadableFormat(date)
		numberOfCommits := helper.NumberOfCommits(date.Weekday() != time.Saturday && date.Weekday() != time.Sunday)

		for i := 0; i < numberOfCommits; i++ {

			helper.EditDummyCommitFile(dummyCommitFilePath)

			_, err = workTree.Add(dummyCommitFileName)
			if err != nil {
				fmt.Println("Error staging file:", err)
				return
			}

			_, err := workTree.Status()
			if err != nil {
				fmt.Println("Error getting the working tree status:", err)
				return
			}

			globalGitConfig, _ := config.LoadConfig(1)

			commit, err := workTree.Commit("Commit #"+strconv.Itoa(i+1)+" for "+formattedDate, &git.CommitOptions{
				Author: &object.Signature{
					Name:  globalGitConfig.User.Name,
					Email: globalGitConfig.User.Email,
					When:  date,
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
	}
}

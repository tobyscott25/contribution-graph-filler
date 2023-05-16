package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/tobyscott25/contribution-graph-filler/commits"
	"github.com/tobyscott25/contribution-graph-filler/dates"
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

	startDate, err := dates.AskUserForDate()
	if err != nil {
		fmt.Println("Error taking user input:", err)
		return
	}

	endDate := time.Now()

	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {

		formattedDate := date.Format("Mon 02 Jan 2006")
		fmt.Println(formattedDate)
		numberOfCommits := commits.NumberOfCommits(date.Weekday() != time.Saturday && date.Weekday() != time.Sunday)
		fmt.Println("NumberOfCommits:", numberOfCommits)

		commits.EditDummyCommitFile(dummyCommitFilePath)

		_, err = workTree.Add(dummyCommitFileName)
		if err != nil {
			fmt.Println("Error staging file:", err)
			return
		}

		status, err := workTree.Status()
		if err != nil {
			fmt.Println("Error getting the working tree status:", err)
			return
		}
		fmt.Println("Git status:")
		fmt.Println(status)

		globalGitConfig, _ := config.LoadConfig(1)

		for i := 0; i < numberOfCommits; i++ {

			commit, err := workTree.Commit("Synthetic commit for "+formattedDate, &git.CommitOptions{
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

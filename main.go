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

	fmt.Println("Enter a start date (DD-MM-YYYY):")
	var startInput string
	fmt.Scanln(&startInput)
	startDate, err := helper.ParseDateInput(startInput)
	if err != nil {
		fmt.Println("Invalid date format. Please try again.", err)
		return
	}

	fmt.Println("Enter an end date (DD-MM-YYYY): (Leave blank to use today as the end date)")
	var input string
	fmt.Scanln(&input)
	var endDate time.Time
	if input == "" {
		endDate = time.Now()
	} else {
		endDate, err = helper.ParseDateInput(input)
		if err != nil {
			fmt.Println("Invalid date format. Please try again.", err)
			return
		}
	}

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

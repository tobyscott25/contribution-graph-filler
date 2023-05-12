package main

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/tobyscott25/contribution-graph-filler/files"
)

func main() {

	// // Get the epoch timestamp from the command-line argument
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run main.go <epoch>")
	// 	return
	// }
	// epochStr := os.Args[1]
	// epoch, err := strconv.ParseInt(epochStr, 10, 64)
	// if err != nil {
	// 	fmt.Println("Error parsing epoch:", err)
	// 	return
	// }

	// // Convert the epoch to a time.Time value and format to be human-readable
	// t := time.Unix(epoch, 0)
	// formattedTime := t.UTC().Format("2 January 2006 15:04:05 MST")

	// fmt.Println("Commit created at: " + formattedTime)

	dummyCommitRepoPath := "../dummy-commit-repo"
	dummyCommitFileName := "dummy-commit-iterations.txt"
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

	commit, err := workTree.Commit("example go-git commit", &git.CommitOptions{
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

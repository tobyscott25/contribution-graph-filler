package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func addLineToEndOfFile(filename string, line string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintln(writer, line)
	if err != nil {
		return err
	}
	return writer.Flush()
}

func editDummyCommitFile(filePath string, dummyCommitIteration int) {

	if !fileExists(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	addLineToEndOfFile(filePath, "I must not tell lies.")
}

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

	editDummyCommitFile(dummyCommitFilePath, 1)

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

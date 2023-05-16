package helper

import (
	"bufio"
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func AddLineToEndOfFile(filename string, line string) error {
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

func EditDummyCommitFile(filePath string) {

	if !FileExists(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}

	AddLineToEndOfFile(filePath, "I must not tell lies.")
}

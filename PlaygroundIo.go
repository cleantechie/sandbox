package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the text")
	scanner.Scan()
	SaveData2(
		"/home/harsha/Devspace/GoSandBox/sandbox/test.txt",
		[]byte(scanner.Text()))
}

func SaveData1(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	fmt.Println("here")
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fp.Write(data); err != nil {
		return err
	}
	return fp.Sync() // fsync
}

// replacing the data atomically
// user either see the old file or new file
func SaveData2(path string, data []byte) error {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	tempFile := fmt.Sprintf("%s_temp_%d", path, random)
	fp, err := os.OpenFile(tempFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Printf("Closing the File %s", tempFile)
		// remove the temporary file
		fp.Close()
		if err != nil {
			os.Remove(tempFile)
		}
	}() // adding () after {} will invoke immediately

	if _, err = fp.Write(data); err != nil { // writes to in-memory copy
		return err
	}
	if err = fp.Sync(); err != nil { // changes flushed to disk
		return err
	}
	fmt.Println("Renaming the File")
	err = os.Rename(tempFile, path) // if everything goes fine rename to existing
	fmt.Println("Renaming Done")
	return err
}

package helper

import (
	"fmt"
	"io"
	"os"
)

func CopyGlobalFilesNetCore(localRepository string) {
	copyFile(getReadmeFile(), localRepository+"/README.md")
	copyFile(getGitIgnoreFile(), localRepository+"/.gitignore")
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func DeleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func createFolder(pathname string) {
	err := os.Mkdir(pathname, 0755)
	if err != nil {
		fmt.Println("Error when trying to get current directory", err)
		panic(err)
	}
}

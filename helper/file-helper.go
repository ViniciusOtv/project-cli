package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyGlobalFilesNetCore(localRepository string) {
	copyFile(getReadmeFile(), localRepository+"/README.md")
	copyFile(getGitIgnoreFile(), localRepository+"/.gitignore")
}

func copyListFile(src []string, applicationPath string) error {
	for _, value := range src {
		in, err := os.Open(value)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(applicationPath + "/Configurations/" + getRawFileName(value) + ".cs")
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}

		out.Close()
	}

	return nil
}

func getRawFileName(file string) string {
	fileName := filepath.Base(file)
	rawName := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	return rawName
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

func CreateFolder(pathname string) {
	err := os.Mkdir(pathname, 0755)
	if err != nil {
		fmt.Println("Error when trying to create current directory", err)
		panic(err)
	}
}

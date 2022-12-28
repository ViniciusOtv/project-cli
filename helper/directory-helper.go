package helper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func CurrentDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("error when trying to get current directory", err)
	}
	fmt.Println(path)
	return path
}

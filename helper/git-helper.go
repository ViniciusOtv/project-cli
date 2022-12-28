package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/fatih/color"
)

func CloneRepository(repository string) bool {
	result := true
	cyanColor := color.New(color.FgCyan).Add(color.Italic)

	fmt.Println("clonning the project...")
	fmt.Println("")
	cyanColor.Println("git clone " + repository)
	command := exec.Command("git", "clone", repository)
	err := command.Run()
	if err != nil {
		fmt.Printf("Error clonning project %v", err)
		result = false
	}

	return result
}

func CommitRepository(pathname string) {

	redColor := color.New(color.FgRed).Add(color.Bold)
	greenColor := color.New(color.FgGreen).Add(color.Bold)

	greenColor.Print(pathname)

	strCommand := "git add ."
	fmt.Println(strCommand)
	command := exec.Command("git", "add", ".")
	command.Dir = pathname

	err := command.Run()
	if err != nil {
		redColor.Println("Error when trying run git add command ", err)
	}

	strCommand = "git status"
	fmt.Println(strCommand)
	command = exec.Command("git", "status")
	command.Dir = pathname
	err = command.Run()
	if err != nil {
		redColor.Println("Error when trying run git status", err)
	}

	strCommand = "git commit -m project cli building initial structure"
	fmt.Println(strCommand)
	command = exec.Command("git", "commit", "-m", "'project cli building initial structure'")
	command.Dir = pathname

	println("committing the application " + command.Dir)

	err = command.Run()
	if err != nil {
		redColor.Println("Error when trying to run git commit", err)
	}

	strCommand = "git push"
	fmt.Println(strCommand)
	command = exec.Command("git", "push")
	command.Dir = pathname
	err = command.Run()
	if err != nil {
		redColor.Println("Error when trying to run git push ", err)
	}
}

func GetUsername() string {
	data, err := ioutil.ReadFile(RootDir() + "/git-user.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func RepositoryLocalPath(repositoryName string) string {
	return CurrentDirectory() + "/" + repositoryName
}

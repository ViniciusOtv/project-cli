package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
	hp "github.com/viniciusOtv/project-cli/helper"
)

type CreateRepositoryResponse struct {
	WebURL string `json:"html_url"`
	SSHURL string `json:"ssh_url"`
}

const httpConflict = 409

func CreateRepository(repositoryName string) CreateRepositoryResponse {

	url := "https://api.github.com/user/repos"

	payload := strings.NewReader("{\n	\"name\": \"" + repositoryName + "\"\n}")

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Basic "+GetGitHubBasicAuth())
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if res.StatusCode == httpConflict {
		fmt.Println("This repository already exist")
		os.Exit(0)
	}

	body, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	var bodyResponse CreateRepositoryResponse

	err = json.Unmarshal(body, &bodyResponse)

	cyanColor := color.New(color.FgCyan).Add(color.Italic)
	redColor := color.New(color.BgRed).Add(color.Italic)

	fmt.Println("")
	cyanColor.Printf("repository created successfully: ")
	redColor.Printf(bodyResponse.WebURL)
	fmt.Println("")
	cyanColor.Print(bodyResponse)

	return bodyResponse
}

func GetGitHubBasicAuth() string {
	return hp.BasicAuth(hp.GetUsername(), getPersonalAccessToken())
}

// func GetUserName() UserResponse {

// 	url := "https://api.github.com/user"

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", url, nil)

// 	if err != nil {
// 		panic(err)
// 	}

// 	req.Header.Add("Authorization", "Basic "+GetGitHubBasicAuth())
// 	req.Header.Add("Content-Type", "application/json")

// 	res, err := client.Do(req)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer res.Body.Close()

// 	if err != nil {
// 		panic(err)
// 	}

// 	body, err := ioutil.ReadAll(res.Body)

// 	defer res.Body.Close()

// 	var bodyResponse UserResponse

// 	err = json.Unmarshal(body, &bodyResponse)
// 	return bodyResponse
// }

func getPersonalAccessToken() string {
	data, err := ioutil.ReadFile(hp.RootDir() + "/personal-access-token.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

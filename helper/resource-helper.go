package helper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func getReadmeFile() string {
	return CurrentDirectory() + "/resources/global/README.md"
}

func getGitIgnoreFile() string {
	return CurrentDirectory() + "/resources/global/.gitignore"
}

func getGlobalConfigutationNetCoreFile() string {
	return RootDir() + "/resources/netcore/global.json"
}

func getNugetNetCoreFile() string {
	return RootDir() + "/resources/netcore/nuget.config"
}

func getNetCoreProgramFile() string {
	return CurrentDirectory() + "/resources/netcore/Program.txt"
}

func getNetCoreStartupFile() string {
	return CurrentDirectory() + "/resources/netcore/Startup.txt"
}

func getNetCoreDefaultControllerFile() string {
	return CurrentDirectory() + "/resources/netcore/BaseController.txt"
}

func getNetCoreSwaggerConfigurationFile() string {
	return CurrentDirectory() + "/resources/netcore/SwaggerConfiguration.txt"
}

func replaceWordInFile(filename string, old string, new string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(old), []byte(new), -1)

	if err = ioutil.WriteFile(filename, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

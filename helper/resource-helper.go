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

func getNetCoreServiceFile() string {
	return CurrentDirectory() + "/resources/netcore/Services/TelemetriaService.txt"
}

func getNetCoreInterfaceFile() string {
	return CurrentDirectory() + "/resources/netcore/Services/Interfaces/ITelemetriaService.txt"
}

func getNetCoreBaseResponseFile() string {
	return CurrentDirectory() + "/resources/netcore/Models/BaseResponse.txt"
}

func getNetCoreHealthCheckResponseFile() string {
	return CurrentDirectory() + "/resources/netcore/Models/HealthCheckResponse.txt"
}

func getNetCoreConfigurationsFile() []string {
	configurations := []string{CurrentDirectory() + "/resources/netcore/Configurations/SwaggerConfiguration.txt",
		CurrentDirectory() + "/resources/netcore/Configurations/DependencyInjectionConfiguration.txt",
		CurrentDirectory() + "/resources/netcore/Configurations/ErrorHandlingMiddlewareConfiguration.txt",
		CurrentDirectory() + "/resources/netcore/Configurations/HealthCheckConfiguration.txt",
		CurrentDirectory() + "/resources/netcore/Configurations/LogConfiguration.txt",
		CurrentDirectory() + "/resources/netcore/Configurations/TelemetriaConfiguration.txt"}

	return configurations
}

func replaceWordInConfigurationsFile(src []string, pathname string, old string, solutionName string) {
	for _, value := range src {
		filename := pathname + "/" + solutionName + "/Configurations/" + getRawFileName(value) + ".cs"

		input, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		output := bytes.Replace(input, []byte(old), []byte(solutionName), -1)

		if err = ioutil.WriteFile(filename, output, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
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

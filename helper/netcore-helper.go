package helper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func CreateNetCoreSolution(pathName string, solutionName string) {

	redColor := color.New(color.FgRed).Add(color.Bold)

	redColor.Println("Creatting NetCore Solution ")

	strCommand := "dotnet new sln -n " + solutionName
	fmt.Println(strCommand)
	command := exec.Command("dotnet", "new", "sln", "-n", solutionName)
	redColor.Println(command.Dir)
	command.Dir = pathName

	err := command.Run()
	if err != nil {
		redColor.Println("Error when trying to create the solution file", err)
		os.Exit(0)
	}
}

func CreateNetCoreAPI(pathname string, solutionName string, repository string) {
	fullName := solutionName
	createNetCoreConsoleApplication(pathname, fullName)
	addProjectToSlnProject(solutionName, pathname)
	setNetCoreCsprojProjectInAPIProject(solutionName, pathname)
	copyNetCoreAPIFileToAPI(solutionName, pathname, fullName, repository)
	createNetCoreTestProject(pathname, solutionName)
	createNetCoreApplicationFolderStructure(pathname, solutionName)
}

func copyNetCoreAPIFileToAPI(solutionName string, pathname string, name string, repository string) {
	applicationPath := pathname + "/" + solutionName
	copyFile(getNetCoreProgramFile(), applicationPath+"/Program.cs")
	copyFile(getNetCoreStartupFile(), applicationPath+"/Startup.cs")

	createFolder(applicationPath + "/Controllers")
	copyFile(getNetCoreDefaultControllerFile(), applicationPath+"/Controllers/BaseController.cs")

	createFolder(applicationPath + "/Configurations")
	copyFile(getNetCoreSwaggerConfigurationFile(), applicationPath+"/Configurations/SwaggerConfiguration.cs")

	programFile := pathname + "/" + solutionName + "/Program.cs"
	startupFile := pathname + "/" + solutionName + "/Startup.cs"
	baseController := pathname + "/" + solutionName + "/Controllers/BaseController.cs"
	swaggerConfiguration := pathname + "/" + solutionName + "/Configurations/SwaggerConfiguration.cs"
	csprojFile := pathname + "/" + solutionName + "/" + solutionName + ".csproj"

	replaceWordInFile(programFile, "[sln]", solutionName)
	replaceWordInFile(startupFile, "[sln]", solutionName)
	replaceWordInFile(baseController, "[sln]", solutionName)
	replaceWordInFile(swaggerConfiguration, "[sln]", solutionName)
	replaceWordInFile(csprojFile, "Microsoft.NET.Sdk", "Microsoft.NET.Sdk.Web")

	pathbase := "/" + repository + "/" + strings.ToLower(strings.Replace(name, ".", "/", -1))
	replaceWordInFile(startupFile, "[pathbase]", pathbase)

	DeleteFile(applicationPath + "/Class1.cs")

}

func createNetCoreTestProject(pathname string, solutionName string) {
	applicationPath := pathname

	name := solutionName + ".Test"

	redColor := color.New(color.FgRed).Add(color.Bold)

	strCommand := "dotnet new MSTest -n " + name
	fmt.Println(strCommand)
	command := exec.Command("dotnet", "new", "MSTest", "-n", name)
	command.Dir = applicationPath

	err := command.Run()
	if err != nil {
		redColor.Println("Erro when trying to create Test project", err)
		os.Exit(0)
	}
}

func createNetCoreConsoleApplication(pathname string, name string) {

	redColor := color.New(color.FgRed).Add(color.Bold)
	command := exec.Command("dotnet", "new", "classlib", "-f", "net6.0", "-n", name)
	command.Dir = pathname

	err := command.Run()
	if err != nil {
		redColor.Println("Error when trying to create th solution file ", err)
		os.Exit(0)
	}
}

func setNetCoreCsprojProjectInAPIProject(solutionName string, pathname string) {
	applicationCsprojPath := pathname + "/" + solutionName + "/" + solutionName + ".csproj"

	new := "<NoWarn>$(NoWarn);1591</NoWarn><GenerateDocumentationFile>true</GenerateDocumentationFile><OutputType>Exe</OutputType><ProjectGuid>{" + uuid.New().String() + "}</ProjectGuid><Version>1.0</Version></PropertyGroup>"
	old := "</PropertyGroup>"

	replaceWordInFile(applicationCsprojPath, old, new)
}

func addProjectToSlnProject(solutionName string, pathname string) {

	redColor := color.New(color.FgRed).Add(color.Bold)
	cyanColor := color.New(color.FgCyan).Add(color.Italic)

	solutionNameWithExtension := solutionName + ".sln"
	applicationCsprojPath := pathname + "/" + solutionName + "/" + solutionName + ".csproj"
	cyanColor.Println("dotnet sln " + solutionNameWithExtension + " add " + applicationCsprojPath)
	command := exec.Command("dotnet", "sln", solutionNameWithExtension, "add", applicationCsprojPath)
	command.Dir = pathname

	err := command.Run()
	if err != nil {
		redColor.Println("Error when trying to add the csproj application to the solution", err)
		os.Exit(0)
	}
}

func createNetCoreApplicationFolderStructure(pathname string, solutionName string) {

	applicationPath := pathname + "/" + solutionName
	repositoryPath := applicationPath + "/" + "Repositories"
	modelInternalPath := applicationPath + "/" + "Models"
	servicePath := applicationPath + "/" + "Services"

	createFolder(applicationPath + "/" + "Enums")
	createFolder(applicationPath + "/" + "Extensions")
	createFolder(applicationPath + "/" + "Mappers")
	createFolder(applicationPath + "/" + "Models")
	createFolder(modelInternalPath + "/" + "Dtos")
	createFolder(applicationPath + "/" + "Repositories")
	createFolder(applicationPath + "/" + "Entities")
	createFolder(applicationPath + "/" + "Requests")
	createFolder(applicationPath + "/" + "Responses")
	createFolder(repositoryPath + "/" + "Interfaces")
	createFolder(applicationPath + "/" + "Services")
	createFolder(servicePath + "/" + "Interfaces")
	createFolder(applicationPath + "/" + "Utils")
	createFolder(applicationPath + "/" + "Validations")

}

func CopyNetCoreFiles(localRepository string) {
	copyFile(getGlobalConfigutationNetCoreFile(), localRepository+"/global.json")
	copyFile(getNugetNetCoreFile(), localRepository+"/nuget.config")
}

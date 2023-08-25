package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	hp "github.com/viniciusOtv/project-cli/helper"
	sv "github.com/viniciusOtv/project-cli/service"
	sh "github.com/viniciusOtv/project-cli/shared/messages"
)

const maxLenghRepositoryName = 30

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create netcore applications",
	Long:  `This command will give more flexibility to create new git repository with netcore applications.`,
	Run: func(cmd *cobra.Command, args []string) {

		cyanColor := color.New(color.FgCyan).Add(color.Italic)

		redColor := color.New(color.BgRed).Add(color.Italic)

		sh.StartupMessage()

		sh.InitiationMessage()

		fmt.Println("")

		cyanColor.Printf("Enter the name of the repository to be created: ")
		var repositoryName string
		fmt.Scan(&repositoryName)

		if len(repositoryName) > maxLenghRepositoryName {
			fmt.Println("The name of repository can be a maximum of 30 characteres!!")
			os.Exit(0)
		}

		fmt.Println("")
		gitRepository := sv.CreateRepository(repositoryName)

		fmt.Println("")
		repository := gitRepository.WebURL

		redColor.Printf(repository)

		if hp.IsWindowsSystem() {
			repository = gitRepository.WebURL
		}

		if !hp.CloneRepository(repository) {
			redColor.Printf("error when cloning repository")
			os.Exit(0)
		}

		localRepository := hp.RepositoryLocalPath(repositoryName)
		hp.CopyGlobalFilesNetCore(localRepository)
		hp.CommitRepository(localRepository)

		fmt.Println("")

		fmt.Println("Creating NetCore project")

		addApi := hp.SubQuestion("Do you want to create an api")

		userSelectedYesToOneQuestion := (addApi)

		if userSelectedYesToOneQuestion {

			sh.DefinitionSolutionNameMessage("NetCore")

			cyanColor.Printf("Enter the name of the solution:  ")
			var solutionName string
			fmt.Scan(&solutionName)

			hp.CreateNetCoreSolution(localRepository, solutionName)
			hp.CopyNetCoreFiles(localRepository)

			hp.CreateNetCoreAPI(localRepository, solutionName, repositoryName)

			hp.CommitRepository(localRepository)

			fmt.Println("")
			fmt.Println("")
			fmt.Println("Process completed sucessfully")

		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

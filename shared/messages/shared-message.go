package messages

import (
	"fmt"

	"github.com/fatih/color"
)

func StartupMessage() {
	cyanColor := color.New(color.FgCyan).Add(color.Italic)

	fmt.Println("")
	cyanColor.Printf("It's great to see you here")
	// cyanColor.Println(hp.GetUsername())

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println("| Welcome to Project Creation Assistant                                                |")
	fmt.Println("----------------------------------------------------------------------------------------")
}

func InitiationMessage() {
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("| We will perform an important step and create the new repository                 |")
	fmt.Println("| * the repository name must not contain capital letters                          |")
	fmt.Println("| * repository name must not contain numbers                                      |")
	fmt.Println("| * repository name can be a maximum of 30 characters                             |")
	fmt.Println("| * repository name must not contain special characters except '-'				   |")
	fmt.Println("-----------------------------------------------------------------------------------")

}

func DefinitionSolutionNameMessage(tecnology string) {
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println("| Now we will define the name of the solution for the project(s) in "+tecnology, "|")
	fmt.Println("| * solution name must not contain numbers                                                  |")
	fmt.Println("| * solution name must not contain special characters                                       |")
	fmt.Println("| * the name of the solution must follow the Pascal Case pattern, example: GatewayPagamento |")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println("")
}

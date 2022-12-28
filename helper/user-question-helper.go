package helper

import (
	"fmt"

	"github.com/fatih/color"
)

func SubQuestion(text string) bool {
	cyanColor := color.New(color.FgCyan).Add(color.Italic)
	cyanColor.Println(text)
	cyanColor.Println("[1] SIM")
	cyanColor.Println("[2] NÃO")

	fmt.Printf("")
	var answer string
	fmt.Scan(&answer)
	fmt.Println("")
	return wasAfirmativeAnswer(answer)
}

func wasAfirmativeAnswer(answer string) bool {
	return answer == "1"
}

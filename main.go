package main

import (
	"fmt"
	"log"
	"os"
)

func newFile() *os.File {
	file, err := os.Create("./dumbAssLoop/dumbAssLoop.go")

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func startGoFile(file *os.File) {
	startGoFile := "package dumbassLoop\n\nimport (\n\t\"fmt\"\n)\n\n"
	file.WriteString(startGoFile)
}

func fileNewMainFunction(file *os.File, text string) {
	newFunc := fmt.Sprintln("func main() {\n", text, "\n}")
	file.WriteString(newFunc)
}

func dumbAssLoop(step int, stopAt int, newFunc string) string {
	newFunc += "\tfmt.Println(\"Hello World\")"
	if step != stopAt {
		newFunc += "\n"
		newFunc = dumbAssLoop(step+1, stopAt, newFunc)
	}
	return newFunc
}

func main() {
	file := newFile()
	startGoFile(file)
	howManyTimes := 10
	var newFunc string
	newFunc = dumbAssLoop(1, howManyTimes, newFunc)
	fileNewMainFunction(file, newFunc)
	file.Close()
}

package main

import (
	"fmt"
	"os"
)

type Context struct {
	ProjectName string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: create go project")
		os.Exit(0)
	}

	projectName := os.Args[1]
	err := createProject(projectName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("project created, enjoy!")
}

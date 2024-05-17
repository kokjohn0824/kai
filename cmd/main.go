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
		//TODO: update usage message
		fmt.Println("usage: create go project")
		os.Exit(0)
	}

	for _, projectName := range os.Args[1:] {
		err := createProject(projectName)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("project created, enjoy!")
}

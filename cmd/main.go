package main

import (
	"fmt"
	"os"
)

type Context struct {
	ProjectName string
}

const (
	verbose = `project created, you can now run the following commands:
go to your created directory 
cd  
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: kai <name>")
		os.Exit(0)
	}

	var name string
	for _, projectName := range os.Args[1:] {
		err := createProject(projectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		name = projectName
	}

	fmt.Print(verbose, name)
}

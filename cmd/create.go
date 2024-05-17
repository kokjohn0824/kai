package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/**
var embedDir embed.FS

func createProject(projectName string) error {
	rootdir := "templates"
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error while getting user directory: %v ", err)
	}
	ctx := Context{ProjectName: projectName}
	projectDir := path.Join(cwd, ctx.ProjectName)

	err = os.MkdirAll(projectDir, 0777)
	if err != nil {
		fmt.Println("error while creating project folder: ", err)
		return err
	}

	err = os.Chmod(projectDir, 0777)
	if err != nil {
		fmt.Println("error while chmod project folder: ", err)
		return err
	}

	walkerr := fs.WalkDir(embedDir, rootdir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("error while walking internal tree: ", err)
			return err
		}
		trimmed := strings.TrimPrefix(p, rootdir)

		if !d.IsDir() {
			f, err := embedDir.ReadFile(p)
			if err != nil {
				fmt.Println("error while opening tmpl: ", err)
				return err
			}

			tmpl, err := template.New(filepath.Base(p)).Parse(string(f))
			if err != nil {
				fmt.Println("error while parsing tmpl: ", err)
				return err
			}
			fp := path.Join(cwd, ctx.ProjectName, strings.TrimSuffix(trimmed, ".tmpl"))

			dir := filepath.Dir(fp)
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				fmt.Println("error while creating dir: ", err)
				return err
			}

			out, err := os.Create(fp)
			if err != nil {
				fmt.Println("error while creating outputfile:", err)
				return err
			}
			defer out.Close()

			err = tmpl.Execute(out, ctx)
			if err != nil {
				fmt.Println("error while Executing tmpl: ", err)
				return err
			}

			fmt.Println("file created: ", fp)
		}

		return nil
	})

	if walkerr != nil {
		fmt.Println("error while walking tree: ", err)
		return err
	}

	return nil
}

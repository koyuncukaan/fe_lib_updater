package main

import (
	"fe_lib_updater/meta"
	"fe_lib_updater/models"
	"fe_lib_updater/projs"
	"fmt"
)

func main() {
	meta.PrintLogo()

	var availableProjects []models.Project
	var workDir string

	for {
		workDir = projs.Entry()
		availableProjects = projs.Finder(workDir)
		if len(availableProjects) > 0 {
			break
		}
		fmt.Println("No projects found. Please enter a valid directory.")
	}

	projectsToBeUpdated := projs.Selector(availableProjects)
	fmt.Println(projectsToBeUpdated)
}

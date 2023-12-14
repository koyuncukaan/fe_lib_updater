package projs

import (
	"bufio"
	"fe_lib_updater/models"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Entry() string {
	fmt.Println("Enter work directory or press Enter for current dir: ")
	var workDir string
	if _, err := fmt.Scanln(&workDir); err != nil {
		fmt.Println(err)
		workDir = "."
	}
	fmt.Println("Work directory: ", workDir)
	return workDir
}

func Finder(workDir string) []models.Project {
	var projects []models.Project
	var currDir string
	projectPrefix := "fe_web_"
	err := os.Chdir(workDir)
	if err != nil {
		fmt.Println(err)
		Finder(workDir)
	}
	currDir, _ = os.Getwd()
	fmt.Println("Current directory: ", currDir)

	entries, err := os.ReadDir(currDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), projectPrefix) {
			project := models.Project{
				Name: entry.Name(),
				Path: filepath.Join(currDir, entry.Name()),
			}
			projects = append(projects, project)
		}
	}
	return projects
}
func Selector(availableProjects []models.Project) []models.Project {
	var selectedProjects []models.Project
	fmt.Println("Available projects:")
	for i, project := range availableProjects {
		fmt.Printf("%d: %s\n", i, project.Name)
	}

	fmt.Println("which projects should be updated? (separated by space)")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	projectsIds := scanner.Text()

	for _, projectId := range strings.Split(projectsIds, " ") {
		index, err := strconv.Atoi(projectId)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if index < len(availableProjects) {
			selectedProjects = append(selectedProjects, availableProjects[index])
		}
	}
	var confirm string
	fmt.Println("Selected projects:")
	for _, project := range selectedProjects {
		fmt.Printf("%s\n", project.Name)
	}
	fmt.Println("Confirm? (y/n)")
	if scanner.Scan() {
		confirm = scanner.Text()
		if confirm == "y" {
			return selectedProjects
		} else {
			return Selector(availableProjects)
		}
	}
	return selectedProjects
}

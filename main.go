package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func projectFinder() {
	var workDir, currDir string
	projectPrefix := "fe_web_"
	fmt.Println("Enter work directory: ")
	fmt.Scanln(&workDir)
	err := os.Chdir(workDir)
	if err != nil {
		fmt.Println(err)
		projectFinder()
	}
	currDir, _ = os.Getwd()
	fmt.Println("Current directory: ", currDir)

	entries, err := os.ReadDir(currDir)
	if err != nil {
		fmt.Println(err)
	}
	for i, entry := range entries {
		if entry.IsDir() && entry.Name()[:7] == projectPrefix {
			fmt.Println(i, entry.Name())
		}
	}
}
func projectSelector(projects *[]int) {
	fmt.Println("which projects should be updated? (separated by space)")
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	projectsIds := scanner.Text()

	for _, projectId := range strings.Split(projectsIds, " ") {
		project, err := strconv.Atoi(projectId)
		if err != nil {
			fmt.Println(err)
		}
		*projects = append(*projects, project)
	}
	var confirm string
	fmt.Printf("projects to be updated: %v, Confirm(y/n)", *projects)
	fmt.Scanln(&confirm)
	if confirm != "y" {
		*projects = nil
		projectSelector(projects)
	}
}

func main() {
	asciiArt := `
  _____           .__  ._____.                         .___       __                
_/ ____\____      |  | |__\_ |__       __ ________   __| _/____ _/  |_  ___________ 
\   __\/ __ \     |  | |  || __ \     |  |  \____ \ / __ |\__  \\   __\/ __ \_  __ \
 |  | \  ___/     |  |_|  || \_\ \    |  |  /  |_> > /_/ | / __ \|  | \  ___/|  | \/
 |__|  \___  >____|____/__||___  /____|____/|   __/\____ |(____  /__|  \___  >__|   
           \/_____/            \/_____/     |__|        \/     \/          \/       
`
	fmt.Println("\033[31m" + asciiArt + "\033[0m")
	var projects []int
	projectFinder()
	projectSelector(&projects)
}

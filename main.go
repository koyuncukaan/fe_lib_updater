package main

import (
	"fmt"
	"os"
)

func printWorkDir() {
	var workDir, currDir string
	projectPrefix := "fe_web_"
	fmt.Println("Enter work directory: ")
	fmt.Scanln(&workDir)
	err := os.Chdir(workDir)
	if err != nil {
		fmt.Println(err)
		printWorkDir()
	}
	currDir, _ = os.Getwd()
	fmt.Println("Current directory: ", currDir)

	entries, err := os.ReadDir(currDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name()[:7] == projectPrefix {
			fmt.Println(entry)
		}
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
	printWorkDir()
}

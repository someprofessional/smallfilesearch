package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SearchOneWordOccurence(word string) []string {

	allResults := []string{}

	searchedWord := strings.ToLower(word)

	outputDir := "../../output/"

	files, err := os.ReadDir(outputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		file, err := os.Open(outputDir + file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		lineBuffer := make([]string, 0, 4)

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Println("Error reading file:", err)
				break
			}
			lineBuffer = append(lineBuffer, line)
			if len(lineBuffer) > 4 {
				lineBuffer = lineBuffer[1:]
			}
			currFile := file.Name()

			var result strings.Builder
			searchValidatedSearchWord := " " + searchedWord + "\n"
			result.WriteString(fmt.Sprintf("\n\nFile name : %s\n", currFile))

			//Here the equality isn't full. if i search for example for test, it will return test even if the word is testing
			if strings.Contains(strings.ToLower(line), searchValidatedSearchWord) {

				result.WriteString(fmt.Sprintf("Searched word : %s\n", searchedWord))
				result.WriteString(fmt.Sprintf(lineBuffer[0]))
				result.WriteString(fmt.Sprintf(lineBuffer[1]))

				allResults = append(allResults, result.String())
			}
		}

	}

	return allResults
}

//func searchMultipleWordOccurence() {

//basically the same as before but with a big twist, i must get the sentence with the number of words before and after corresponding to the query
//so i also must find a way to correlate that to the original files

// open files
// look where the word is
// add result into a struct ?
// do it for each files
// do something to show from wich file it comes from
// send back the results structured
//}

//transform that into a worker pool
//https://gobyexample.com/worker-pools

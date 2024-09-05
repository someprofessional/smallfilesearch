package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SearchOneWordOccurence(word string) string {

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

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Println("Error reading file:", err)
				break
			}
			if strings.Contains(line, searchedWord) {
				return fmt.Sprintf("I found it! : %s", file.Name())
			}
		}

	}

	return "well well well, That's not supposed to happen ..."
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

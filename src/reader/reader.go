package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type FileLoadedMsg struct {
	Files []string
}

type WordOccurences struct {
	Word        string
	Occurence   int
	WordNumbers []int
}

func LoadFilesCmd() tea.Cmd {
	return func() tea.Msg {

		inputDir := "../../input"
		outputDir := "../../output"

		currFiles := make([]string, 0)

		files, err := os.ReadDir(inputDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			currFiles = append(currFiles, file.Name())

			file, err := os.Open(inputDir + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			wordMap := make(map[string]*WordOccurences)
			wordPosition := 1

			for scanner.Scan() {
				line := scanner.Text()
				words := splitIntoWords(line)

				for _, word := range words {
					word = strings.ToLower(word)
					if _, exists := wordMap[word]; !exists {
						wordMap[word] = &WordOccurences{
							Word:        word,
							Occurence:   0,
							WordNumbers: []int{},
						}
					}
					wordMap[word].Occurence++
					wordMap[word].WordNumbers = append(wordMap[word].WordNumbers, wordPosition)
					wordPosition++
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			var sb strings.Builder
			for _, wordInfo := range wordMap {
				sb.WriteString(fmt.Sprintf("Word: %s\n", wordInfo.Word))
				sb.WriteString(fmt.Sprintf("Occurence: %d\n", wordInfo.Occurence))
				sb.WriteString(fmt.Sprintf("Positions: %v\n", wordInfo.WordNumbers))
				sb.WriteString("\n")
			}
			curredTitle := strings.TrimPrefix(file.Name(), "../../input/")
			newTitle := "tokenized_" + curredTitle
			newFile, err := os.Create(outputDir + "/" + newTitle)
			if err != nil {
				fmt.Println(err)
			}
			defer newFile.Close()

			_, err = newFile.WriteString(sb.String())
			if err != nil {
				fmt.Println("Error writing to file:", err)
			}
		}

		return FileLoadedMsg{Files: currFiles}
	}
}

func splitIntoWords(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z]{4,}`)
	matches := re.FindAllString(text, -1)

	var cleanedWords []string
	for _, match := range matches {
		cleanedWord := cleanWord(match)
		if cleanedWord != "" {
			cleanedWords = append(cleanedWords, cleanedWord)
		}
	}
	return cleanedWords
}

func cleanWord(word string) string {

	re := regexp.MustCompile(`[ _";:]`)
	return re.ReplaceAllString(word, "")
}

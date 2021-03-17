package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	dictionaryFile := flag.String("dictionary",
		"txt",
		"please specify absolute path of dictionary file")

	inputFile := flag.String("input",
		"txt",
		"please specify absolute path of input file")

	flag.Parse()
	inputFileData, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("bad file")
		fmt.Println(err)
	}

	defer inputFileData.Close()
	inputText := ""
	scanner := bufio.NewScanner(inputFileData)
	for scanner.Scan() {
		inputText = scanner.Text()
	}

	dictionaryFileDate, err := os.Open(*dictionaryFile)
	if err != nil {
		fmt.Println("bad file")
		fmt.Println(err)
	}
	defer dictionaryFileDate.Close()
	dictionaryScanner := bufio.NewScanner(dictionaryFileDate)
	dictionaryMap := make(map[string]string)
	var text strings.Builder
	dictionaryList := make([]string, 0)
	for dictionaryScanner.Scan() {
		dictionaryText := dictionaryScanner.Text()
		dictionaryList = append(dictionaryList, dictionaryText)
		text.WriteString(dictionaryText)
		if len(dictionaryText) < 2 || len(dictionaryText) > 105 {
			fmt.Println("Each word in the dictionary is must be between 2 and 105 letters long")
			os.Exit(1)
		}
		if _, ok := dictionaryMap[dictionaryText]; ok {
			fmt.Println("No two words in the dictionary are the same.")
			os.Exit(1)
		} else {
			dictionaryMap[dictionaryText] = dictionaryText
		}
	}
	if len(text.String()) > 105 {
		fmt.Println("The sum of lengths of all words in the dictionary does not exceed 105")
		os.Exit(1)
	}
	dictionaryTxtMap := make(map[string]string)
	for _, dictionaryText := range dictionaryList {
		count := 0
		if len(dictionaryText) > 0 {
			Perm([]rune(dictionaryText), func(sText []rune) {
				if _, ok := dictionaryTxtMap[dictionaryText]; !ok {
					if strings.Contains(inputText, string(sText)) {
						count += 1
					}
					dictionaryTxtMap[dictionaryText] = dictionaryText
				}
			})
		}
		fmt.Printf("%v Number of occurence for %v ", count, dictionaryText)
		fmt.Printf("\n")
	}

}

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
	file, err := os.Open(*dictionaryFile)
	if err != nil {
		fmt.Println("bad file")
		fmt.Println(err)
	}

	defer file.Close()
	text := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()
	}

	infile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("bad file")
		fmt.Println(err)
	}
	defer infile.Close()
	inscanner := bufio.NewScanner(infile)
	for inscanner.Scan() {
		intext := inscanner.Text()
		count := 0
		if strings.Contains(text, intext){
		count += 1
		}
		if len(intext) > 0 {
			for i := 1; i < len(intext); i++ {
				stext := ScrambleStr(intext, int64(i))
				if strings.Contains(text, stext){
					count += 1
				}
			}
		}
		fmt.Printf("%v Number of occurence for %v ", count, intext)
		fmt.Printf("\n")
	}
}

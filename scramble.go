package main

import (
	"bufio"
	"io"
	"math/rand"
)

func ScrambleStr(in string, ti int64) string {
	rand.Seed(ti)
	inRune := []rune(in)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func Occurrences(word string, r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	wordCount := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if scanner.Text() == word {
			wordCount++
		}
	}
	return wordCount, scanner.Err()
}


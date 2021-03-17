package main

import (
	"math/rand"
)

func ScrambleStr(text string, number int64) string {
	rand.Seed(number)
	textRune := []rune(text)
	rand.Shuffle(len(textRune), func(i, j int) {
		textRune[i], textRune[j] = textRune[j], textRune[i]
	})
	return string(textRune)
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

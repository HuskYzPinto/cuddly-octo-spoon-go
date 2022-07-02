package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countVowels(word string) (count int) {
	vowels := "aeiou"
	for _, char := range word {
		if strings.Contains(vowels, string(char)) {
			count += 1
		}
	}
	return
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert a word: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(countVowels(strings.TrimSpace(input)))
}

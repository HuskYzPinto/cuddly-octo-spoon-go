package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pigLatin(word string) string {
	vowels := "aeiou"
	stringSlice := strings.Split(word, "")
	var consonantIndex int
	var text string
	for index, char := range stringSlice {
		if !strings.Contains(vowels, char) {
			consonantIndex = index
			break
		}
	}

	suffix := stringSlice[consonantIndex] + "ay"
	text = strings.Join(stringSlice[:consonantIndex], "") + strings.Join(stringSlice[consonantIndex+1:], "")

	return text + suffix
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input your word here: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(pigLatin(strings.TrimSpace(input)))
}

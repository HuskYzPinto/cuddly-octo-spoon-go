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
	for index, char := range stringSlice {
		if !strings.Contains(vowels, char) {
			suffix := stringSlice[index] + "ay"
			text := strings.Join(stringSlice[:index], "") + strings.Join(stringSlice[index+1:], "")
			return suffix + text
		}
	}
	return ""
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

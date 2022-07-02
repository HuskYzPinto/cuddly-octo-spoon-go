package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) int {
	wordList := strings.Split(text, " ")
	return len(wordList)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input a text: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(countWords(strings.TrimSpace(input)))
}

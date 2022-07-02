package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(text string) (result string) {
	for _, char := range text {
		result = string(char) + result
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input your text here: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.TrimSpace(input)
	fmt.Println(reverseString(trimmedInput))

}

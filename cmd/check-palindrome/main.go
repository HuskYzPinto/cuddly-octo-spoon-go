package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkPalindrome(word string) bool {
	reversedString := ""
	for _, value := range word {
		reversedString = string(value) + reversedString
	}
	if reversedString == word {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input your text: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t\n", checkPalindrome(strings.TrimSpace(input)))
}

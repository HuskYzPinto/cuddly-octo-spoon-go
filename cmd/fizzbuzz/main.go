package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func fb(number int) {
	fmt.Println(number)
	for i := 1; i <= number; i++ {
		var result string
		if i%3 == 0 {
			result += "fizz"
		}
		if i%5 == 0 {
			result += "buzz"
		}
		if result != "" {
			fmt.Println(result)
		} else {
			fmt.Println(i)
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please insert your number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	trimmedInput := strings.TrimSpace(input)
	n, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fb(n)
}

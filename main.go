package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}
func GetUserInput() int {
	var input int
	fmt.Print("Enter a number between 1 and 10: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You guessed", input)
	}
	return input
}
func IsMatch(RandomN int, input int) bool {
	if RandomN > input {
		fmt.Println("Your guess is too small")
		return false
	} else if RandomN < input {
		fmt.Println("Your guess is too small")
		return false
	} else {
		fmt.Println("You guessed it!")
		return true
	}

}
func main() {

	RandomN := GetRandomNumber()
	count := 0
	for count < 3 {
		input := GetUserInput()
		fmt.Println(IsMatch(RandomN, input))
		count++
	}

}

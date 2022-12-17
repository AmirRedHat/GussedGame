package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_numbers(length int) []int {
	arr := make([]int, length)
	var i int = 0
	for i < length {
		arr[i] = i
		i++
	}
	return arr
}

func main() {

	// generating an array of numbers
	arr1 := generate_numbers(100)

	// choosing a random number as answer
	var pickedup_number int
	rand.Seed(time.Now().Unix())
	pickedup_number = rand.Intn(len(arr1))
	fmt.Println(pickedup_number)

	// working with pointers
	var user_number int

	fmt.Print("Enter a number between 0 and 100: ")
	fmt.Scanln(&user_number)

	for pickedup_number != user_number {

		fmt.Println("Incorrect!")
		fmt.Print("Enter a number between 0 and 100: ")
		fmt.Scanln(&user_number)
	}

	fmt.Println("you got it!")


}

package main

// Understnding Golang

/* Hello World**/
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Variables
	var message string = "Hello World!"
	fmt.Println(message)

	// Short declaration method
	age := 24
	fmt.Println(age)

	// Contants
	const PI = 3.142
	fmt.Println("PI:", PI)

	/* Data Types*/
	var integer int = 42
	var floatingPoint float64 = 0.343
	var isTrue bool = true
	var text string = "Hello World!"

	fmt.Println(integer, floatingPoint, isTrue, text)

	/* Control Flow  */

	// If statement

	age = 18
	if age >= 18 {
		fmt.Println("You're an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	// for loop

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Friday":
		fmt.Println("It's Friday")

	default:
		fmt.Println("It's Weekend")
	}

	// Arrays
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println("Array", numbers)

	// Slices
	slice := numbers[1:]
	fmt.Println("Slice", slice)

	// Maps
	person := map[string]interface{}{
		"name":      "john",
		"age":       30,
		"isStudent": false,
	}

	fmt.Println("Map:", person)
	// NOTE: Array's have a fixed value while slices are dynamic
	// Maps are key-value pairs

	// Instance of struct
	person2 := Person{
		Name:    "Alice",
		Age:     23,
		Address: "123, Main St",
	}

	fmt.Println("Name", person2.Name)
	fmt.Println("Age", person2.Age)
	fmt.Println("Address", person2.Address)

	// Variables and Pointers
	value := 42
	pointer := &value

	fmt.Println("Value:", value)
	fmt.Println("Pointer:", pointer)
	fmt.Println("Dereferenced:", *pointer) // NOTE: Pointer stores the memory address of a variable

	// Function Results
	sum := add(3, 2)
	fmt.Println("Sum", sum)

	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("Quotient", quotient, "Remainder", remainder)

	// Concurrencies and Goroutines
	var wg sync.WaitGroup

	// Launch a goroutine
	wg.Add(1)
	go printNumbers(&wg)

	// Run a function in the main goroutine
	printLetters()

	// Wait for the goroutines to finish
	wg.Wait()
}

// Struct
type Person struct {
	Name    string
	Age     int
	Address string // Struct allows you define your own data types
}

// Functions with parameter and return values
func add(a, b int) int {
	return a + b
}

// Function with multiple returns
func divideAndRemainder(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor

	return quotient, remainder
}

// Concurrencies and Goroutines

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)

		fmt.Println(i, " ")
	}
}

func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(string(char), " ")
	}
}

// Goroutines enables concurrent executions

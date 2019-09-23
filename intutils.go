package main

// CHANGED HAHAHAHA
import (
    "fmt"
)

// This function returns the smaller number of the two arguments...
func IntMin(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

// some comment in not changed line...
func main() {
	fmt.Println("This program does nothing. Was written to test testing and Github actions...")
	fmt.Println("It can help you decide between two numbers and return the smaller one...")
	fmt.Printf("Between 2 and 3 the smaller is: %d\n", IntMin(2,3))
	fmt.Printf("Between 2 and -3 the smaller is: %d\n", IntMin(2,-3))
}

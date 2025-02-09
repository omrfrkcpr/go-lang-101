package main

import "fmt"

// A loop statement allows us to execute code multiple times, in a loop.

// A loop can be used for many different use cases, one is to iterate over elements in a list.

var taskItems = []string{"first task", "second task", "third task"}

// For each task in the list, we want to execute the same logic. 

// Loop Types in Go:

// 1. Basic For loop: A loop that executes a block of code a specific number of times.
// 2. For loop with condition: A loop that continues executing as long as a specified condition is true.
// 4. For range loop: A loop that iterates over elements in slices, arrays or maps.

func main(){
	// Basic For loop

	// i : index of first task
	// i < range : condition to check if the index is less than the length of the range

	for index, task := range taskItems{
		fmt.Println(index+1, task)
	}

		// Blank Identifier (_): Go requires you to use all variables you declare. To ignore a variable we dont want to use. Unlike other languages, in Go we need to make unused variables explicit.

	for _, task := range taskItems{
		fmt.Println(task)
	}


}
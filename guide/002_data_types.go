package main

import "fmt"

// Data types specify the type of data that can be stored and manipulated by a program.

// Different programming languages have variations in their handling of data types and some offer specialized types beyond the common ones.

// Common data types: string, boolean, map, array, slices, numbers (int, float, complex), etc.

// Integers = whole numbers - positive and negative numbers
// Floating = decimal numbers

func dataTypes(){
	fmt.Println("string", "string")
	fmt.Println("boolean", true)
	fmt.Println("int", 10)
	fmt.Println()
	fmt.Println("float", 10.5)
	fmt.Println("complex", 10 + 5i)
}

func main() {
	// Calling dataTypes function to display the data types
	dataTypes()

	var mainData = "This is a main data"
	sameMainData := true
	fmt.Println(mainData)
	fmt.Println(sameMainData)

	// Arrays and slices in Go can hold only elements of the same type.
	var mainSlice = []string {"item1", "item2", "item3", mainData}
	fmt.Println("Main Slice:",mainSlice) 

	// Arrays have a fixed size. Slice is an abstraction of an Array. 
	var mainArray = [3]int{1, 2, 3}
	fmt.Println("Main Array:",mainArray)

	// Maps in Go are unordered collections of key-value pairs.
}
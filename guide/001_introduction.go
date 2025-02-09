package main

// We need to start with package name. All Go files start with package. Defines the namespaces for the code in the file, used to organize and group code.

// Go's standard library comes with core packages for us to use.

// fmt = Formatting for I/O

// os = Operating system interaction

// time = Time and date utilities

// net = Networking

import "fmt"

func main() {
	fmt.Print("Hello, World!\n") // Add a new line after the string
	fmt.Println("Hello, World!") // Add automatically new lines after the string
}

// To run the Go file: $ go run index.go


 


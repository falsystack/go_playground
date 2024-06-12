package main

import (
	"hello/dependency_injections"
	"os"
)

func main() {
	dependency_injections.Greet(os.Stdout, "Elodie")
}

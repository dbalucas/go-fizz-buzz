package main

import "fmt"

func main() {

	var limit = 1
	for limit <= 80 {
		if limit%3 != 0 && limit%5 != 0 && limit%15 != 0 {
			fmt.Print(limit)
		}
		if limit%3 == 0 {
			fmt.Print("fizz")
		}
		if limit%5 == 0 {
			fmt.Print("buzz")
		}
		if limit != 80 {
			fmt.Print(", ")
		}
		limit++
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Test Begin:")
	var str string
	for {
		fmt.Scan(&str)
		switch str {
		case "quit":./
			fmt.Println("over,quit")
			return
		case "version":
			fmt.Println("version 1.9")
		}

	}

}

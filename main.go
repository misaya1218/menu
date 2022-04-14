package main

import "fmt"

func main() {
	fmt.Println("Tes1t Begin:")
	var str string
	for {
		fmt.Scan(&str)
		switch str {
		case "quit":./
			fmt.Println("test over1 !,quit")
			return
		case "version":
			fmt.Println("version 1.9")
		}

	}

}

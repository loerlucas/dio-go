package main

import "fmt"

func main() {
	//tabuada do 3
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
	//Bricadeira do Pin/Pan
	for i := 1; i < 101; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PIN/PAN")
		} else if i%3 == 0 {
			fmt.Println("PIN")
		} else if i%5 == 0 {
			fmt.Println("PAN")
		} else {
			fmt.Println(i)
		}
	}
}

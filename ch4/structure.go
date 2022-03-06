package main

import "fmt"

func main()  {
	tryLoop()
	forLabel()
}

func tryLoop() {

	words := []string{"a", "bob", "apple", "orange", "this is a word"}
	for idx, word := range words {
		switch size := len(word); size {
		case 1, 2, 3:
			fmt.Println("this is a small word", idx)
		case 4, 5, 6:
			fmt.Println("this is a good word", idx)
		default:
			fmt.Println("it is too big word")
		}
	}
}

func forLabel(){

	loop:
	for i := 0; i < 10; i++ {
		switch {
		case i % 2 == 0 :
			fmt.Println(i, "is even")
		case i % 3 == 0 :
			fmt.Println(i, " is divded by 3, not 2")
		case i % 7 == 0 :
			fmt.Print("exit the loop")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}
}

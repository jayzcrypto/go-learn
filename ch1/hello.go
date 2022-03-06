package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const x int = 10
	fmt.Print("the value of x = " + string(x))
	fmt.Print("greetings \n \" Go Lang\"\n" + string(x))
	fmt.Print(`Good morning! 
hopefully can used to build "Cronos"
`)

	var b [3]int
	var a = [3]int{1, 2, 3}

	fmt.Print(a)
	b[0] = 2
	b[1] = 4
	b[2] = 6

	fmt.Print(b)

	var m = []int{4, 5, 6}
	var n = []int{8, 9, 10}
	m = append(m, n...)
	fmt.Println(m)
	lenCap()
	tryMap()
	trySet()
	tryStruct()
	tryIfBlock()

}

func lenCap() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

}

func tryMap() {

	var totalWins = map[string]int{}
	totalWins["hodl"] = 1
	fmt.Println(totalWins["hodl"])
	totalWins["moons"]++
	fmt.Println(totalWins["moons"])
	var v, ok = totalWins["tendis"]
	fmt.Println(v, ok)

	//delete from a map
	delete(totalWins, "hodl")
	fmt.Println(totalWins["hodl"])
}

func trySet()  {
	var intSet = map[int]bool{}
	var vals = []int{1, 2, 4, 5, 7, 5, 10, 15}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet[5], len(intSet))
}

func tryStruct() {
	type person struct {
		name string
		age int
		pet string
	}
	bob := person{
		"bob",
		30,
		"alice",
	}
	fmt.Println(bob.name)

	var manager struct {
		name string
		age int
	}
	manager.name = "Danny"
	manager.age = 40

	fmt.Println(manager.name)

	 pet := struct {
		name string
		kind string
	}{
		"wang cai",
		"dog",
	}
	fmt.Println(pet.name)

	 type firstPerson struct {
		 name string
		 age int
	 }

	 var g  struct{
		 name string
		 age int
	 }

	 f := firstPerson{
		 name : "jay z",
		 age : 40,
	 }

	g = f

	fmt.Println(g == f)
}

func tryIfBlock()  {

	n := rand.Intn(10)
	fmt.Println(n)
	if n == 0 {
		fmt.Println("number is 0")
	} else if n > 5 {
		fmt.Println("number is too big")
	} else {
		fmt.Println("it's a good number")
	}

}

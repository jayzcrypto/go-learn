package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	fucAsVal()
	innerFuc()
	v := baseMulti(2)(3)
	fmt.Println(v)
	v = baseMulti(4)(5)
	fmt.Println(v)
	//closureCall()
	callByValue()
}

func closureCall () {
	f, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer closer()
	data := make([]byte, 2048)

	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}


func callByValue() {
	type Person struct {
		name string
		age int
	}

	a , b, c := 2, "hello", Person{"alice", 20}

	func (i int, s string, p Person){
		i = i * 2
		s = "goodbye"
		p.age = 23
		p.name = "bob"
	}(a, b, c)
	fmt.Println(a, b, c)

	amap := map[int]string {
		1 : "this",
		2: "that",
	}
	func (amap map[int]string){
		amap[2] = "hello"
		amap[3] = "goodbye"
		delete(amap, 1)
	}(amap)
	fmt.Println(amap)
}

func getFile(name string)( *os.File, func(), error) {

	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}

	return file, func() {
		file.Close()
	}, nil
}

func baseMulti(i int) func(int) int {
	return func(j int) int {
		return i * j
	}
}


type opFuncType func(int, int) int
func innerFuc()  {

	for i := 0 ; i < 5; i++ {

		func(j int) {
			fmt.Println(" Looping through, current value is ", j)
		}(i)
	}
}


func fucAsVal(){
	expressions := [][]string{
		{"2", "+", "3"},
		{"5", "-", "2"},
		{"6", "/", "3"},
		{"2", "x", "4"},
		{"x"},
		{"x", "y"},
	}

	for idx, expr := range  expressions {
		if len(expr) != 3 {
			fmt.Println(idx, "invalid expression", expr)
			continue
		}

		p1, err := strconv.Atoi(expr[0])

		if err != nil {
			fmt.Println(idx, "invalid expr", err)
			continue
		}

		p2, err := strconv.Atoi(expr[2])

		if err != nil {
			fmt.Println(idx, "invalid expr", err)
			continue
		}

		op := optMap[expr[1]]

		fmt.Println("execution value = " , op(p1, p2))
	}
}

var optMap = map[string]opFuncType {
	"+" : add,
	"-" : sub,
	"x" : mul,
	"/" : div,
}
func add(a int, b int)  int {
	return a + b
}

func sub(a int, b int) int  {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}

func div(a int, b int)  int{
	return a / b
}



package main

import (
	"fmt"
	"time"
)

func main(){
	p := Person{
		"Jay",
		"Zeng",
		39,
	}
	fmt.Println(p.String())
	var c Counter
	fmt.Println(c)
	c.Increment()
	fmt.Println(c)
	doUpdateWrong(c)
	fmt.Println("in main after doUpdateWrong", c.String())
	doUpdateRight(&c)
	fmt.Println("in main after doUpdate Right", c.String())
	var i int = 100
	var s Score = 300
	var hs HighScore = 500

	s = Score(i)
	hs = HighScore(s)

	var empoyee = Employee(p)
	employee := Employee{
		FirstName: "Bob",
		LastName: "Lee",
	}

	fmt.Println("hs =" , hs, "employee = ", empoyee, employee )

	//type inner outer
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello"}
	fmt.Println(o.Double())

}

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

type Person struct {
	FirstName string
	LastName string
	Age int
}

type Employee Person
type Score int
type HighScore Score




func (p Person) String() string {
	return fmt.Sprintf("%s %s Age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c *Counter) Increment(){
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total %d , lastUpdate %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in Do Update Wrong", c.String())
}

func doUpdateRight(c *Counter)  {
	c.Increment()
	fmt.Println("In update right", c.String())
}



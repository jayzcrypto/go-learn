package main

import "fmt"

func main()  {
	theBasics()
	derefStruct()
	derefPointer()
}

func theBasics() {
	var x = 10
	var y = true
	var pointerX = &x
	var pointerY = &y
	var pointerZ *string

	fmt.Println(x, pointerX, &x)
	fmt.Println(y, pointerY, &y)
	fmt.Println(pointerZ)
//	fmt.Println(*pointerZ)
}

func derefStruct() {
	type Person struct {
		firstname string
		middleName *string
		lastname string
	}

	helper := func(str string) *string {
		return &str
	}

	p := Person{
		"Rachel",
		helper("Karren"),
		"Green",
	}

	fmt.Println(p)
	fmt.Println(*p.middleName)

	p.middleName = helper("bob")

	fmt.Println(p)
	fmt.Println(*p.middleName)


	type Foo struct {
		Field1  string
		Field2 string
	}

	makeFoo := func (f *Foo) error {
		f.Field1 = "foo"
		f.Field2 = "bar"
		return nil
	}

	f := Foo{}
	err := makeFoo(&f)
	if err != nil {
		return
	}
	fmt.Println(f)

	makeFoo2 := func() (Foo, error) {
		f := Foo{
			Field1: "foo1",
			Field2: "bar1",
		}
		return f, nil
	}
	fmt.Println(makeFoo2())
}

func derefPointer() {

	//copy in by value, so px in the parameter is not the same as the one in the outer fuc
	updatedFailed := func(px *int) {
		x := 20
		px = &x
	}

	//tho px in the parameter is a copy of the outer &x, the update fuc
	// is doing is referring to the value of the &x and update the value directly
	updated := func(px *int) {
		*px = 20
	}

	x := 10
	updatedFailed(&x)
	fmt.Println(x)
	updated(&x)
	fmt.Println(x)
}





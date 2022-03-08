package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main()  {

	c := Client{
		L: LogicProvider{},
		B: LogicProvider{},
	}
	c.Program()
	data, err :=processJson()
	if err == nil {
		for k, v := range  data {
			fmt.Println("Key =", k, "| value = ", v)
		}
	}

	var i interface{}
	var mine MyInt = 20
	i = mine
	//i2 := i.(MyInt)
	err2 := printError(i)
	fmt.Println(err2)
}

func printError(val interface{}) error {
	v, ok := val.(int)
	if !ok {
		return  fmt.Errorf("unexpected type for %v",v)
	}
	return nil
}

// MyInt type assertion & switch
type MyInt int

func processJson() (map[string]interface{}, error){

	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("test.json")
	if err != nil {
		return nil, err
	}
 //  fmt.Println(contents)
	err1 := json.Unmarshal(contents, &data)
	return data, err1

}

type LogicProvider struct {}

func (lp LogicProvider) Process(data string) string {
	fmt.Println(data)
	return data
}

func(lp LogicProvider) Think(data string) string {
	fmt.Println(data)
	return data
}

type Logic interface {
	Process(data string) string
}

type Behavior interface {
	Think(data string) string
}

type Client struct {
	L Logic
	B Behavior
}

func (c Client) Program(){
	data := "hello world"
	c.B.Think(data)
	c.L.Process(data)

}

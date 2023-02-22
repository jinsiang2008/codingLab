// Follow the doc: https://learnku.com/go/t/24715
// Feb 22, 2023 Shanghai

package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	gender string
}

func (p *person) describe() {
	fmt.Println("%v is %v years old.", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p *person) setName(name string) {
	p.name = name
}

func main() {
	// var ap *int
	// a := 12
	// ap = &a

	// fmt.Println(a)
	// fmt.Println(*ap)

	//指定属性和值
	p1 := person{name: "Bob", age: 42, gender: "Male"}
	//指定值
	p2 := person{"Alice", 30, "Female"}

	pp1 := &p1
	pp2 := &p2
	
	fmt.Println(pp1.name)
	fmt.Println(pp2.age)

	pp1.describe()

	pp1.setAge(10)
	fmt.Println(pp1.age)

	pp2.setName("Alice2")
	fmt.Println(pp2.name)
}
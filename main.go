package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type SS struct {
	Age int
}

func (s SS) Get() int {
	return s.Age
}

func (s SS) Set(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func test(i int) {
	fmt.Printf("%p\n", &i)
}

func main() {
	// ss := &SS{}

	var a = &SS{}
	fmt.Println(a)

	// f(ss) //value
	i := 10
	fmt.Printf("%p\n", &i)
	test(i)
}

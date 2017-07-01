package main

import "fmt"

//go 语言中 你可以给任意类型添加方法
//除了指针类型
type Integer int

//面向对象的方式
func (a Integer) Less(b Integer) bool {
	return a < b
}

//面向过程的方式
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

//这里修改了a的值，所以必须要用指针，因为go也是基于值传递的
func (a *Integer) Add(b Integer) {
	*a += b
}
func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
	a.Add(2)
	fmt.Println("a = ", a)
	if Integer_Less(a, 2) {
		fmt.Println(a, "Less 2")
	}
}

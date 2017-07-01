package main

import "fmt"

var v1 int

var v2 string
var v3 [10]int  //数组
var v4 []int    //数组切片
var v5 struct { //结构体
	f int
}
var v6 *int //指针

var v7 map[string]int //map，key为string类型，value为int类型
//声明多个变量
var (
	v8 int
	v9 string
)

var v10 func(a int) int // a为参数 int为返回类型

//声明初始化的变量
var v11 int = 10
var v12 = 10 //编译器自动推导类型
//v13 := 10 := 在这里定义是不行的 不能定义在函数体里面 函数里面定义的变量是必须使用的
func main() {
	v13 := 10 //编译器自动推导类型,而且变量必须是新的变量
	v1 = 123
	v8 = 321
	v1, v8 = v8, v1 //变量交换，而不需要中间变量
	fmt.Println(v13)
	fmt.Println(v1)
}

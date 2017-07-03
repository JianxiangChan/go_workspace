package main

import "fmt"

import "addless"

import "two"
import "one"
import "writer"

func main() {

	var a addless.Integer = 1
	var b addless.LessAdder = &a
	//3.5.1-3.5.3接口基本使用方法
	fmt.Println(a, b)
	a.Add(2)
	fmt.Println(a, b)
	b.Add(2)
	fmt.Println(a, b)

	//这两个接口并无不同
	var file1 one.ReadWrite = new(one.File)
	var file2 two.IStream = file1
	var file3 one.ReadWrite = file2

	var file4 one.ReadWrite = new(one.File)
	var file5 writer.Writer = file4

	//这里会报错 因为writer里面的方法只是one的子集
	var file6 writer.Writer = new(one.File)
	//	var file7 one.ReadWrite = file6
	fmt.Println(file1, file2, file3, file5, file6)

	//3.5.4接口查询
	if file7, ok := file6.(writer.Writer); ok {
		fmt.Println("ok")
		fmt.Println(file7)
	}

	//查询类型
	if file7, ok := file6.(*one.File); ok {
		fmt.Println("ok")
		fmt.Println(file7)
	}

	//3.5.5类型查询
	//var v1 interface{} = "abc"
	switch v := b.(type) {

	case addless.LessAdder:
		fmt.Println("string", v)
	default:
		fmt.Println("string2", v)
	}
	//现实使用的时候就是将接口查询和类型查询结合起来使用，使用先用类型查询
	//然后用接口查询查看是否实现了对应方法，

	//3.5.6 接口组合 见one.go

	//any类型的
	//interface可以指向任意类型 Println就是通过这个来实现的

}

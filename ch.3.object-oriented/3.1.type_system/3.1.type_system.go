package main

import "fmt"

type Integer int

func (a Integer) less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func (a *Integer) Modify(b Integer) {
	*a += b
}

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func main() {

	//3.1.1 为类型添加方法
	var a Integer = 1
	if a.less(2) {
		fmt.Println(a, "Less 2")
	}
	//go语言也是基于值传递的
	a.Add(2)
	fmt.Println("a =", a)

	//参考/src/pkg/http/header.go里面部分代码
	//header类型就是map类型

	//3.1.2 值语义和应用语义
	b := a
	b.Modify(2)
	fmt.Println("a =", a)
	fmt.Println("b =", b) //对b的修改不影响a
	//go的四种引用语义 都是靠指针来实现的 具体见笔记

	//3.1.3 结构体
	//go语言的结构体和C语言并没有什么不同 放弃了继承在内的大量的面向对象的特性
	var c Rect
	c = Rect{x: 12, y: 12, width: 12, height: 12}
	fmt.Println("s =", c.Area())

}

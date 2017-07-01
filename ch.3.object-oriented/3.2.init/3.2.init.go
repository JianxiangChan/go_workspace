package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

//go语言没有构造函数，需要自己创建一个全局的创建函数来完成
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
func main() {
	//任何未被初始化的函数都会被设为“0”值
	//bool 为false int为0 string为空
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	rect5 := NewRect(0, 0, 100, 200)

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)
	fmt.Println(rect5)
}

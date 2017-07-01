package main

//Golang的log包短小精悍，可以非常轻松的实现日志打印转存功能。
import (
	"fmt"
	//	"log"
)

//定义一个base类
type Base struct {
	Name string
}

func (base Base) Foo() {
	//do things
	fmt.Println("func Base.Foo")
	return
}

func (base Base) Bar() {
	//do things
	fmt.Println("func Base.Bar")
	return
}

type Name struct {
	Name string
}

//定义了一个Foo类，并继承了Bar方法
type Foo struct {
	//Base
	*Base
	//只访问最外层的成员
	Name string
	//但是同一类里面不能有两个相同名字的成员

	//Name这里就会报错 但是这个错误不是一定发生的 如果没有使用这两个成员的话 这里就需要自己去确认
}

func (foo *Foo) Bar() {
	fmt.Println("func Foo.Bar")
	foo.Base.Bar()

}

//type Job struct {
//	Command string
//	*log.Logger
//}

//func (job *Job) Start() {

//	job.Logger

//}

func main() {
	//这里注意了 如果以指针来进行派生 就需要一个实例化的Base指针，
	//不然就内存错误了
	a := Foo{}
	a.Base = new(Base)
	a.Foo()
	a.Bar()
	a.Base.Bar()
	a.Base.Foo()
	fmt.Println(a.Base.Name, a.Name)
}

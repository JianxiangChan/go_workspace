package addless

//import "fmt"

type Integer int

//3.5.3 接口赋值
func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func (a *Integer) Modify(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

var a Integer = 1

//1.将对象实例赋值给接口
var b LessAdder = &a

//这里会编译错误 因为这里的add不会自动生成一个 func (a Integer)Add(b Integer)
//var b LessAdder = a

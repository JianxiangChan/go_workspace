package vis

//go的可见性是通过符号大小写来决定的
//其访问控制是包一级的 同一个包类其他类型也可以访问私有成员
type Rect struct {
	x, y          float64
	width, height float64
}

//包外不可见
func (r *Rect) Area() float64 {
	return r.width * r.height
}

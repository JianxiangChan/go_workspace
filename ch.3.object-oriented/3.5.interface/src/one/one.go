package one

type File struct {
	Name string
}

//var A int

func (f *File) Read(buf []byte) (n int, err error) {
	return
}
func (f *File) Write(buf []byte) (n int, err error) {
	return
}

type Read interface {
	Read(buf []byte) (n int, err error)
}
type Write interface {
	Write(buf []byte) (n int, err error)
}
type ReadWrite interface {
	//	Read(buf []byte) (n int, err error)
	//	Write(buf []byte) (n int, err error)
	//3.5.6 接口组合
	Read //只包含方法 不包含成员变量
	Write
}

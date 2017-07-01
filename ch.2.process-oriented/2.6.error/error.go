package main

import "os"
import "fmt"

type PathError struct {
	Op   string
	Path string
	Err  error
}

//自定义的error方法
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func main() {

	//2.6.1 错误处理
	fi, err := os.Stat("a.txt")
	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			fmt.Println(fi)
		}
	}

}

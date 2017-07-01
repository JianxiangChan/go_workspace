package mymath

import "errors"

// 2.5.1 函数定义
func Add(a int, b int) (ret int, err error) { //注意 大写的字母开头才能被其他包调用！简直坑爹！
	// func add(a, b int)(ret int, err error)    //参数类型相同
	// func add(a, b int)int //只有一个返回值
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil //支持多重返回值
}

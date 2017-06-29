package main

import "fmt"

func main() {
	a := 3
	a = example_ifelse(a)
	fmt.Println(a)
	example_switchcase(a)
	example_for()
	example_goto()
}

//2.4.1 if&else 条件语句

//奇怪的是我这里是可以正常编译的，书上这里说return是不能出现在函数的if 和else的中间的
//个人猜测是 安装的go的版本不同，go的新特性支持这种写法，我这是官网下载的最新版本
func example_ifelse(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

//2.4.2选择语句

//switch case fallthrough
//表达式中不设定是常量或者整数，
//不需要break来明确跳出
//fallthrough才会继续执行下一个case
//switch后面如果没有表达式，就和if&else一样
func example_switchcase(x int) {
	temp := x
	switch temp {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4,5,6")
	//如果switch 后面加了表达式 就不能使用判断的方式来执行，否则报错类型不匹配
	//	case 7 <= x && x <= 9:
	//		fmt.Println("7-9")
	default:
		fmt.Println("Default")
	}
}

//2.4.3循环语句
//只支持平行赋值的方法
func example_for() {
	sum := 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[i], a[j]
	}
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop //打断J这一层循环
			}
			fmt.Println(i)
		}
	}
}

//2.4.4 跳转语句

//谨慎使用
func example_goto() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

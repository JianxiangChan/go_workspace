package main

import "fmt"
import "math"

//数值运算 和C一样
/*
布尔类型bool
整型 int8 byte int16 int uint uintptr int32 uint32 int64 uint64
浮点类型 float32 float64
复数类型 complex64 complex128
字符串 string
字符类型 rune
错误类型 error
指针 pointer
数组 array
切片 slice
字典 map
结构体 struct
接口 interface
*/
func main() {
	/*
		不能自动或者强制类型转换
		var b bool
		b = 1
		b = bool(1)
	*/
	var b bool
	b = (1 != 0)
	fmt.Println("Result:", b)

	var value2 int32
	value1 := 64           //这里编译器自动认为是int型
	value2 = int32(value1) //go 不认为int32 和 int是同一个类型 所以要强制类型转换
	fmt.Println("Result:", value2)

	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal")
	}

	//可以直接与字面常量进行比较
	var k int32
	var l int64
	k, l = 1, 2
	/*
		if k == l {
			fmt.Println("i & j is equal.")
		}
	*/
	if k == 1 || l == 2 {
		fmt.Println("i and j are equal.")
	}

	//运算符 异或取反 ^

	//2.3.3浮点型
	//floa32 想当时 C的float float64相当于double
	var fvalue1 float32
	fvalue1 = 12
	//编译器自动推导成double型
	fvalue2 := 12.0
	IsEqual(float64(fvalue1), fvalue2, 0.00000001)
	//fvalue1 = fvalue2 所以这句话会报错！！！！
	fmt.Printf("fvalue1 = %f,fvalue2 = %f\r\n", fvalue1, fvalue2)

	//2.3.4复数类型
	var mvalue1 complex64
	mvalue1 = 3.2 + 12i
	mvalue2 := 3.2 + 12i
	mvalue3 := complex(3.2, 12)
	fmt.Println(mvalue1)
	fmt.Println(mvalue2)
	fmt.Println(mvalue3)
	fmt.Println(real(mvalue1))
	fmt.Println(imag(mvalue1))

	//2.3.5 字符串

	//字符串操作
	var str string
	str = "hello world"
	str1 := "你好，世界"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str1, len(str1))
	fmt.Printf("The first character of \"%s\" is %c. \n", str, ch)
	//字符串的内容可以用数组下标的方式取得，但是不能赋值和修改
	//str[0] = 'x' //错误
	//Go支持UTF-8和Unicode
	str2 := str + str1
	fmt.Printf("The length of \"%s\" is %d \n", str2, len(str2))
	fmt.Printf("The second character of \"%s\" is %c. \n", str, "hello"[1])

	//字符串遍历
	n := len(str2)
	for i := 0; i < n; i++ {
		ch := str2[i]
		fmt.Println(i, ch)
	}
	//range 遍历数组并返回数组下标和实际值
	for i, ch := range str2 {
		fmt.Println(i, ch)
	}

	//字符类型 byte 和rune

	//数组
	/*
		[32]byte
		[2*N] struct{x, y int32} //复杂类型数组
		[1000]*float64 //指针数组
		[3][5]int
		[2][2][2]float64 //等同于[2]([2]([2]float64))
		数组长度是不可改变的
		长度获取和字符串一样
	*/
	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(),array values:", array)

	//2.3.8切片数组

	//1.创建切片数组
	//基于已创建的数组创建
	//这里 array[a:b]取得是 array[a] 到 array[b-1] 的元素
	var mySlice []int = array[1:4]
	fmt.Println("Element of mySlice")

	//_,这里表示省略
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//直接创建
	mySlice1 := make([]int, 5)

	fmt.Println("Element of mySlice1")
	for _, v := range mySlice1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//动态增减元素
	mySlice2 := make([]int, 5, 10) //预留十个存储空间
	mySlice2 = append(mySlice2, 1, 2, 3)

	fmt.Println("Element of mySlice2")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice4 := make([]int, 2)
	mySlice2 = append(mySlice2, mySlice4...) //注意...的用法 相当于打散mySlice4元素的值赋值给切片数组。这里不能少了

	fmt.Println("Element of mySlice2")
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice3 := []int{1, 2, 3, 4, 5} //直接创建并初始化
	fmt.Println("Element of mySlice3")
	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//数组切片给人最直观的感受就是[]里面是空的 而数组是固定的
	//基于切片数组创建切片数组
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3]
	fmt.Println("Element of newSlice")
	for _, v := range newSlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//2.3.9 map
	type PersonInfo struct {
		ID      string
		Name    string
		Address string
	}

	//这里的初始化语法不太明白
	//personDB 为变量名 string 为键的类型， PersonInfo则是其中所存放的值类型
	var personDB map[string]PersonInfo
	//make()
	//personDB = make(map[string] PersonInfo)
	personDB = make(map[string]PersonInfo, 100) //指定存储能力
	//插入map信息
	personDB["1234"] = PersonInfo{"1234", "Tom", "Room 203, ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101, ..."}

	//从这个map查找键为“1234”的信息
	person, ok := personDB["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "wiht ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	delete(personDB, "1234")

	person, ok = personDB["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "wiht ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}
func IsEqual(f1 float64, f2 float64, p float64) bool {
	return math.Abs(f1-f2) < p
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(),array values:", array)
}

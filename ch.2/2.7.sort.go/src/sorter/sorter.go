/*主程序程序要求
1. 获取并解析命令行的输入
2. 从对应文件中读取输入数据
3.对用对应的排序函数
4，将排序的结果输入到对应的文件中去
5.打印排序所花费的时间
*/

package main

//命令行解析接口
import "flag"
import "fmt"

// bufio 包实现了带缓存的 I/O 操作
// 它封装一个 io.Reader 或 io.Writer 对象
// 使其具有缓存和一些文本读写功能
import "bufio"

//io 包为I/O原语提供了基础的接口.
//它主要包装了这些原语的已有实现，如 os 包中的那些，抽象成函数性的共享公共接口，加上一些其它相关的原语。
//由于这些接口和原语以不同的实现包装了低级操作，因此除非另行通知，否则客户不应假定它们对于并行执行是安全的。
import "io"

//os包中实现了平台无关的接口，设计向Unix风格，但是错误处理是go风格，
//当os包使用时，如果失败之后返回错误类型而不是错误数量．
import "os"

//基本字符串类型的转换包
import "strconv"

//对输入的命令格式定义
var infile *string = flag.String("i", "infile", "File contains value for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted value")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	//打开infile文件
	file, err := os.Open(infile)
	if err != nil {
		//打开文件失败
		fmt.Println("Failed to open the inputfile", infile)
		return
	}
	//确保关闭文件句柄
	defer file.Close()

	//读取文件
	br := bufio.NewReader(file)

	//创建一个长度为0的切片数组
	values = make([]int, 0)

	//一行一行解析数组里面的内容
	for {
		//逐行读取文件
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			//当读到了 end of file帧后 不返回err 而直接返回nil 表示读取正常
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected") //缓冲溢出
		}

		//将读到的字符串一步步转化为数组
		str := string(line) //转换字符数组为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value) //放到values里面
	}

	return
}

func writeValues(values []int, outfile string) error {
	//创建一个以outfile为名字的dat
	file, err := os.Create(outfile)

	if err != nil {
		fmt.Println("failed to creat the output file", outfile)
		return err
	}

	defer file.Close()

	//将数组转化为字符串然后末尾换行写入到outfile
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	//解析命令行
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}

	writeValues(values, *outfile)
}

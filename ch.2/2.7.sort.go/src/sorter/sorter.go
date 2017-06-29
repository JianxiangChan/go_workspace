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

func main() {
	//解析命令行
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}
}

//测试的包的函数必须要与源程序相同
package bubblesort

//go 自带的testing包 直接可以进行单元测试、性能分析、输出结果验证
import "testing"

//函数以Test开头 且 函数第一个字母大写
func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 ||
		values[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 3 5")
	}
}
func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 5")
	}
}

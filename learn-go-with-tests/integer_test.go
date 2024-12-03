package learn_go_with_tests

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

// note: 示例
// 示例是存在于一个包的`_test.go`文件中的函数，如果你的代码发生改变，导致示例不再有效，
// 那么你的构建也会失败。
// 如果删除注释`// Output:`，示例函数将不会执行。虽然函数会被编译，但是它不会执行。
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	// 封装一个函数来进行断言
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// note: `t.Helper()`告诉测试套件这个是辅助函数，通过这样做，当测试失败时所报告当行号将是函数当调用行！
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("chris")
		want := "hello, chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		assertCorrectMessage(t, got, want)
	})
}

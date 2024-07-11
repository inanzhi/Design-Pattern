package benchmark_test

import (
	"github.com/inanzhi/factualcase/learn/gotest/benchmark"
	"testing"
)

// 执行方法 go test -bench=.  //好像不能执行所有的test
// go test -bench=. <import-path>
// go test -run=^$ -bench=.
// go test -bench=. -benchmem
// 记住这些命令应该在你的Go项目的根目录下执行，或者在包含有基准测试的包的目录下执行。
// -bench 匹配性能测试  .为flag的参数  匹配所有的性能测试
// 基准测试/性能测试的规则是BenchmarkXxx
func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	//b.N表示循环执行次数  N值是动态调整的，底层可靠的计算出程序的执行时间后会停止
	for i := 0; i < b.N; i++ {
		benchmark.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmark.MakeSliceWithPreAlloc()
	}
}

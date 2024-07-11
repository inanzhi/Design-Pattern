package example_test

import "github.com/inanzhi/factualcase/learn/gotest/example"

// go test example_test.go
// go test执行所有的测试 包括Example
// 使用OutPut:检测 output不分大小写  比较字符串时会忽略前后的空白字符
// 无output该函数无法执行
// 只有三种检测
// 检测单行输出
func ExampleSayHello() {
	example.SayHello()
	//OutPut:    Hello World
}

// 检测多行输出  //注意名称是Example+待测方法名  名称不对 匹配不上
func ExampleSayGoodbye() {
	example.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	example.PrintNames()
	// Unordered Output:
	// Jim
	// Bob
	// Tom
	// Sue
}

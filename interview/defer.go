package interview

import "fmt"

// 打印结果为：
// 打印后
// 打印中
// 打印前
// 考察defer语句与panic()函数的执行流程，
// defer是后进先出，而panic(运行时恐慌)需要等待defer结束后才会向上传递，
// 本例中会先按照先进后出的顺序执行完defer语句，最后才会执行panic
func Defer_call() {
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()

	panic("触发异常")
}



package channel

import (
	"fmt"
	"time"
)

// channel的使用
func ChannelUse() {
	// 初始化一个存放数据的Channel作为数据队列
	dataChan := make(chan int, 5)

	// 向channel中写入数据
	go func(datas chan<- int) {
		for i := 1; i <= 10; i++ {
			dataChan <- i
			fmt.Println("put: ", i)
		}
	}(dataChan)

	// 从channel中读取数据
	go func(datas <- chan int) {
		for {
			var b int
			b = <- dataChan
			fmt.Println("read: ", b)
		}
	}(dataChan)

	time.Sleep(time.Second * 3)
}

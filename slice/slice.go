package slice

import "fmt"

// 切片的理解
// 每次len达到底层数组的cap时进行扩容
// 每次扩容后底层数组长度扩为原来的2倍
func SliceUse() {
	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("append 之前 s: ", s)
	fmt.Println("s.len: ", len(s))
	fmt.Println("s.cap: ", cap(s))

	for i := 10; i <= 20; i++ {
		s = append(s, i)
		// 每次append后输出s
		fmt.Printf("------------第%d次append-------------", (i-9))
		fmt.Println("s: ", s)

		// 每次append后输出s的len和cap
		fmt.Println("s.len: ", len(s))
		fmt.Println("s.cap: ", cap(s))
		fmt.Println("-------------------------")
	}
}

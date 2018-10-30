package interview

import "fmt"

type Student struct {
	Name string
	Age  int
}

func ParseStudent() {
	stuDict := make(map[string]*Student)
	stus := []Student{
		{Name: "jeason", Age: 24},
		{Name: "jack", Age: 25},
		{Name: "rose", Age: 23},
	}

	// 错误的foreach写法，结果是：stuDict字典中三个key——jeason, jack, rose对应的value都将是最后一次迭代的值——{rose 23}
	for _, stu := range stus {
		stuDict[stu.Name] = &stu
	}

	// 正确的foreach赋值方法如下：
	for i, stu := range stus {
		stuDict[stu.Name] = &stus[i]
	}

	for k, v := range stuDict {
		fmt.Println(k, " : ", *v)
	}
}

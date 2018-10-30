package json

import (
	"encoding/json"
	"fmt"
)

// json解析库的使用
// 使用json.Unmarshal(data []byte, v interface{})将字节数组解析为某个具体的类型变量
// 使用json.Marshal(v interface{})将某个类型的变量序列化为json字节数组
func JsonUnmarshal() {
	var bts []byte
	bts = []byte("{\"name\":\"jeason\", \"age\":24, \"sex\":\"male\"}")

	var user User
	err := json.Unmarshal(bts, &user)
	if err != nil {
		fmt.Println("json.Unmarshal error: ", err)
	} else {
		fmt.Println("user: ", user)
	}
}

func JsonMarshal() {
	var user = User{
		Name: "jeasonwang",
		Age:  24,
		Sex:  "male",
	}

	data, err := json.Marshal(&user)
	if err != nil {
		fmt.Println("json.Marshal error: ", err)
	} else {
		fmt.Println("data: ", data)
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

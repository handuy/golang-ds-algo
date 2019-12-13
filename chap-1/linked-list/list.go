package main

import (
	"fmt"
	"container/list"
)

func main(){
	var strList list.List
	strList.PushBack("Java")
	strList.PushBack("Golang")
	strList.PushBack("Python")

	// B1: lấy ra element số 1 trong list
	// B2: Kiểm tra xem element đó có nil hay không. Nếu là nil chứng tỏ list này là rỗng
	// B3: In ra giá trị của element. Sau đó chuyển sang element kế tiếp bằng cách gọi method Next()
	// B4: Lặp lại B2 và B3 cho đến hết List
	for el := strList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
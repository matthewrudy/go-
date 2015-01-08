package 動物

import (
	"fmt"
)

type 動物 interface {
	叫() T字符串
}

type T字符串 string

func F摩(朋友 動物) {
	fmt.Println("你：可愛的動物")
	fmt.Println("牠：", 朋友)
}

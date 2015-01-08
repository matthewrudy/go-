package 動物

import (
	"fmt"
	"github.com/matthewrudy/go-zh/类型"
)

type 動物 interface {
	F叫() 类型.T字符串
}

func F摩(朋友 動物) {
	fmt.Println("你： 可愛的動物")
	fmt.Println("牠：", 朋友.F叫(), "!!!")
}

package main

import (
	"github.com/matthewrudy/go-zh/動物"
	"github.com/matthewrudy/go-zh/类型"
)

type T狗 struct {
}

func (g *T狗) F叫() 类型.T字符串 {
	return "汪汪"
}

func main() {
	汪汪 := &T狗{}
	動物.F摩(汪汪)
}

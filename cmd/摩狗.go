package main

import (
	"github.com/matthewrudy/go-zh/動物"
)

type 狗 struct {
	名字 string
}

func (g *狗) 叫() 動物.T字符串 {
	return "汪汪"
}

func main() {
	汪汪 := 狗{}
	動物.F摩(汪汪)
}

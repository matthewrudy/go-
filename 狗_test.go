package go中文

import (
	"testing"
)

func Test_叫(t *testing.T) {
	狗 := 狗{
		名字: "盤瓠",
	}

	話 := 狗.叫()

	if 話 != "汪汪" {
		t.Error("狗不是說‘汪汪’嗎？", 話)
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	// \ 反斜杠具有特殊含义，告诉程序字符串中的反斜杠是一个单纯的反斜杠，在 \ 前再加 \
	path := "C:\\gocode\\src\\github.com\\zhangliying0113\\studygo\\day01"
	fmt.Println(path)

	s := "I'm OK"
	fmt.Println(s)

	// 多行字符串
	s2 := `
	世界那么大
		我想去看看
	`
	fmt.Println(s2)

	s3 := `C:\gocode\src\github.com\zhangliying0113\studygo\day01`
	fmt.Println(s3)

	// 字符串操作
	fmt.Println(len(s3))

	name := "神奇"
	world := "的世界"

	ss := name + world
	fmt.Println(ss)

	ss1 := fmt.Sprintf("%s%s", name, world)

	fmt.Printf("%s%s", name, world)

	fmt.Println(ss1)

	// 分隔
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "世界"))
	fmt.Println(strings.Contains(ss, "世上"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "神奇"))

	// 后缀
	fmt.Println(strings.HasSuffix(ss, "神奇"))

	// 字符位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	// 将 [C: gocode src github.com zhangliying0113 studygo day01] 元素用 ++ 拼接
	// 结果：C:++gocode++src++github.com++zhangliying0113++studygo++day01
	fmt.Println(strings.Join(ret, "++"))
}

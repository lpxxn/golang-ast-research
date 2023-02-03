package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	var src = []byte(`println("hello world")`)

	var fSet = token.NewFileSet()
	var file = fSet.AddFile("hello.go", fSet.Base(), len(src))

	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fSet.Position(pos), tok, lit)
	}
	fmt.Println("=====================================")
	/*
		Filename表示文件名，Offset对应文件内的字节偏移量(从0开始)，Line和Column对应行列号(从1 开始)。比较特殊的是Offset成员，它用于从文件数据定位代码，但是输出时会将偏移量转换为行列号 输出。
		打印位置信息时，根据文件名、行号和列号有6种组合:
	*/
	a := token.Position{Filename: "hello.go", Line: 1, Column: 2}
	b := token.Position{Filename: "hello.go", Line: 1}
	c := token.Position{Filename: "hello.go"}

	d := token.Position{Line: 1, Column: 2}
	e := token.Position{Line: 1}
	f := token.Position{Column: 2}
	fmt.Println(a.String())
	fmt.Println(b.String())
	fmt.Println(c.String())
	fmt.Println(d.String())
	fmt.Println(e.String())
	fmt.Println(f.String()) // Output: -  行号从1开始，是必须的信息，如果缺少行号则输出“-”表示无效的位置。
}

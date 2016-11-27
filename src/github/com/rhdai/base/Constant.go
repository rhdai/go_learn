package main

import "fmt"

func main() {
	/*
	    通过 const 关键字定义常量
	    Go语言预定义了这些常量: true 、 false 和 iota
	    iota 比较特殊,可以被认为是一个可被编译器修改的常量,在每一个 const 关键字出现时被重置为0,
	    然后在下一个 const 出现之前,每出现一次 iota ,其所代表的数字会自动增1。
	*/
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0// 无类型浮点常量
	const (
		size int64 = 1024
		eof = -1
		// 无类型整型常量
	)
	const u, v float32 = 0, 3// u = 0.0, v = 3.0,常量的多重赋值
	const a, b, c = 3, 4, "foo"// a = 3, b = 4, c = "foo", 无类型整型和字符串常量

	const(
		c0 = iota;
		c1 = iota;
		c2 = iota;
	)
	const(
		c3 = iota;
	)
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

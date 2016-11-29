package main

import "fmt"

func main() {
	/* 左花括号{ 必须与for处于同一行。
	 Go语言中的for循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别
	是， Go语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多
	个变量。*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// Go语言的for循环同样支持continue和break来控制循环，但是它提供了一个更高级的
	//break，可以选择中断哪一个循环，如下例：

}

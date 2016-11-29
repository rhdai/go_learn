package main

import "fmt"

func main() {
	/* 左花括号{ 必须与switch处于同一行；
	 条件表达式不限制为常量或者整数；
	 单个case中，可以出现多个结果选项；
	 与C语言等规则相反， Go语言不需要用break来明确退出一个case；
	 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
	 可以 不设定 switch 之 后的条 件表达式， 在此种情况 下，整个 switch 结构与 多个
	 if...else... 的逻辑作用等同。*/
	var a int =1
	switch a {
	case 0:
		fallthrough
	case 1:
		fmt.Println("1")
		fallthrough //执行下一个case
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}
}

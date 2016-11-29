package main

import "fmt"

func main() {
	/*条件语句不需要使用括号将条件包含起来() ；
	  无论语句体内有几条语句，花括号{} 都是必须存在的；
	  左花括号{ 必须与if或者else处于同一行；
	  在if之后，条件语句之前，可以添加变量初始化语句，使用; 间隔；*/
	var a int = 1;
	if a>0 {
		fmt.Println("a是正数")
	}else{
		fmt.Println("a是负数")
	}
	fmt.Println(example(0))
}
func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

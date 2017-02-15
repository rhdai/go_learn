package main

import "fmt"
/**
	规则：小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用。
		这个规则也适用于类型和变量的可见性

 */

func main() {
	sum1 := Add1(1, 2);
	fmt.Println(sum1);
	sum2 := Add2(1, 2);
	fmt.Println(sum2);
	myfunc1(1,2,3,4);
	myfunc2(1,2.0,3,true);
	MyPrintf(1,2.0,3,true);
}

func Add1(a int,b int)(sum int){
	return a + b;
}

func Add2(a,b int)int {
	return a+b;
}

func myfunc1(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc2(args ... interface{})  {//interface{}代表能接受任何参数
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args... interface{}){
	for _,arg := range args{
		switch arg.(type) {
			case int:{
				fmt.Println(arg , "is int")
			}
			case string:{
				fmt.Println(arg , "is string")
			}
			case bool:{
				fmt.Println(arg , "is bool")
			}
		}
	}
}
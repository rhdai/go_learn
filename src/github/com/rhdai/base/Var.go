package main

import "fmt"

func main()  {
	/*
	    1.变量定义
		关键字 var ,类型信息放在变量名之后;
		可以将若干个需要声明的变量放置在一起,免得程序员需要重复写 var 关键字
		变量声明语句不强制使用分号作为结束符
	*/
	var v1 int;
	var v2 string;
	var v3 [10]int;//数组
	var v4 []int;//数组切片
	var v5 struct{
		f int;
	}
	var v6 *int;//指针
	var v7 map[int]string;//map key是int value是string
	var v8 func(a int) int;

	var(
		v9 int;
		v10 string;
	)
	fmt.Println(v1);
	fmt.Println(v2);
	fmt.Println(v3);
	fmt.Println(v4);
	fmt.Println(v5);
	fmt.Println(v6);
	fmt.Println(v7);
	fmt.Println(v8);
	fmt.Println(v9);
	fmt.Println(v10);

	/*
	    变量初始化
		(冒号和等号的组合 := ),用于明确表达同时进行变量声明和初始化的工作
		go语言支持多赋值
	*/
	var v11 int = 11// 正确的使用方式1
	var v12 = 12// 正确的使用方式2,编译器可以自动推导出v12的类型
	v13 := 13// 正确的使用方式3,编译器可以自动推导出v13的类型
	var v14 int;
	v14 = 14;
	fmt.Println(v11);
	fmt.Println(v12);
	fmt.Println(v13);
	fmt.Println(v14);
	//多重赋值
	v13,v14 = v14,v13;
	fmt.Println(v13);
	fmt.Println(v14);
	var nickname string;
	_, _, nickname = GetName();
	fmt.Println(nickname);
}

func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

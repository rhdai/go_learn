package main

import "fmt"

func main() {
	/*
	    创建数组切片的方法主要有两种:
	    	1.基于数组
	    	2.直接创建  使用内置函数make()
	    	3.基于数组切片创建数组切片
	    	4.内容复制
	*/

	//定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//基于数组创建数组切片
	var mySlice []int = myArray[1:5];

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//创建一个初始元素个数为5的数组切片,元素初始值为0:
	mySlice1 := make([]int, 5)
	//创建一个初始元素个数为5的数组切片,元素初始值为0,并预留10个元素的存储空间:
	mySlice2 := make([]int, 5, 10)
	//直接创建并初始化包含5个元素的数组切片:
	mySlice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice1, mySlice2, mySlice3)


	/*
	    数组和数组切片元素遍历的两种方式:
	    	循环
	    	range关键词
	*/
	for i := 0; i < len(mySlice3); i++ {
		fmt.Println("方式一：mySlice3[", i, "] =", mySlice3[i])
	}

	for i, v := range mySlice3 {
		fmt.Println("方式二：mySlice3[", i, "] =", v)
	}

	/*
	    len() 函数返回的是数组切片中当前所存储的元素个数
	    cap() 函数返回的是数组切片分配的空间大小
	    append() 继续新增元素
	*/
	fmt.Println("len(mySlice2):", len(mySlice2))
	fmt.Println("cap(mySlice2):", cap(mySlice2))

	mySlice2 = append(mySlice2, 1, 2, 3)
	for i, v := range mySlice2 {
		fmt.Println("mySlice3[", i, "] =", v)

	}
	mySlice4 := []int{8, 9, 10}
	// 给mySlice2后面添加另一个数组切片
	mySlice2 = append(mySlice2, mySlice4...)
	for i, v := range mySlice2 {
		fmt.Println("mySlice3[", i, "] =", v)

	}

	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:3] // 基于oldSlice的前3个元素构建新数组切片
	printArry(newSlice);


	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	printArry(slice1);
	printArry(slice2);
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	printArry(slice1);
	printArry(slice2);

}

func printArry(arr []int) {
	for i, v := range arr {
		fmt.Println("arr[", i, "] =", v)

	}
}

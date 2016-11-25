package main

import (
	"fmt"
	//"os"
)

func main() {
	fmt.Println("hello, world.你好,世界");
 	var a int = 8;
	var b int = 6;
	const pi float64 = 3.141592653
	fmt.Println(pi)
	const p  = 2<<3;
	fmt.Println(p)
//	const home  = os.Getenv("path")
	a,b = b,a;
	fmt.Println(a,b);
	var c,d int = testfunc(a,b);
	fmt.Println(c,d)
//	var nickName string;
//	_,_,nickName := getName();
	fmt.Println(getName())
}

func testfunc(a int,b int)(add int ,sub int) {
	return a+b,a-b;
}

func getName()(firName string,secName string,nickName string){
	return "a","b","c"
}
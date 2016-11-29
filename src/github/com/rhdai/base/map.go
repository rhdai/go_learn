package main

import "fmt"

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func main() {
	var personDB map[string] PersonInfo;//申明变量
	personDB = make(map[string] PersonInfo)//创建map
	//personDB = make( map[ string] PersonInfo, 100)//创建一个初始化容量为100的map
	//personDB = map[ string] PersonInfo{
	//	"1234": PersonInfo{"1", "Jack", "Room 101,..."},//创建并初始化map
	//}
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}//相当于java里的map.put("key",value)
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	// 从这个map查找键为 "12345" 的信息
	//person, ok := personDB["12345"]
	//if ok {
	//	fmt.Println("Found person", person.Name, "with ID 1234.")
	//} else {
	//	fmt.Println("Did not find person with ID 1234.")
	//}
	delete(personDB, "12345")//删除key为12345的键值对

	person, ok := personDB["12345"]//相对于java更简单明了
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}

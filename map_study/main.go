package main

import (
	"fmt"
	"sort"
)

func main() {
	//map有两种初始化方法

	//方法一：先声明，再初始化或者通过:=将声明和初始化合并为一条语句
	var testMap map[string]int
	//testMap["zero"] = 0,此处由于testMap仅仅是声明，还未初始化，为nil，赋值会panic
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(testMap) //map是无序的，每次打印结果当中顺序可能不一
	testMap["four"] = 4
	fmt.Println(testMap)

	//方法二：通过make来初始化一个新字典
	var testMap1 = make(map[string]int)
	testMap1["zero"] = 0 //通过make创建的map，可以直接赋值，因为make就是分配内存空间
	fmt.Println(testMap1)

	//查找元素
	value, ok := testMap["one"]
	if ok {
		fmt.Println("search one value ", value)
	}

	value = testMap["two"]
	fmt.Println("two value is ", value) //一个返回值，返回value
	value = testMap["five"]
	fmt.Println("five value is ", value) //key不存在时不会panic或者error，返回map的value的零值

	//删除元素
	delete(testMap, "one")
	fmt.Println("delete one, current testMap is ", testMap) //如果key不存在或者map尚未初始化，delete调用不会报错

	//遍历
	for key, value := range testMap {
		fmt.Println("key is ", key, "value is ", value)
	}

	for _, value = range testMap {
		fmt.Println("value is ", value)
	}

	for key := range testMap { //只获取键名
		fmt.Println("key is ", key)
	}

	//由于字典是无序的，可以通过对键和值创建slice，然后对slice进行排序后来实现
	keys := make([]string, 0)
	for k, _ := range testMap {
		keys = append(keys, k)
	}
	sort.Strings(keys) //使用sort.Strings对string类型的slice排序，还有sort.Ints对int类型的slice进行排序
	for _, k := range keys {
		fmt.Println("sorted: ", k, testMap[k])
	}
}
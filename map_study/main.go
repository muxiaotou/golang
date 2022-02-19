package main

//https://www.cnblogs.com/snowInPluto/p/7477365.html

import (
	"fmt"
	"sort"
)

func main() {
	//map有两种初始化方法

	//方法一：先声明，再初始化或者通过:=将声明和初始化合并为一条语句
	var testMap map[string]int
	//testMap["zero"] = 0,此处由于testMap仅仅是声明，还未初始化，为nil，赋值会panic
	//nil的map，或者key不存在的map读，均会返回map对应value的零值
	fmt.Println(testMap["bbb"])
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

	//map初始化容量(len)为3，实际可以超过3,这一点和slice不一样
	//map make时len不是必须的，当有新元素添加时，都会检查是否需要扩容
	//slice是make是必须指定len，因此确定了内存空间范围，直接赋值时超过空间大小就越界了，slice提供了append操作，新元素
	//添加时，会检查是否需要扩容
	testMap2 := make(map[string]int, 3)
	testMap2["one"] = 1
	testMap2["two"] = 2
	testMap2["three"] = 3
	testMap2["four"] = 4
	fmt.Println(testMap2)

	//map初始赋值，一行和多行的区别：最后一行是否有逗号分隔
	m := map[string]int{"beijing": 1, "xi'an": 2}
	m1 := map[string]int{
		"beijing": 1,
		"xi'an":   2,
	}
	fmt.Println(m, m1)
}

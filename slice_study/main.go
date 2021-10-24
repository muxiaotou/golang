package main

import "fmt"

//https://www.cnblogs.com/qcrao-2018/p/10631989.html
func main() {
	//golang可以基于数组、切片，以及直接创建三种方式来创建切片
	//先定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	//基于数组创建切片
	q2 := months[3:6]
	summber := months[5:8]

	fmt.Println(q2, summber)

	//直接创建
	mySlice1 := make([]int, 5) //分配了内存空间，打印结果是对应元素的零值，即[0,0,0,0,0]
	mySlice2 := []int{1, 2, 3, 4, 5}
	var mySlice3 []int //没有分配内存空间，打印结果是[],即nil
	fmt.Println(mySlice1, mySlice2, mySlice3)
	if mySlice3 == nil {
		fmt.Println("mySlice3 is nil")
	}

	//len()获取元素个数来遍历
	for i := 0; i < len(months); i++ {
		fmt.Println("summer[", i, "]=", months[i])
	}

	//range关键字来遍历
	for i, v := range months {
		fmt.Println("summer[", i, "]=", v)
	}

	//只打印索引
	for i := range months {
		fmt.Println(i)
	}

	//只需要value
	for _, v := range months {
		fmt.Println(v)
	}

	//slice的长度len和容量cap
	//1)对于济源数组和切片创建的切片而言，默认容量是从新切片起始索引到对应底层数组的结尾索引
	//2)对于通过内置make函数创建的切片而言，在没有指定容量参数的情况下，默认容量和切片长度一致
	var oldSlice = make([]int, 5, 10)
	fmt.Println("oldSlice len is ", len(oldSlice))
	fmt.Println("oldSlice cap is ", cap(oldSlice))

	//oldSlice，默认值是[0,0,0,0,0]，append后结果为[0,0,0,0,0,1,2,3,4]
	newSlice := append(oldSlice, 1, 2, 3, 4)
	fmt.Println(newSlice)

	//直接append一个slice，注意末尾的...符号
	appendSlice := []int{2, 3, 4, 5, 6, 7}
	newSlice = append(oldSlice, appendSlice...)
	fmt.Println(newSlice)

	//自动扩容
	//如果通过append追加的元素个数超出oldSlice的默认容量，则底层会自动进行扩容，扩容并不会改变原来的slice，而是会生成一个容量更大的切片
	//然后把原有的元素和新元素拷贝到新切片中

	//go没有提供python类似的del函数来删除特定的元素，但是可以通过append来实现插入及删除操作
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8}

	tmp := append([]int{}, slice3[3:]...)            //为插入操作，需要新建一个临时的slice来保存后续的元素数据
	slice4 := append(append(slice3[:3], 10), tmp...) //插入

	slice5 := append(slice3[:4], slice3[5:]...) //删除
	fmt.Println(slice4, slice5)

	fmt.Println(len(slice5), cap(slice5))
	slice6 := append(slice5[:3], []int{9, 8, 7, 6, 5, 4, 3, 2, 1}...)
	fmt.Println(slice5, slice6)

	//上面需要新建一个tmp临时的slice，是因为slice是基于底层数组实现的，slice3\slice4\slice5数组指针都指向了同一个数组，因此对任意slice的
	//修改均会影响所有slice
	//除非某个slice的cap不够了，append则会重新生成了底层数组,而不会影响原有的slice，例如上面最后打印的slice5和slice6

}

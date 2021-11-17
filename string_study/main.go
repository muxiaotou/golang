package main

import (
	"fmt"
	"reflect"
)

func main() {
	//使用+连接字符串
	fmt.Println("i am" + "mutouchen")

	//使用Sprintln连接字符串
	s := fmt.Sprintln("aaaaa", "bbbbbb")
	fmt.Println(s)

	//go支持两种方式遍历字符串
	//字节数组的方式遍历
	str := "Hello,世界"
	for i := 0; i < len(str); i++ {
		fmt.Println(reflect.TypeOf(str[i])) //for的valve是uint8，byte是uint8的别名
		fmt.Println(i, str[i])              //依据下标取字符串中的字符，str[i]的类型为byte，因为字符串的底层实际就是[]byte，通过len获取的是字符串的字节长度，是从ascii字符集的角度切入
	}
	//此处len(str)预期应该是8个字符，实际len(str)是12，因为字母acsii字符集占1个byte，汉字unicode字符集占2个byte，而golang
	//使用的是UTF-8编码规范，汉字占3个byte，因此是12
	//UTF-8是正对unicode的一种可边长的字符编码，字符占1个byte，汉字占了3个byte

	//以unicode字符方式遍历
	for i, ch := range str {
		fmt.Println(reflect.TypeOf(ch)) //range的value是int32，rune是int32的别名
		fmt.Println(i, ch)              //ch的类型为rune，通过range遍历字符串是从unicode字符集的角度切入
	}
	//range打印出来的是unicode字符编码，需要使用string(ch)的方式可以转换为中文

	//以上，[]byte []rune可以和string进行类型互转

	//golang提供strconv包来实现基本数据类型与字符串之间的转化
	//strconv.Atoi() 字符串转化为整形
	//strconv.Itoa() 整形转化为字符串

	//只打印索引
	for i := range str {
		fmt.Println(i)
	}
}

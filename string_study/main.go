package main

//string默认是UTF-8编码的unicode字符集
//unicode字符集包含汉字
//ascii编码规则：ascii字符集不包含汉字，每个ascii码占用1个byte来存储，byte的最高bit置0
//utf-8编码规则当中，对于ascii字符集的编码规则和ascii编码规则一致

//字符：
//各种文字和符号的的总称，可以是一个中文汉字、一个英文字母、一个阿拉伯数字、一个标点符号等

//字符集：
//多个字符的集合。如128个字符的ASCII字符集，GB2312的简体中文字符集，unicode字符集等
//ASCII字符集有ASCII编码映射到计算机进行存储和传输，而unicode仅仅为字符集，由UTF编码来规定unicode字符集如何在
//计算机当中进行存储

//字符编码：
//字符编码即将字符集中的字符映射成其他形式，以便在计算机中进行存储和传输的规则，常见编码规则有ASCII编码和UTF编码等
//UTF-8是一种针对unicode字符集的可变长度字符编码规则，单字节符号，编码方式和ASCII一样

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
	//len获取的是string占用byte的个数，一个字母占1个byte，一个汉字则占用多个byte
	//此方法打印除了str底层占用的每个byte，而一个byte在utf-8当中与ascii当中方式一致，因此对单字母无问题，对于汉字就会有问题
	for i := 0; i < len(str); i++ {
		fmt.Println(reflect.TypeOf(str[i])) //for的valve是uint8，byte是uint8的别名
		fmt.Println(i, str[i])              //依据下标取字符串中的字符，str[i]的类型为byte，因为字符串的底层实际就是[]byte，通过len获取的是字符串的字节长度，类似是从ascii字符集的角度切入
		fmt.Printf("%c\n", str[i])          //%c,打印对应的字符
	}
	//此处len(str)预期应该是8个字符，实际len(str)是12，因为字母acsii字符集占1个byte，而golang
	//使用的是UTF-8编码规范，汉字占3个byte，因此是12
	//UTF-8是针对unicode字符集的一种可变长的字符编码，字符占1个byte，汉字占了3个byte

	//以unicode字符方式遍历
	//range是去依次遍历str字符串当中的每个字符，golang默认使用utf-8编码，utf-8用来编码unicode字符集，因此可以打印出汉字
	for i, ch := range str {
		fmt.Println("range: ", reflect.TypeOf(ch)) //range的value是int32，rune是int32的别名
		fmt.Println(i, string(ch))                 //ch的类型为rune，通过range遍历字符串是从unicode字符集的角度切入
	}
	//range打印出来的是unicode字符集，需要使用string(ch)的方式可以转换为中文

	//以上，[]byte可以和string进行类型互转，因为string是uint8序列构成的，存储的是字符的UTF-8编码，一个汉字占3 byte
	//但是[]rune转string好像直接不行，待查询好的方法-----TODO
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	c := []rune{'你', '好'} //rune是int32，每个汉字占用4个byte呢
	fmt.Println(string(b), string(c))
	for _, ch := range c {
		fmt.Println("[]rune to string: ", string(ch))
	}

	//string转[]rune和[]byte
	d := "你好，我是木头 aaa"
	d1 := []rune(d)
	fmt.Println("string to []rune: ", d1, len(d1))
	e := "hello,i am mutou"
	e1 := []byte(e)
	fmt.Println("string to []byte: ", e1, len(e1))

	//golang提供strconv包来实现基本数据类型与字符串之间的转化
	//strconv.Atoi() 字符串转化为整形
	//strconv.Itoa() 整形转化为字符串

	//只打印索引
	for i := range str {
		fmt.Println(i)
	}

	//空字符串零值是“”，但不是nil，nil也不能赋值给string类型，因此不存在值为nil的string
	//var tmp string
	//tmp = nil
	//fmt.Println(tmp)

	//range出来的value都是int32类型，无关字母还是汉字，UTF-8编码
	str1 := "abcdefg"
	for _, v := range str1 {
		fmt.Println(v)
	}

	//打印中文
	ss := "中国"
	for _, v := range ss {
		fmt.Printf("%c", v)
		fmt.Println(string(v))
	}
}

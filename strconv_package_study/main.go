package main

import (
	"fmt"
	"strconv"
	"time"
)

//strconv 包提供了字符串与整型互转
//Atoi  是ParseInt的便捷版，等价于ParseInt(s , 10 , 0),ParseInt 三个参数含义分别是：待转的字符串、转换的进制、转换后的int范围(int,int8,int32等)
//Itoa  是FormatInt的便捷版，等价于FormatInt(a, 10)

func main() {
	//strconv.ParseInt 字符串按照不同进制转换成十进制，并可以定义目标十进制的int范围
	s, _ := strconv.ParseInt("123", 10, 0)
	fmt.Printf("%d  %T\n", s, s)

	s1, _ := strconv.ParseInt("12390", 10, 8)
	fmt.Printf("%d  %T\n", s1, s1) //127

	s2, _ := strconv.ParseInt("012", 0, 0) //base为0时，根据string前缀来判断按照十进制、八进制、十六进制进行转换
	fmt.Println(s2)

	s3, _ := strconv.ParseInt("12", 8, 0) //base为8时，会将12当做八进制转换成int
	fmt.Println(s3)

	s4, _ := strconv.ParseInt("AA", 16, 0) //base为16时，会将AA当做十六进制转换成int
	fmt.Println(s4)

	s5, _ := strconv.ParseInt("AA", 32, 0) //转换三十二进制
	fmt.Println(s5)

	//strconv.FormatInt 字符串按照给定的进制，将数字转换为对应的字符串字符
	s6 := strconv.FormatInt(128, 10)
	fmt.Printf("%s %T\n", s6, s6)

	s7 := strconv.FormatInt(128, 16)
	fmt.Printf("%s %T\n", s7, s7)

	//对比一下Sprintf转换字符串和FormatInt转换字符串的时间
	starttime := time.Now()
	fmt.Sprintf("%x", 126)
	fmt.Println(time.Now().Sub(starttime))
}

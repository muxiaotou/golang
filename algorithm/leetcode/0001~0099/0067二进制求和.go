//+build ignore
package main

import (
	"fmt"
	"strconv"
)

//解题思路：
//计划倒序开始计算，从最右边开始计算两个二进制数的和

func binariadd(a, b string) (res string) {
	if len(a) < len(b) {
		a, b = b, a //保持a最长，因为下一步要对短的字符串进行补零操作
	}

	for i := len(a) - len(b); i > 0; i-- {
		b = "0" + b //短的string高位补零，+字符串的相加
	}

	var flag int //定义一个变量来作为进位标识

	for i := len(a) - 1; i >= 0; i-- {
		aNum := a[i] - '0'
		bNum := b[i] - '0'
		temp := int(aNum) + int(bNum) + flag
		fmt.Println("temp", temp)
		if temp == 3 {
			res = "1" + res
			flag = 1
		} else if temp == 2 {
			res = "0" + res
			flag = 1
		} else {
			//res = string(temp) + res ，直接使用string转换int达不到预期结果
			//    在Golang中int转string不能直接使用string(int)方式转换，使用fmt.Sprintf()即可。
			//    string(int)中,把int当成UTF-8中二进制码,转换成了对应的字符,汉字,符号等。
			//	  https://blog.csdn.net/qq_31930499/article/details/108185583
			res = strconv.Itoa(temp) + res
			fmt.Println("res:", temp, string(temp), res)
			flag = 0
		}
	}

	if flag == 1 {
		res = "1" + res
	}
	return
}

func main() {
	fmt.Println(binariadd("111", "100"))
}

//字符串IP与数字互转
//判断IP是否合法
//判断是ipv4还是v6
//判断IP是否属于某个网段
//判断电话号码是否合法
package main

//IP特点，每段使用.点好分隔，每段的数字取值时0~255，二进制使用8位即可以表示
import (
	"fmt"
	"strconv"
	"strings"
)

//192.168.1.1，字符串IP转数字
func strtoip(s string) (res int) {
	strlist := strings.Split(s, ".")
	pos := 24 //标记每个IP段数字移动的位置
	for _, v := range strlist {
		i, _ := strconv.Atoi(v)
		i <<= pos     //当前段的ip数字移动到对应位置
		res = res | i //起到一个相加的作用，将IP断的每个断拼接起来，保存到最终结果res当中
		//fmt.Printf("%T\n", i)
		pos -= 8 //下一个断少移动8位
	}
	return
}

//数字转IP
func iptostr(n int) (res string) {
	strlist := make([]string, 4)
	for i := 0; i < len(strlist); i++ {
		strlist[len(strlist)-i-1] = strconv.Itoa(n & 0xff) //每次取数字的低8位，存入slice当中
		n >>= 8                                            //数字右移8位
	}
	for k, v := range strlist {
		res = res + v
		if k < len(strlist)-1 {
			res = res + "."
		}
	}
	return
}
func main() {
	a := 8
	b := 1
	fmt.Println(a | b)
	s := "222.173.108.86"
	//fmt.Printf("%T", s)
	fmt.Println(strtoip(s)) //3735907414

	fmt.Println(iptostr(3735907414)) //222.173.108.86
}

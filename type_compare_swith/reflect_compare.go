package main

import (
	"fmt"
	"reflect"
)

/*
	reflect.DeepEqual函数可以用来比较任意类型的变量的值是否深度一致
	通常用来比较两个slice、struct(包含有不可比较元素)、map等是否相等
*/

func main() {

	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	m2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	/*
		直接比较map，不允许
		但可以使用reflect.DeepEqual进行深度值的比较，如果类型和值相同，则结果为true
	*/
	//fmt.Println(m1 == m2)   invalid operation: m1 == m2 (map can only be compared to nil)
	fmt.Println(reflect.DeepEqual(m1, m2)) //true
	m2 = map[string]int{"foo": 1, "bar": 3}
	fmt.Println(reflect.DeepEqual(m1, m2)) //false

	m3 := map[string]interface{}{
		"foo": []int{1, 2},
	}
	m4 := map[string]interface{}{
		"foo": []int{1, 2},
	}
	fmt.Println(reflect.DeepEqual(m3, m4)) //true

	var m5 map[string]interface{}
	var m6 map[string]interface{}
	fmt.Println(reflect.DeepEqual(m5, nil)) //false, 这个比较特殊哈，两个值必须同为nil或者同为non-nil，因为必须满足类型相同
	fmt.Println(reflect.DeepEqual(m5, m6))  //true
	fmt.Println(m5 == nil)                  //true

	/*
		直接进行slice比较，不被允许
		但可以使用reflect.DeepEqual对slice进行深层次的比较，只有类型和值相同，则深度比较结果就位true
	*/
	s := []string{
		"foo",
	}
	fmt.Println(reflect.DeepEqual(s, []string{"foo"})) //true
	fmt.Println(reflect.DeepEqual(s, []string{"bar"})) //false

	fmt.Println(reflect.DeepEqual(s, nil))                 //false，nil类型和s不相同，两者值是相同的，因此不相等
	fmt.Println(reflect.DeepEqual([]string{}, []string{})) //true
}

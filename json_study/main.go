package main

import (
	"encoding/json"
	"fmt"
)

func struct_to_json() {

	//struct to json string
	type Student1 struct {
		Name string
		Age  int
		//age   int  //大写字母开头会被序列化，小写字母不会被序列化,因为小写字母为私有变量，无法取到反射信息
		Skill string
	}

	stu1 := Student1{"beibei", 3, "football"}
	stu11 := Student1{Name: "mutou", Skill: "basketball"} //struct内未赋值字段，序列号之后为对应零值
	data1, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("marshal fail")
		return
	}
	fmt.Println("marshal success: ", string(data1))

	data11, err := json.Marshal(stu11)
	if err != nil {
		fmt.Println("marshal fail")
		return
	}
	fmt.Println("marshal success: ", string(data11))

	//marshal success:  {"Name":"beibei","Age":3,"Skill":"football"}
	//marshal success:  {"Name":"mutou","Age":0,"Skill":"basketball"}
}

func json_to_struct() {

	type Student1 struct {
		Name  string
		Age   int
		Skill string
	}
	str1 := `{"Name":"beibei","Age":3,"Skill":"football"}`      //反引号字符串的文本不会被解析，双引号字符串的文本会被解析
	str2 := `{"Name":"beibei","Skill":"football"}`              //json string内未对struct 字段赋值，反序列化后字段未自身零值
	str3 := `{"Name":"beibei","Skill":"football","Info":"vip"}` //json string内有字段不在struct时，反序列号时忽略

	var stu1 Student1
	var stu2 Student1
	var stu3 Student1

	err := json.Unmarshal([]byte(str1), &stu1)
	if err != nil {
		fmt.Println("unmarshal err ")
		return
	}
	fmt.Println("unmarshal success, ", stu1)

	err = json.Unmarshal([]byte(str2), &stu2)
	if err != nil {
		fmt.Println("unmarshal err ")
		return
	}
	fmt.Println("unmarshal success, ", stu2)

	err = json.Unmarshal([]byte(str3), &stu3)
	if err != nil {
		fmt.Println("unmarshal err:  ", err)
		return
	}
	fmt.Println("unmarshal success, ", stu3)

	//unmarshal success,  {beibei 3 football}
	//unmarshal success,  {beibei 0 football}
	//unmarshal success,  {beibei 0 football}
}

func tag_define() {

	type Student struct {
		Name string `json:"name"` //加了tag之后，按照tag的名称显示和识别
		Age  int
		Info string `json:"info,omitempty"` //omitempty,当info为空时忽略，而不是显示零值，omit 省略
		//Info string `json:"info"` //当info为空时不会忽略，显示零值,即空字符串
	}

	stu := Student{Name: "mutou", Age: 12}
	str1, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("marshal fail, ", err)
		return
	}
	fmt.Println("marshal success, ", string(str1))

	str2 := `{"name":"chen","Age":18}`
	var stu1 Student
	err = json.Unmarshal([]byte(str2), &stu1)
	if err != nil {
		fmt.Println("unmarshal fail, err is ", err)
		return
	}
	fmt.Println("unmarshal success, ", stu1)

	//marshal success,  {"name":"mutou","Age":12}
	//unmarshal success,  {chen 18}
}

func map_to_json() {
	//map key无大写开头的要求
	m := make(map[string]interface{})
	m["name"] = "chenli"
	m["age"] = 38

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("map marshal fail, ", err)
		return
	}
	fmt.Println("map marsh success , ", string(data))

	m1 := make(map[string]interface{})
	err = json.Unmarshal(data, &m1)
	if err != nil {
		fmt.Println("map unmarshal fail, ", err)
		return
	}
	fmt.Println("map unmarshal success , ", m1)

	//map marsh success ,  {"age":38,"name":"chenli"}
	//unmarshal success ,  map[age:38 name:chenli]
}

func main() {
	struct_to_json()
	json_to_struct()
	tag_define()
	map_to_json()
}

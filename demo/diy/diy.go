package diy

//根据首字母大小写来确定可访问的权限，无论方法名、常量、变量名还是结构体名字，如果首字母大写，则可以被其他包访问，如果首字母小写，则只能在本包内使用
//结构体中属性名首字母小写时无法进行数据解析(如json解析)，所以需要解析时属性名首字母要大写

//定义结构体，大写外部包可以使用
type Diy struct {
	A int64   //大写成员可以导出，即外部包可以使用
	b float64 //小写成员不可以到处，即包内可以使用
}

//引用结构体的方法，引用传递，会改变原有结构体的值
func (d *Diy) Set(a int64, b float64) { //结构体方法无返回值
	d.A = a
	d.b = b
	return
}

//值结构体的方法，值传递，不会改变原有结构体的值
func (d Diy) Set2(a int64, b float64) {
	d.A = a
	d.b = b
	return
}

//首字母小写方法，不能导出被其他包使用
func (d Diy) set(a int64, b float64) {
	d.b = b
	d.A = a
	return
}

//首字母小写函数，不能导出，仅能在同一个包内使用
func sumb(a, b int64) int64 {
	return a + b
}

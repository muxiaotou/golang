package main

//和条件语句一样，Switch坐花括号{必须与switch处于同一行；
//单个case中，可以出现多个结果选项(通过逗号分隔)
//golang不需要使用break来明确退出一个case
//只有在case中明确添加fallthtough关键字，才会继续执行行紧跟的下一个case
//可以不设定switch之后的条件表达式，此种情况下switch结果与if...else...的逻辑作用相同(比如以下方式二)，switch后设定了变量，则case后面的值是需要和这个变量判定是否相等(比如以下方式一)

//switch 有两种方式：

//方式一：
//switch var {
//	case val1:
//		...
//	case val2:
//		...
//	default:
//		...
//}

//方式一可以用逗号来分隔不同的分支条件 case var1,var2:
//case 60:
//	fallthrough

//方式二：
//score := 100
//swith {
//	case score >= 90:
//		fmt.Println("Grade:A")
//	case score >=80 && score < 90:
//		fmt.Println("Grade:B")
//	...
//	default:
//		fmt.Println("Grade:F")
//}

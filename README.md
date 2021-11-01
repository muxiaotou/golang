    only study
    
    整理一些各种文档、书籍的学习笔记
    
    2021年10月16日
    书籍：go专家编程
    进度：
    文档：https://geekr.dev/golang-tutorial
    进度：https://geekr.dev/posts/go-conditional-statement 10月24日

#零值
    array、struct，每个元素或字段都是对应类型的零值
    slice、map、对于零值是nil
    bool: false, int: 0, float: 0.0, string: "", pointer\function\interface\slice\channel\map: nil
#var声明变量
    var 变量名 类型 = 表达式，如果省略了表达式，就使用类型的零值初始化变量
    num := 仅在函数内定义正确，函数外定义失败
    根据首字母大小写来确定可访问的权限，无论方法名、常量、变量名还是结构体名字，
    如果首字母大写，则可以被其他包访问，如果首字母小写，则只能在本包内使用(本包内的不同go文件可以访问)
    结构体中属性名首字母小写时无法进行数据解析(如json解析)，所以需要解析时属性名首字母要大写
    
#驼峰命名法
    首个单词小写，之后每个新单词的首字母大写，如userName，如果希望被外部包使用的全局变量，则需要首字母大写
    
#字符、字符串
    golang当中没有专门的字符类型，一般单个ascii字符用byte保存,单个汉字用rune来保存，此类字符用单引号''包起来
    byte：uint8的别名，用来表示ascii字符
    rune：int32的别名，用来表示汉子字符
    string：底层是一个byte数组
    
    字符串遍历有两种方法，可以看string_study当中的示例代码
    
#值传递
    https://www.cnblogs.com/snowInPluto/p/7477365.html 
    
#defer
    defer在函数执行完毕或者运行抛出panic前执行
    文件的关闭  defer f.Close()
    如果定义了多个defer语句，按照先进后出的顺序执行


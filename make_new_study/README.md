    主要学习一下var & make & new 在声明变量关于array、slice、map、struct这些类型上的区别。
    
#var声明变量
    var 变量名 类型 = 表达式，如果省略了表达式，就使用类型的零值初始化变量
    
#零值
    array、struct，每个元素或字段都是对应类型的零值
    slice、map、对于零值是nil
    bool: false, int: 0, float: 0.0, string: "", pointer\function\interface\slice\channel\map: nil
    
#make和new
    new：func new(Type) *Type
    创建一个Type类型的匿名变量，初始为Type类型的零值，返回变量地址，返回的指针类型为*Type
    make：make(t Type, size ...TntegerType) Type
    分配并初始化一个类型为slice、map或channel的对象，返回类型与Type相同，而非指向它的指针
        
    
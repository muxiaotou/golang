    1.go 语言内置的error接口
    func Foo(param int) (n int, err error) {
        ...
    }
    
    error是一个interface
    type error interface {
        Error() string
    }
    
    2.通过标准错误包errors提供的New()方法创建一个error类型的错误实例
    errors.New("oh ! no ! errors")
    
    
    
    
    问题描述：
        在编译openapi代码时，提示依赖的common包里面import的一个包找不到，对比了此前代码，common包里面import这块没有变动。
    原因定位：
        改动前，依赖的common包里的包import后，且代码当中有使用，但在openapi代码当中没有调用此包的相关变量及方法，改动后
        openapi有调用此包的变量及方法，因此出现错误：
        ppproot@localhost chenli # tree 
        .
        ├── chenli.go
        └── pack
            ├── i.go
            └── li
                └── li.go
        
        cat chenli.go 
            package main
        
            import (
                    "fmt"
                    "beibei.com/chenli/pack"
            )
        
            func main() {
                    fmt.Println("Hello chenli")
                    pack.Ppp1()
            }
        
        cat i.go 
            package pack
            
            import (
                    "fmt"
            )
            
            func Ppp1() {
                    fmt.Print("ppp")
            }
        
        cat li.go 
            package pack
            
            import (
                    "fmt"
                    "beibei.com/parallel_create/common1"
            )
            
            func Ppt() {
                    fmt.Println("pppt")
                    _ := common1.check()
            }
            
            func Ppp() {
                    fmt.Print("ppp")
            }
        
        如上例，chenli.go会import pack，但没有importpack下的li，所以li下面的li.go调用不到，即使common1的包不存在，也不会check
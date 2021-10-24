    整理一下go mod的学习，这篇文章讲的很详细：https://www.jianshu.com/p/760c97ff644c
    
GOPATH
    
    1.GOPATH包管理方式约定$GOPATH目录当中需要有三个子目录：src、pkg、bin，其中src当中存放源代码
    2.所有的包管理依赖手动管理
    3.依赖包没有版本可言
GO MOD
    
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的
    当modules功能启动后，依赖包的存放位置变为$GOPATH/pkg/mod，允许同一个package多个版本共存且多个项目可以共享缓存的module
    
    mkdir gomod
    cd gomod
    go mod init good  初始化项目，生成go.mod文件
    $ cat go.mod
    module gomod
    
    go 1.13
    
    go mode提供了module,require,replace和exclude四个命令
    module：语句指定包的名字(路径)
    requite：语句指定的依赖性模块
    replace：语句可以替换依赖性模块
    exclude：语句可以忽略依赖项模块
    
    如本例子，执行go run main.go，自动查找依赖，并自动下载
    module gomod
    
    go 1.13
    
    require github.com/gin-gonic/gin v1.7.4 // indirect
    
    replace beibei.com/bs/bs-gateway => ../bs/bs-gateway 替换依赖模块为go.mod文件相对路径下面的某个包
    golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a 替换依赖包为一个可以访问的库
    
    https://www.cnblogs.com/kingstarer/p/13461048.html 
    当go命令运行，会先查找当前目录有没有go.mod，如果没有会一直向上找，直到找到go.mod，并以找到go.mod的目录做为根目录进行第三方包引用 。
    
    

    
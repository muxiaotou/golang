    
    https://mp.weixin.qq.com/s/NGgn7E5AJX8hEVPtQ2DaDw

    Go的自动化测试框架要点如下：
    - 测试代码以xxx_test.go方式命名
    - 测试代码中import “testing”
    - 测试函数形如 func Testxyz(t *Testing.T) {…}
    - 执行测试：go test

    go test默认运行所在目录下的xxx_test.go 的测试文件
    
    $ go test
    PASS
    ok      _/D_/zbs-code/src/beibei.com/golang/gotest      0.317s

    $ go test -v 打印详细的测试结果
    === RUN   TestSplit
    --- PASS: TestSplit (0.00s)
    === RUN   TestSplitWithComplexSep
    --- FAIL: TestSplitWithComplexSep (0.00s)
    split_test.go:21: expected: [a d], got: [a cd]
    FAIL
    exit status 1
    FAIL    _/D_/zbs-code/src/beibei.com/golang/gotest      0.328s

    $ go test -run=Sep -v 仅仅运行正则匹配Sep的测试

    测试函数以Test、Benchmark、Example和可选的函数名组成，可选函数名首字母必须大写
    func TestXxx(t *Testing.T) 为测试函数，测试一些逻辑行为是否正确
    func BenchmarkXxx(t *Testing.T) 为基准测试函数，测试函数的性能
    func ExampleXxx(t *Testing.T) 为示例函数，为函数提供示例文档


    
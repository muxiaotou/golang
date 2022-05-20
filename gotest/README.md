    
    Go的自动化测试框架要点如下：
    - 测试代码以xxx_test.go方式命名
    - 测试代码中import “testing”
    - 测试函数形如 func Testxyz(t *Testing.T) {…}
    - 执行测试：go test

    go test默认运行所在目录下的xxx_test.go 的测试文件

    测试函数以Test、Benchmark、Example和可选的函数名组成，可选函数名首字母必须大写
    func TestXxx(t *Testing.T) 为测试函数，测试一些逻辑行为是否正确
    func BenchmarkXxx(t *Testing.T) 为基准测试函数，测试函数的性能
    func ExampleXxx(t *Testing.T) 为示例函数，为函数提供示例文档


    
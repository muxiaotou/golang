    
    https://mp.weixin.qq.com/s/NGgn7E5AJX8hEVPtQ2DaDw

# go test
    Go的自动化测试框架要点如下：
    - 测试代码以xxx_test.go方式命名
    - 测试代码中import “testing”
    - 测试函数形如 func Testxyz(t *Testing.T) {…}
    - 执行测试：go test

    go test默认运行所在目录下所有的xxx_test.go 的测试文件，这些test.go文件没有也不需要main()作为函数入口
    
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
    $ go test -short ，当Test测试用例当中有short的判断
        if testing.Short() {
		    t.Skip("short model, skip test")
	    }
    时，跳过后续的测试代码

    子测试，在Test测试用例内执行一组测试用例，这样就不需要为不同的测试数据定义多个测试函数了
    t.Run("case1",func(t *testing.T){...})
    可以在上面子测试的匿名函数内t.Parallel(),多组自测试并行运行
    表格测试方法：多组测试用例时，可以定义一个struct的slice作为测试用例，然后遍历表格，使用t.Run对每个条目执行必要的测试

    测试函数以Test、Benchmark、Example和可选的函数名组成，可选函数名首字母必须大写
    func TestXxx(t *testing.T) 为测试函数，测试一些逻辑行为是否正确
    func BenchmarkXxx(t *testing.T) 为基准测试函数，测试函数的性能
    func ExampleXxx(t *testing.T) 为示例函数，为函数提供示例文档

    t *Testing.T,t参数用于报告测试失败和附加的日志信息。
    当一个测试的测试函数返回时，又或当一个测试函数调用FailNow、Fatal、Fatalf、SkipNow、Skip或者Skipf中的任意一个时，该测试宣告结束
    t.Log 打印日志
    t.Logf 格式化打印日志
    t.Error 相当于调用Log之后调用Fail，标记失败，但仍继续执行该测试
    t.Errorf 相当于调用Logf之后调用Fail，标记失败，但仍继续执行该测试
    t.Fail 标记当前测试为失败，但是仍继续执行该测试
    t.FailNow 标记当前测试失败，并停止执行该测试，在此之后，测试过程将在下一个测试中继续,即停止当前Test的测试，其他Test的测试继续
    t.Fatal 相当于调用Log之后调用FailNow
    t.Fatalf 相当于调用Logf之后调用FailNow
    
    

# 代码覆盖率 go test -cover
    go test -cover 输出单元测试用例对代码的覆盖率  ， coverage: 87.5% of statements
    go test -cover -v -coverprofile=c.out  将覆盖率信息输出到文件，使用go tool cover -html=c.out来处理生成的覆盖率信息
    
# testify 
    testing包没有提供断言，因此使用testing包时会写出许多的if...else语句，testify/assert、testify/require包提供了许多断言函数

    go get github.com/stretchr/testify

    import "github.com/stretchr/testify/assert"
    import "github.com/stretchr/testify/require"

    assert.Equal(t, want, got)
    require.Equal(t, want, got)
    上述写法需要传入testing.T类型变量t，可以创建assert对象，assert := assert.New(t),assert对应因此拥有了所有断言防范，
    后续不再需要传入testing.T参数了。
    assert := assert.New(t)
	assert.Equal(want, got, "they should be equal")

    require拥有和assert包所有的断言函数，唯一区别是require遇到失败的用例就立即终止本次测试，而assert仅仅只是标记失败
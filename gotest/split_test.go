package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

//没有也不需要main()作为函数入口，以Test开头的函数默认都会被自动执行

func TestSplit(t *testing.T) {
	//t.Parallel() //各Test用例之间并行，因此每次go  test执行的最终结果顺序不一定一致
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}

	t.Log("test t.Log")
	//t.Error("test t.Error")
	//t.FailNow()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
	fmt.Println("t.Log")
}

//go test -run Sep  仅仅运行正则匹配Sep的Test用例
func TestSplitWithComplexSep(t *testing.T) {
	//t.Parallel()
	got := Split("abcd", "bc")
	//want := []string{"a", "c"}
	want := []string{"a", "d"}

	//testing未提供断言，需要自己if...else...来进行判断
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
	//使用testify/assert包
	//assert.Equal(t, want, got)
	assert := assert.New(t)
	assert.Equal(want, got, "they should be equal")
	t.Log("assert equal test")

	//使用testify/require包
	require.Equal(t, want, got)
	t.Log("require equal test")
}

//go test -short 由于testing.Short()判断为真，会从t.Skip return返回，从而跳过后续的Println代码
func TestSplitSkip(t *testing.T) {
	//t.Parallel()
	if testing.Short() {
		t.Skip("short model, skip test")
	}
	fmt.Println("test skip")
}

//子测试 t.Run ，并行运行t.Parallel, 以下定义一个slice of struct，表格驱动测试的方式
func TestSplitAll(t *testing.T) {
	//t.Parallel()
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() //多次执行，多个自测试每次执行顺序不一定相同
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expect: %v, got: %v", tt.want, got)
			}
		})
	}
}

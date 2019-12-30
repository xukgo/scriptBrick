package unitTest

import (
	"bytes"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strconv"
	"testing"
)

func BenchmarkFuncDefine1(b *testing.B) {
	funcMap := make(map[string]funcField.IScriptStringMinor)
	funcMap["sum"] = new(SumMinor)
	funcMap["stringJoin"] = new(StringJoinMinor)
	exp := "stringJoin(abc,sum(100.789,30.56),def100)"

	rootDefine, err := funcField.ParseStringFuncDefine(exp)
	if err != nil {
		b.Fail()
	}
	err = rootDefine.InitFunc(funcMap)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		val, err := rootDefine.Excute(nil)
		if err != nil {
			b.Fail()
		}
		if val != "abc131.349def100" {
			b.Fail()
		}
	}
}

type SumMinor struct {
}

func (this *SumMinor) EvalScript(ctx interface{}, args ...string) (interface{}, error) {
	var sum float64
	for idx := range args {
		v, err := strconv.ParseFloat(args[idx], 64)
		if err != nil {
			return nil, err
		}
		sum += v
	}
	return sum, nil
}

//每个参数都是数字
//func (this *SumMinor) CheckParams(args ...string) error {
//
//}

type StringJoinMinor struct {
}

func (this *StringJoinMinor) EvalScript(ctx interface{}, args ...string) (interface{}, error) {
	bf := new(bytes.Buffer)
	for idx := range args {
		bf.WriteString(args[idx])
	}
	return bf.String(), nil
}

//func (this *StringJoinMinor) CheckParams(args ...string) error {
//
//}

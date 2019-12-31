package expressionField

import (
	"fmt"
	"github.com/xukgo/scriptBrick/mathEngine"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strconv"
	"testing"
)

func BenchmarkCalcObject1(b *testing.B) {
	funcMap := make(map[string]funcField.IScriptObjectMinor)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(CalcFuncMinor)
	exp := "sum(calc(1000+222.12+(100*2)),sum(1,2))"
	calcMino := NewCalcObjectMinor(funcMap)
	err := calcMino.Init(exp)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		val, err := calcMino.Calc(nil)
		if err != nil {
			b.Fail()
		}
		if val != 1425.12 {
			b.Fail()
		}

	}
}

type SumObjectMinor struct {
}

func (this *SumObjectMinor) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	var sum float64
	for idx := range args {
		str := fmt.Sprintf("%v", args[idx])
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		sum += v
	}
	return sum, nil
}

func (this *SumObjectMinor) CheckArgCount(count int) bool {
	return count > 0
}

type CalcFuncMinor struct {
}

func (this *CalcFuncMinor) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	str := fmt.Sprintf("%v", args[0])
	res,err := mathEngine.ParseAndExec(str)
	return res,err
}

func (this *CalcFuncMinor) CheckArgCount(count int) bool {
	return count == 1
}

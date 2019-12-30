package expressionField

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strconv"
	"testing"
)

func BenchmarkCalcObject1(b *testing.B) {
	funcMap := make(map[string]funcField.IScriptObjectMinor)
	funcMap["sum"] = new(SumObjectMinor)
	exp := "1000+sum(100.23,100.11,sum(1,2))+(100*2)"
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
		if val != 1403.34 {
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

func (this *SumObjectMinor) CheckArgValid(ctx interface{}, args ...interface{}) error {
	if len(args) == 0{
		return fmt.Errorf("args count cannot be 0")
	}
	return nil
}

func (this *SumObjectMinor) CheckArgCount(count int) bool {
	return count > 0
}

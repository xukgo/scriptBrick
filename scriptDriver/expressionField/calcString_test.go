package expressionField

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strconv"
	"testing"
)

func BenchmarkCalcString1(b *testing.B) {
	funcMap := make(map[string]funcField.IScriptStringMinor)
	funcMap["sum"] = new(SumStringMinor)
	exp := "1000+sum(100.23,100.11,sum(1,2))+(100*2)"
	calcMino := NewCalcStringMinor(funcMap)
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

type SumStringMinor struct {
}

func (this *SumStringMinor) EvalScript(ctx interface{}, args ...string) (interface{}, error) {
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

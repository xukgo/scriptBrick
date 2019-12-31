package unitTest

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"github.com/xukgo/scriptBrick/scriptDriver/brick"
	"math"
	"strconv"
	"testing"
)

func BenchmarkCalcObject1(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptMinor)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brick.CalcExpBrick)
	exp := "sum(sum(3,2),sum(1,2.12))"
	calcMino := scriptDriver.NewFuncExecNode(funcMap)
	err := calcMino.Init(exp)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		val, err := calcMino.Calc(nil)
		if err != nil {
			b.Fail()
		}
		gap := math.Abs(float64(val.(float64)) - 8.12)
		if gap > 0.01 {
			b.Fail()
		}

	}
}

func BenchmarkCalcObject2(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptMinor)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brick.CalcExpBrick)
	exp := "sum(sum(3,2),calc(1000+sum(1,2)+222.12-(100*2)),sum(1,2))"
	calcMino := scriptDriver.NewFuncExecNode(funcMap)
	err := calcMino.Init(exp)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		val, err := calcMino.Calc(nil)
		if err != nil {
			b.Fail()
		}
		if val != 1033.12 {
			b.Fail()
		}

	}
}

func BenchmarkCalcObject3(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptMinor)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brick.CalcExpBrick)
	exp := "sum(calc(1000+sum(1,2)+calc(11+22*1)),sum(1,2))"
	calcMino := scriptDriver.NewFuncExecNode(funcMap)
	err := calcMino.Init(exp)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		val, err := calcMino.Calc(nil)
		if err != nil {
			b.Fail()
		}
		if val != 1039.00 {
			b.Fail()
		}

	}
}

type SumObjectMinor struct {
}

func (this *SumObjectMinor) CloneBasic() scriptDriver.IScriptMinor {
	return this
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

func (this *SumObjectMinor) GetIsExpressionArg(int) bool {
	return false
}

func (this *SumObjectMinor) AfterInitCorrectArg(dict map[string]scriptDriver.IScriptMinor, index int, arg *scriptDriver.FuncNodeArg) error {
	return nil
}

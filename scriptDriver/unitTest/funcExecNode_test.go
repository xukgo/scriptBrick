package unitTest

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"github.com/xukgo/scriptBrick/scriptDriver/brickBomb"
	"math"
	"strconv"
	"testing"
)

func BenchmarkCalcObject1(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "sum(sum(3,2),sum(1,2.12))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, nil)
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		gap := math.Abs(float64(val.(float64)) - 8.12)
		if gap > 0.01 {
			b.Fail()
			return
		}

	}
}

func BenchmarkCalcObject2(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "sum(sum(3,2),calc('1000+sum(1,2)+222.12-(100*2)'),sum(1,2))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, nil)
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		if val != 1033.12 {
			b.Fail()
			return
		}

	}
}

func BenchmarkCalcObject3(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "sum(calc('1000+sum(1,2)+calc(11+22*1)'),sum(1,2))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, nil)
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		if val != 1039.00 {
			b.Fail()
			return
		}

	}
}

func BenchmarkCalcObject4(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "calc('1.1+2.2').sum(sum(1,2)).sum(10,20).sum(calc('100+200'))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, nil)
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		if val != 336.3 {
			b.Fail()
			return
		}

	}
}

func BenchmarkCalcObject5(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "calc('1.1+sum(sum(1,2)).sum(10,20)+2.2').sum(calc('100+200'))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, nil)
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		if val != 336.3 {
			b.Fail()
			return
		}

	}
}

func BenchmarkCalcObject6(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["calc"] = new(brickBomb.CalcExpBrick)
	exp := "preEval(calc('1.1+sum(sum(1,2).sum(10,20))+2.2').sum(calc('100+200')))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp, scriptDriver.NewPreConstParam(nil, "preEval"))
	if err != nil {
		b.Fail()
		return
	}

	for i := 0; i < b.N; i++ {
		val, err := brick.Build(nil)
		if err != nil {
			b.Fail()
			return
		}
		if val != 336.3 {
			b.Fail()
			return
		}

	}
}

type SumObjectMinor struct {
}

func (this *SumObjectMinor) CloneBasic() scriptDriver.IScriptBrick {
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

//func (this *SumObjectMinor) AfterInitCorrectArg(dict map[string]scriptDriver.IScriptBrick, index int, arg *scriptDriver.BrickArg) error {
//	return nil
//}

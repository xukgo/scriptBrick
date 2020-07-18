package unitTest

import (
	"github.com/xukgo/scriptBrick/scriptDriver"
	"github.com/xukgo/scriptBrick/scriptDriver/brickBomb"
	"math"
	"testing"
)

func Test_constBuild(t *testing.T) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["add"] = new(AddObjectMinor)
	var exps []string
	exps = append(exps, "add(sum(3,2),add(1,2.12,sum(1,1)),sum(1,2))")
	exps = append(exps, "sum(sum(3,2),add(1,2.12,sum(1,1)),sum(1,2))")
	exps = append(exps, "add(sum(3,2),add(1,2.12,add(1,1)),sum(1,2))")
	exps = append(exps, "sum(sum(3,2),sum(1,2.12,add(1,1)),sum(1,2))")

	for _, exp := range exps {
		brick, err := scriptDriver.CreateBrick(funcMap, exp)
		if err != nil {
			t.FailNow()
			return
		}

		val, err := brick.Build(nil)
		if err != nil {
			t.FailNow()
			return
		}
		gap := math.Abs(float64(val.(float64)) - 13.12)
		if gap > 0.01 {
			t.FailNow()
			return
		}
	}
}
func BenchmarkCalcObject1(b *testing.B) {
	funcMap := make(map[string]scriptDriver.IScriptBrick)
	funcMap["preEval"] = new(scriptDriver.PreBuildBrick)
	funcMap["sum"] = new(SumObjectMinor)
	funcMap["add"] = new(AddObjectMinor)
	exp := "sum(sum(3,2),sum(1,2.12),add(1,2))"
	brick, err := scriptDriver.CreateBrick(funcMap, exp)
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
		gap := math.Abs(float64(val.(float64)) - 11.12)
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
	brick, err := scriptDriver.CreateBrick(funcMap, exp)
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
	brick, err := scriptDriver.CreateBrick(funcMap, exp)
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
	brick, err := scriptDriver.CreateBrick(funcMap, exp)
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
	brick, err := scriptDriver.CreateBrick(funcMap, exp)
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
	brick, err := scriptDriver.CreateBrick(funcMap, exp) //, scriptDriver.NewPreConstParam(nil, "preEval")
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

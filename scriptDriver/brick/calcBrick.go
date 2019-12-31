package brick

import "github.com/xukgo/scriptBrick/scriptDriver"

type CalcExpBrick struct {
	funcDot *CalcFuncDot
}

func (this *CalcExpBrick) Clone() scriptDriver.IScriptMinor {
	return new(CalcExpBrick)
}
func (this *CalcExpBrick) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	return this.funcDot.Calc(ctx)
}

func (this *CalcExpBrick) CheckArgCount(count int) bool {
	return count == 1
}

func (this *CalcExpBrick) GetIsExpressionArg(index int) bool {
	return true
}

func (this *CalcExpBrick) AfterInitCorrectArg(dict map[string]scriptDriver.IScriptMinor, index int, funcArg *scriptDriver.FuncNodeArg) error {
	checkArgExpFuncDict := make(map[string]scriptDriver.CheckExpressionArgFunc)
	for key, val := range dict {
		checkArgExpFuncDict[key] = val.GetIsExpressionArg
	}

	funcDot := new(CalcFuncDot)
	err := funcDot.Init(dict, funcArg.Content)
	if err != nil {
		return err
	}
	this.funcDot = funcDot

	//funcArg.Content = ""
	//funcArg.MType = funcField.TYPE_FUNC
	//funcArg.Func = rootDefine
	return nil
}

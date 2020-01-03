package brick

import "github.com/xukgo/scriptBrick/scriptDriver"

type CalcExpBrick struct {
	calcNode *CalcBrickNode
}

func (this *CalcExpBrick) CloneBasic() scriptDriver.IScriptBrick {
	return new(CalcExpBrick)
}
func (this *CalcExpBrick) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	return this.calcNode.Calc(ctx)
}

func (this *CalcExpBrick) CheckArgCount(count int) bool {
	return count == 1
}

func (this *CalcExpBrick) AfterInitCorrectArg(dict map[string]scriptDriver.IScriptBrick, index int, funcArg *scriptDriver.BrickArg) error {
	calcDot := new(CalcBrickNode)
	err := calcDot.Init(dict, funcArg.Content)
	if err != nil {
		return err
	}
	this.calcNode = calcDot

	//funcArg.Content = ""
	//funcArg.MType = funcField.TYPE_FUNC
	//funcArg.Func = rootDefine
	return nil
}

package brick

import (
	"fmt"
	"github.com/xukgo/scriptBrick/mathEngine"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"strings"
)

type CalcBrickNode struct {
	expression string
	bricks     []*scriptDriver.Brick
}

func (this *CalcBrickNode) Init(dict map[string]scriptDriver.IScriptBrick, exp string) error {
	exp = strings.TrimSpace(exp)
	this.expression = exp
	this.bricks = nil

	msegs, err := scriptDriver.GetFuncSentences(exp)
	if err != nil {
		return err
	}

	checkArgExpFuncDict := make(map[string]scriptDriver.CheckExpressionArgFunc)
	for key, val := range dict {
		checkArgExpFuncDict[key] = val.GetIsExpressionArg
	}

	for idx := range msegs {
		funseg := msegs[idx]
		brick, err := scriptDriver.ParseBrick(funseg, checkArgExpFuncDict)
		if err != nil {
			return err
		}

		err = scriptDriver.InitFuncDefine(dict, brick)
		if err != nil {
			return err
		}

		this.bricks = append(this.bricks, brick)
		exp = strings.ReplaceAll(exp, funseg, formatSegIndex(idx))
	}
	this.expression = exp

	for idx := range this.bricks {
		err := scriptDriver.CheckMinorArgCountValid(this.bricks[idx])
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *CalcBrickNode) Calc(ctx interface{}) (interface{}, error) {
	if len(this.bricks) == 0 {
		return mathEngine.ParseAndExec(this.expression)
	}

	exp := this.expression
	for idx := range this.bricks {
		v, err := this.bricks[idx].Build(ctx)
		if err != nil {
			return nil, err
		}
		exp = strings.ReplaceAll(exp, formatSegIndex(idx), fmt.Sprintf("%v", v))
	}
	return mathEngine.ParseAndExec(exp)
}

func formatSegIndex(index int) string {
	return fmt.Sprintf("func[%02d]", index)
}

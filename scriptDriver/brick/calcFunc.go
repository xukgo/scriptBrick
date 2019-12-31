package brick

import (
	"fmt"
	"github.com/xukgo/scriptBrick/mathEngine"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"strings"
)

type CalcFuncDot struct {
	expression string
	funMinors  []*scriptDriver.FuncNodeMinor
}

func (this *CalcFuncDot) Init(dict map[string]scriptDriver.IScriptMinor, exp string) error {
	exp = strings.TrimSpace(exp)
	this.expression = exp
	this.funMinors = nil

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
		fdefine, err := scriptDriver.ParseFuncDefine(funseg, checkArgExpFuncDict)
		if err != nil {
			return err
		}

		err = scriptDriver.InitFuncDefine(dict, fdefine)
		if err != nil {
			return err
		}

		this.funMinors = append(this.funMinors, fdefine)
		exp = strings.ReplaceAll(exp, funseg, formatSegIndex(idx))
	}
	this.expression = exp

	for idx := range this.funMinors {
		err := scriptDriver.CheckMinorArgCountValid(this.funMinors[idx])
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *CalcFuncDot) Calc(ctx interface{}) (interface{}, error) {
	if len(this.funMinors) == 0 {
		return mathEngine.ParseAndExec(this.expression)
	}

	exp := this.expression
	for idx := range this.funMinors {
		v, err := this.funMinors[idx].Excute(ctx)
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

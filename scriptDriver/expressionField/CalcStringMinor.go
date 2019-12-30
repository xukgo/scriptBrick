package expressionField

import (
	"fmt"
	"github.com/xukgo/scriptBrick/mathEngine"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strings"
)

type CalcStringMinor struct {
	dict      map[string]funcField.IScriptStringMinor
	exp       string
	funMinors []*funcField.FuncStringMinor
}

func NewCalcStringMinor(dict map[string]funcField.IScriptStringMinor) *CalcStringMinor {
	model := new(CalcStringMinor)
	if dict == nil {
		dict = make(map[string]funcField.IScriptStringMinor)
	}
	model.dict = dict
	return model
}

func (this *CalcStringMinor) Init(exp string) error {
	exp = strings.TrimSpace(exp)
	this.exp = exp
	this.funMinors = nil

	msegs, err := funcField.GetFuncSentences(exp)
	if err != nil {
		return err
	}
	for idx := range msegs {
		funseg := msegs[idx]
		rootDefine, err := funcField.ParseStringFuncDefine(funseg)
		if err != nil {
			return err
		}
		err = rootDefine.InitFunc(this.dict)
		if err != nil {
			return err
		}

		this.funMinors = append(this.funMinors, rootDefine)
		exp = strings.ReplaceAll(exp, funseg, formatSegIndex(idx))
	}
	this.exp = exp

	for idx := range this.funMinors{
		err := checkStringMinorArgCountValid(this.funMinors[idx])
		if err != nil{
			return err
		}
	}
	return nil
}

func checkStringMinorArgCountValid(item *funcField.FuncStringMinor)error {
	if !item.RealFuncMinor.CheckArgCount(len(item.FuncArgs)){
		return fmt.Errorf("%s func args count is not valid, count is %d",item.FuncName, len(item.FuncArgs))
	}

	for idx := range item.FuncArgs{
		if item.FuncArgs[idx].MType != funcField.TYPE_FUNC{
			continue
		}

		err := checkStringMinorArgCountValid(item.FuncArgs[idx].Func)
		if err != nil{
			return err
		}
	}

	return nil
}

func (this *CalcStringMinor) Calc(ctx interface{}) (interface{}, error) {
	if len(this.funMinors) == 0 {
		return mathEngine.ParseAndExec(this.exp)
	}

	exp := this.exp
	for idx := range this.funMinors {
		v, err := this.funMinors[idx].Excute(ctx)
		if err != nil {
			return nil, err
		}
		exp = strings.ReplaceAll(exp, formatSegIndex(idx), fmt.Sprintf("%v", v))
	}
	return mathEngine.ParseAndExec(exp)
}

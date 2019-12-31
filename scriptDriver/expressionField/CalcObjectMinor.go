package expressionField

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver/funcField"
	"strings"
)

type CalcObjectMinor struct {
	dict      map[string]funcField.IScriptObjectMinor
	funMinor *funcField.FuncObjectMinor
}

func NewCalcObjectMinor(dict map[string]funcField.IScriptObjectMinor) *CalcObjectMinor {
	model := new(CalcObjectMinor)
	if dict == nil {
		dict = make(map[string]funcField.IScriptObjectMinor)
	}
	model.dict = dict
	return model
}

func (this *CalcObjectMinor) Init(exp string) error {
	exp = strings.TrimSpace(exp)

	rootDefine, err := funcField.ParseObjectFuncDefine(exp)
	if err != nil {
		return err
	}
	err = rootDefine.InitFunc(this.dict)
	if err != nil {
		return err
	}

	this.funMinor = rootDefine
	err = checkObjectMinorArgCountValid(this.funMinor)
	if err != nil{
		return err
	}

	return nil
}

func checkObjectMinorArgCountValid(item *funcField.FuncObjectMinor)error {
	if !item.RealFuncMinor.CheckArgCount(len(item.FuncArgs)){
		return fmt.Errorf("%s func args count is not valid, count is %d",item.FuncName, len(item.FuncArgs))
	}

	for idx := range item.FuncArgs{
		if item.FuncArgs[idx].MType != funcField.TYPE_FUNC{
			continue
		}

		err := checkObjectMinorArgCountValid(item.FuncArgs[idx].Func)
		if err != nil{
			return err
		}
	}

	return nil
}

func (this *CalcObjectMinor) Calc(ctx interface{}) (interface{}, error) {
	var err error
	var v interface{}
	v, err = this.funMinor.Excute(ctx)
	if err != nil {
		return nil, err
	}
	return v,nil
}

//func formatSegIndex(index int) string {
//	return fmt.Sprintf("func[%02d]", index)
//}

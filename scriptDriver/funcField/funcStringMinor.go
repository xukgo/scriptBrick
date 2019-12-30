package funcField

import (
	"fmt"
	"strings"
)

type FuncStringMinor struct {
	FuncName      string
	FuncArgs      []*FuncStringArg
	RealFuncMinor IScriptStringMinor
}

func NewFuncStringMinor(fname string, fargs []*FuncStringArg) *FuncStringMinor {
	model := new(FuncStringMinor)
	model.FuncName = fname
	model.FuncArgs = fargs
	return model
}

type FuncStringArg struct {
	MType   int
	Content string
	Func    *FuncStringMinor
}

func NewFuncStringArg(mtype int, indata interface{}) *FuncStringArg {
	model := new(FuncStringArg)
	model.MType = mtype
	if mtype == TYPE_FUNC {
		model.Func = indata.(*FuncStringMinor)
	} else {
		model.Content = indata.(string)
	}
	return model
}

func (this *FuncStringMinor) Excute(ctx interface{}) (interface{}, error) {
	var arr []string
	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			f := this.FuncArgs[idx].Func
			v, err := f.Excute(ctx)
			if err != nil {
				return nil, err
			}
			arr = append(arr, fmt.Sprintf("%v", v))
		} else {
			arr = append(arr, this.FuncArgs[idx].Content)
		}
	}
	return this.RealFuncMinor.EvalScript(ctx, arr...)
}

func (this *FuncStringMinor) InitFunc(factory map[string]IScriptStringMinor) error {
	lcFactory := make(map[string]IScriptStringMinor)
	for k, v := range factory {
		lcFactory[strings.ToLower(k)] = v
	}

	return this.initFuncArg(lcFactory)
}
func (this *FuncStringMinor) initFuncArg(factory map[string]IScriptStringMinor) error {
	if len(this.FuncName) == 0 {
		return fmt.Errorf("func name is empty")
	}
	f, find := factory[strings.ToLower(this.FuncName)]
	if !find {
		return fmt.Errorf("func name is not found instance:%s", this.FuncName)
	}
	this.RealFuncMinor = f

	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			err := this.FuncArgs[idx].Func.initFuncArg(factory)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

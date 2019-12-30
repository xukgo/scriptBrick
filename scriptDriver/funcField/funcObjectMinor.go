package funcField

import (
	"fmt"
	"strings"
)

type FuncObjectMinor struct {
	FuncName      string
	FuncArgs      []*FuncObjectArg
	RealFuncMinor IScriptObjectMinor
}

func NewFuncObjectMinor(fname string, fargs []*FuncObjectArg) *FuncObjectMinor {
	model := new(FuncObjectMinor)
	model.FuncName = fname
	model.FuncArgs = fargs
	return model
}

type FuncObjectArg struct {
	MType   int
	Content string
	Func    *FuncObjectMinor
}

func NewFuncObjectArg(mtype int, indata interface{}) *FuncObjectArg {
	model := new(FuncObjectArg)
	model.MType = mtype
	if mtype == TYPE_FUNC {
		model.Func = indata.(*FuncObjectMinor)
	} else {
		model.Content = indata.(string)
	}
	return model
}

func (this *FuncObjectMinor) Excute(ctx interface{}) (interface{}, error) {
	var arr []interface{}
	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			f := this.FuncArgs[idx].Func
			v, err := f.Excute(ctx)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		} else {
			arr = append(arr, this.FuncArgs[idx].Content)
		}
	}
	return this.RealFuncMinor.EvalInstance(ctx, arr...)
}

func (this *FuncObjectMinor) InitFunc(factory map[string]IScriptObjectMinor) error {
	lcFactory := make(map[string]IScriptObjectMinor)
	for k, v := range factory {
		lcFactory[strings.ToLower(k)] = v
	}

	return this.initFuncArg(lcFactory)
}
func (this *FuncObjectMinor) initFuncArg(factory map[string]IScriptObjectMinor) error {
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

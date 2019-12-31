package scriptDriver

import (
	"fmt"
	"strings"
)

type FuncNodeMinor struct {
	FuncName      string
	FuncArgs      []*FuncNodeArg
	RealFuncMinor IScriptMinor
}

func NewFuncNodeMinor(fname string, fargs []*FuncNodeArg) *FuncNodeMinor {
	model := new(FuncNodeMinor)
	model.FuncName = fname
	model.FuncArgs = fargs
	return model
}

type FuncNodeArg struct {
	MType   int
	Content string
	Func    *FuncNodeMinor
}

func NewFuncNodeArg(mtype int, indata interface{}) *FuncNodeArg {
	model := new(FuncNodeArg)
	model.MType = mtype
	if mtype == TYPE_FUNC {
		model.Func = indata.(*FuncNodeMinor)
	} else {
		model.Content = indata.(string)
	}
	return model
}

func (this *FuncNodeMinor) Excute(ctx interface{}) (interface{}, error) {
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
	return this.RealFuncMinor.Eval(ctx, arr...)
}

func (this *FuncNodeMinor) InitFunc(dict map[string]IScriptMinor) error {
	lcFactory := make(map[string]IScriptMinor)
	for k, v := range dict {
		lcFactory[strings.ToLower(k)] = v
	}

	return this.initFuncArg(lcFactory)
}
func (this *FuncNodeMinor) initFuncArg(dict map[string]IScriptMinor) error {
	if len(this.FuncName) == 0 {
		return fmt.Errorf("func name is empty")
	}
	f, find := dict[strings.ToLower(this.FuncName)]
	if !find {
		return fmt.Errorf("func name is not found instance:%s", this.FuncName)
	}
	this.RealFuncMinor = f.Clone()

	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			err := this.FuncArgs[idx].Func.initFuncArg(dict)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

package scriptDriver

import (
	"fmt"
	"strings"
)

type Brick struct {
	FuncName      string
	FuncArgs      []*BrickArg
	RealFuncMinor IScriptBrick
}

func (this *Brick) Build(ctx interface{}) (interface{}, error) {
	var arr []interface{}
	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			f := this.FuncArgs[idx].Func
			v, err := f.Build(ctx)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		} else if this.FuncArgs[idx].MType == TYPE_STRING {
			arr = append(arr, this.FuncArgs[idx].Content)
		} else {
			arr = append(arr, this.FuncArgs[idx].Value)
		}
	}
	return this.RealFuncMinor.Eval(ctx, arr...)
}

func (this *Brick) InitFunc(dict map[string]IScriptBrick) error {
	lcFactory := make(map[string]IScriptBrick)
	for k, v := range dict {
		lcFactory[strings.ToLower(k)] = v
	}

	return this.initBrickArg(lcFactory)
}
func (this *Brick) initBrickArg(dict map[string]IScriptBrick) error {
	if len(this.FuncName) == 0 {
		return fmt.Errorf("func name is empty")
	}
	f, find := dict[strings.ToLower(this.FuncName)]
	if !find {
		return fmt.Errorf("func name is not found instance:%s", this.FuncName)
	}
	this.RealFuncMinor = f.CloneBasic()

	for idx := range this.FuncArgs {
		if this.FuncArgs[idx].MType == TYPE_FUNC {
			err := this.FuncArgs[idx].Func.initBrickArg(dict)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

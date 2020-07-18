package unitTest

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"strconv"
)

type AddObjectMinor struct {
}

func (this *AddObjectMinor) CloneBasic() scriptDriver.IScriptBrick {
	return this
}

func (this *AddObjectMinor) SurplusContext() bool {
	return false
}

func (this *AddObjectMinor) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	var sum float64
	for idx := range args {
		str := fmt.Sprintf("%v", args[idx])
		v, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		sum += v
	}
	return sum, nil
}

func (this *AddObjectMinor) CheckArgCount(count int) bool {
	return count > 0
}

//func (this *AddObjectMinor) AfterInitCorrectArg(dict map[string]scriptDriver.IScriptBrick, index int, arg *scriptDriver.BrickArg) error {
//	return nil
//}

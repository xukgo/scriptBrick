package scriptDriver

import (
	"fmt"
)

func InitFuncDefine(dict map[string]IScriptMinor, fdefine *FuncNodeMinor) error {
	var err error
	err = fdefine.InitFunc(dict)
	if err != nil {
		return err
	}

	err = correctArg(dict, fdefine)
	if err != nil {
		return err
	}

	err = CheckMinorArgCountValid(fdefine)
	if err != nil {
		return err
	}

	return nil
}

func correctArg(dict map[string]IScriptMinor, item *FuncNodeMinor) error {
	if item.RealFuncMinor == nil {
		return nil
	}

	for idx := range item.FuncArgs {
		err := item.RealFuncMinor.AfterInitCorrectArg(dict, idx, item.FuncArgs[idx])
		if err != nil {
			return err
		}
	}

	for _, funcArg := range item.FuncArgs {
		if funcArg.MType == TYPE_STRING {
			continue
		}
		if funcArg.Func == nil {
			continue
		}
		err := correctArg(dict, funcArg.Func)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckMinorArgCountValid(item *FuncNodeMinor) error {
	if !item.RealFuncMinor.CheckArgCount(len(item.FuncArgs)) {
		return fmt.Errorf("%s func args count is not valid, count is %d", item.FuncName, len(item.FuncArgs))
	}

	for idx := range item.FuncArgs {
		if item.FuncArgs[idx].MType != TYPE_FUNC {
			continue
		}

		err := CheckMinorArgCountValid(item.FuncArgs[idx].Func)
		if err != nil {
			return err
		}
	}

	return nil
}

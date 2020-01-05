package scriptDriver

import (
	"fmt"
)

func InitFuncDefine(dict map[string]IScriptBrick, brick *Brick) error {
	var err error
	err = brick.InitFunc(dict)
	if err != nil {
		return err
	}

	err = correctArg(dict, brick)
	if err != nil {
		return err
	}

	err = CheckMinorArgCountValid(brick)
	if err != nil {
		return err
	}

	return nil
}

func correctArg(dict map[string]IScriptBrick, item *Brick) error {
	if item.RealFuncMinor == nil {
		return nil
	}

	for idx := range item.FuncArgs {
		mountCallback, ok := item.RealFuncMinor.(IBrickMountCallback)
		if !ok {
			continue
		}
		err := mountCallback.AfterInitCorrectArg(dict, idx, item.FuncArgs[idx])
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

func CheckMinorArgCountValid(item *Brick) error {
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

func checkIsQuotMarkSentence(str string) bool {
	if len(str) < 2 {
		return false
	}
	head := str[0]
	tail := str[len(str)-1]
	if head != tail {
		return false
	}

	if head == '"' {
		return true
	}
	if head == '\'' {
		return true
	}

	return false
}

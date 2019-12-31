package scriptDriver

import (
	"fmt"
	"strings"
)

func CreateBrick(dict map[string]IScriptBrick,exp string) (*Brick,error) {
	if dict == nil {
		dict = make(map[string]IScriptBrick)
	}

	exp = strings.TrimSpace(exp)

	checkArgExpFuncDict := make(map[string]CheckExpressionArgFunc)
	for key, val := range dict {
		checkArgExpFuncDict[key] = val.GetIsExpressionArg
	}

	brick, err := ParseBrick(exp, checkArgExpFuncDict)
	if err != nil {
		return nil,err
	}

	err = InitFuncDefine(dict, brick)
	if err != nil {
		return nil,err
	}

	return brick,nil
}

func innerCheckIsExpressionArg(index int) bool {
	return false
}

func ParseBrick(exp string, dict map[string]CheckExpressionArgFunc) (*Brick, error) {
	exp = strings.TrimSpace(exp)
	if !checkBracketsMatch(exp) {
		return nil, fmt.Errorf("脚本括号数量不正确")
	}

	funcName, funArgs, err := SplitFuncExpression(exp)
	if err != nil {
		return nil, err
	}

	model := new(Brick)
	model.FuncName = funcName

	isExpFunc, find := dict[funcName]
	if !find {
		isExpFunc = innerCheckIsExpressionArg
	}

	var fargs []*BrickArg
	for idx := range funArgs {
		if isExpFunc(idx) {
			fargs = append(fargs, NewBrickArg(TYPE_STRING, funArgs[idx]))
			continue
		}

		sarr, err := GetFuncSentences(funArgs[idx])
		if err != nil {
			return nil, err
		}
		if len(sarr) > 0 {
			fd, err := ParseBrick(funArgs[idx], dict)
			if err != nil {
				return nil, err
			}
			fargs = append(fargs, NewBrickArg(TYPE_FUNC, fd))
		} else {
			fargs = append(fargs, NewBrickArg(TYPE_STRING, funArgs[idx]))
		}
	}
	model.FuncArgs = fargs
	return model, nil
}
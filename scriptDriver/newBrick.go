package scriptDriver

import (
	"fmt"
	"strings"
)

func CreateBrick(dict map[string]IScriptBrick, exp string) (*Brick, error) {
	if dict == nil {
		dict = make(map[string]IScriptBrick)
	}

	exp = strings.TrimSpace(exp)
	brick, err := ParseBrick(exp)
	if err != nil {
		return nil, err
	}

	err = InitFuncDefine(dict, brick)
	if err != nil {
		return nil, err
	}

	return brick, nil
}

//func innerCheckIsExpressionArg(index int) bool {
//	return false
//}

func ParseBrick(exp string) (*Brick, error) {
	exp = strings.TrimSpace(exp)
	if !checkBracketsMatch(exp) {
		return nil, fmt.Errorf("脚本括号数量不正确")
	}

	funExpArr, err := SplitScriptExpression(exp)
	if err != nil {
		return nil, err
	}
	if len(funExpArr) == 0 {
		return nil, fmt.Errorf("没有识别到函数")
	}

	var lastBrick *Brick = nil
	for _, funExp := range funExpArr {
		funcName, funArgs, err := UnlashScriptExpression(funExp)
		if err != nil {
			return nil, err
		}

		brickModel := new(Brick)
		brickModel.FuncName = funcName

		var fargs []*BrickArg
		if lastBrick != nil {
			fargs = append(fargs, NewBrickArg(TYPE_FUNC, lastBrick))
		}

		for idx := range funArgs {
			funArgs[idx] = strings.TrimSpace(funArgs[idx])
			if checkIsQuotMarkSentence(funArgs[idx]) {
				funArgs[idx] = funArgs[idx][1 : len(funArgs[idx])-1]
				fargs = append(fargs, NewBrickArg(TYPE_STRING, funArgs[idx]))
				continue
			}

			sarr, err := GetFuncSentences(funArgs[idx])
			if err != nil {
				return nil, err
			}
			if len(sarr) > 0 {
				fd, err := ParseBrick(funArgs[idx])
				if err != nil {
					return nil, err
				}
				fargs = append(fargs, NewBrickArg(TYPE_FUNC, fd))
			} else {
				fargs = append(fargs, NewBrickArg(TYPE_STRING, funArgs[idx]))
			}
		}
		brickModel.FuncArgs = fargs
		lastBrick = brickModel
	}

	return lastBrick, nil
}

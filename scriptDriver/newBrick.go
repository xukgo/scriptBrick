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

	//if preParam == nil {
	//	return brick, nil
	//}
	//
	//err = preBuild(brick, preParam)
	//if err != nil {
	//	return nil, err
	//}
	//return brick, nil
	return constBuild(brick)
}

func constBuild(brick *Brick) (*Brick, error) {
	if brick == nil {
		return nil, nil
	}

	if brick.RealFuncMinor == nil {
		return brick, nil
	}

	//if brick.RealFuncMinor.SurplusContext() {
	for idx := range brick.FuncArgs {
		if brick.FuncArgs[idx].Func == nil {
			continue
		}
		nb, err := constBuild(brick.FuncArgs[idx].Func)
		if err != nil {
			return brick, err
		}
		if _, ok := nb.RealFuncMinor.(*ConstBrick); ok {
			val, err := nb.Build(nil)
			if err != nil {
				return brick, err
			}
			brick.FuncArgs[idx] = NewBrickArg(TYPE_OBJECT, val)
		} else {
			brick.FuncArgs[idx].Func = nb
		}
	}

	allConstArg := true
	for idx := range brick.FuncArgs {
		if brick.FuncArgs[idx].Func != nil && !brick.FuncArgs[idx].Func.RealFuncMinor.SurplusContext() {
			allConstArg = false
			break
		}
		if !brick.FuncArgs[idx].CheckIsConstValue() {
			allConstArg = false
			break
		}
	}

	if allConstArg && brick.RealFuncMinor.SurplusContext() {
		val, err := brick.Build(nil)
		if err != nil {
			return brick, err
		}

		resBrick := new(Brick)
		resBrick.FuncName = "const"
		resBrick.RealFuncMinor = NewConstBrick(val)
		return resBrick, nil
	}
	//}

	return brick, nil
}

func ParseBrick(exp string) (*Brick, error) {
	exp = strings.TrimSpace(exp)
	if !checkBracketsMatch(exp) {
		return nil, fmt.Errorf("脚本括号数量不正确:%s", exp)
	}

	funExpArr, err := SplitScriptExpression(exp)
	if err != nil {
		return nil, err
	}
	if len(funExpArr) == 0 {
		return nil, fmt.Errorf("没有识别到函数:%s", exp)
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

//func preBuild(brick *Brick, preParam *PreConstParam) error {
//	if brick == nil {
//		return nil
//	}
//	if brick.FuncName == preParam.PreFuncName {
//		val, err := brick.Build(preParam.Context)
//		if err != nil {
//			return err
//		}
//		brick.RealFuncMinor = new(PreBuildBrick)
//		brick.FuncArgs = []*BrickArg{NewBrickArg(TYPE_OBJECT, val)}
//		return nil
//	}
//
//	for idx := range brick.FuncArgs {
//		err := preBuild(brick.FuncArgs[idx].Func, preParam)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

//func innerCheckIsExpressionArg(index int) bool {
//	return false
//}

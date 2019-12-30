package funcField

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func GetFuncSentences(script string) ([]string, error) {
	//reg := regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9_]{1,64})(\s*)(\([^\)]*\))`)

	var arr []string
	script = strings.TrimSpace(script)
	reg := regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9_]{1,64})(\s*)(\()`)
	for {
		matchStr := reg.FindString(script)
		if len(matchStr) == 0 {
			break
		}

		str, err := getFuncSentence(script, strings.Index(script, matchStr))
		if err != nil {
			return nil, err
		}
		arr = append(arr, str)
		script = strings.Replace(script, str, "@@", 1)
	}
	return arr, nil
}

func getFuncSentence(exp string, start int) (string, error) {
	if !checkBracketsMatch(exp) {
		return "", fmt.Errorf("脚本括号数量不正确")
	}

	bf := new(bytes.Buffer)
	bracktSign := 0
	exp = exp[start:]
	srclen := len(exp)
	for i := 0; i < srclen; i++ {
		if exp[i] == '(' {
			bracktSign++
		} else if exp[i] == ')' {
			if bracktSign == 1 {
				bf.WriteByte(exp[i])
				break
			}
			bracktSign--
		}
		if bracktSign < 0 {
			return "", fmt.Errorf("表达式括号匹配格式不正确")
		}
		bf.WriteByte(exp[i])
	}
	return bf.String(), nil
}

func ParseObjectFuncDefine(exp string) (*FuncObjectMinor, error) {
	exp = strings.TrimSpace(exp)
	if !checkBracketsMatch(exp) {
		return nil, fmt.Errorf("脚本括号数量不正确")
	}

	funcName, funArgs, err := splitFuncExpression(exp)
	if err != nil {
		return nil, err
	}

	model := new(FuncObjectMinor)
	model.FuncName = funcName

	var fargs []*FuncObjectArg
	for idx := range funArgs {
		sarr, err := GetFuncSentences(funArgs[idx])
		if err != nil {
			return nil, err
		}
		if len(sarr) > 0 {
			fd, err := ParseObjectFuncDefine(funArgs[idx])
			if err != nil {
				return nil, err
			}
			fargs = append(fargs, NewFuncObjectArg(TYPE_FUNC, fd))
		} else {
			fargs = append(fargs, NewFuncObjectArg(TYPE_STRING, funArgs[idx]))
		}
	}
	model.FuncArgs = fargs
	return model, nil
}

func ParseStringFuncDefine(exp string) (*FuncStringMinor, error) {
	exp = strings.TrimSpace(exp)
	if !checkBracketsMatch(exp) {
		return nil, fmt.Errorf("脚本括号数量不正确")
	}

	funcName, funArgs, err := splitFuncExpression(exp)
	if err != nil {
		return nil, err
	}

	model := new(FuncStringMinor)
	model.FuncName = funcName

	var fargs []*FuncStringArg
	for idx := range funArgs {
		sarr, err := GetFuncSentences(funArgs[idx])
		if err != nil {
			return nil, err
		}
		if len(sarr) > 0 {
			fd, err := ParseStringFuncDefine(funArgs[idx])
			if err != nil {
				return nil, err
			}
			fargs = append(fargs, NewFuncStringArg(TYPE_FUNC, fd))
		} else {
			fargs = append(fargs, NewFuncStringArg(TYPE_STRING, funArgs[idx]))
		}
	}
	model.FuncArgs = fargs
	return model, nil
}
func checkBracketsMatch(exp string) bool {
	leftCount := strings.Count(exp, "(")
	rightCount := strings.Count(exp, ")")
	if leftCount != rightCount {
		return false
	}

	if leftCount <= 0 {
		return false
	}
	return true
}

func splitFuncExpression(exp string) (string, []string, error) {
	if !checkBracketsMatch(exp) {
		return "", nil, fmt.Errorf("脚本括号数量不正确")
	}
	leftIdx := strings.Index(exp, "(")
	rightIdx := strings.LastIndex(exp, ")")
	if leftIdx > rightIdx {
		return "", nil, fmt.Errorf("括号顺序格式不正确")
	}

	funcName := exp[:leftIdx]
	funcName = strings.TrimSpace(funcName)
	if rightIdx-leftIdx == 1 {
		return funcName, nil, nil
	}

	var args []string
	bracktSign := 0
	exp = exp[leftIdx+1 : rightIdx]
	srclen := len(exp)
	bf := new(bytes.Buffer)
	for i := 0; i < srclen; i++ {
		if exp[i] == '(' {
			bracktSign++
		} else if exp[i] == ')' {
			bracktSign--
		}
		if bracktSign < 0 {
			return "", nil, fmt.Errorf("表达式括号匹配格式不正确")
		}

		if bracktSign == 0 && exp[i] == ',' {
			args = append(args, strings.TrimSpace(bf.String()))
			bf.Reset()
		} else {
			bf.WriteByte(exp[i])
		}
	}
	args = append(args, strings.TrimSpace(bf.String()))
	bf.Reset()
	return funcName, args, nil
}

package scriptDriver

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

const FUNC_MARK = "@@"
const CONTINUE_FUNC_MARK = "@@.@@"

func GetFuncExpression(script string) (string, []string, error) {
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
			return "", nil, err
		}
		script = strings.Replace(script, str, FUNC_MARK, 1)
		script = trimFuncMarkSpace(script)
		if strings.Index(script, CONTINUE_FUNC_MARK) >= 0 {
			arr[len(arr)-1] = arr[len(arr)-1] + "." + str
			script = strings.Replace(script, CONTINUE_FUNC_MARK, FUNC_MARK, 1)
		} else {
			arr = append(arr, str)
		}
	}
	return script, arr, nil
}

func trimFuncMarkSpace(script string) string {
	leftSpaceMark := " " + FUNC_MARK
	rightSpaceMark := FUNC_MARK + " "
	for {
		length := len(script)
		script = strings.ReplaceAll(script, leftSpaceMark, FUNC_MARK)
		if length == len(script) {
			break
		}
	}
	for {
		length := len(script)
		script = strings.ReplaceAll(script, rightSpaceMark, FUNC_MARK)
		if length == len(script) {
			break
		}
	}
	return script
}
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

//把一个方法脚本分解成函数名和参数
func UnlashScriptExpression(exp string) (string, []string, error) {
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

//把函数连接分解成一个个函数
func SplitScriptExpression(exp string) ([]string, error) {
	arr, err := GetFuncSentences(exp)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

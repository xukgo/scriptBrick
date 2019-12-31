package scriptDriver

import (
	"strings"
)

type FuncExecNode struct {
	dict     map[string]IScriptMinor
	funMinor *FuncNodeMinor
}

func NewFuncExecNode(dict map[string]IScriptMinor) *FuncExecNode {
	model := new(FuncExecNode)
	if dict == nil {
		dict = make(map[string]IScriptMinor)
	}
	model.dict = dict
	return model
}

func (this *FuncExecNode) Init(exp string) error {
	exp = strings.TrimSpace(exp)

	checkArgExpFuncDict := make(map[string]CheckExpressionArgFunc)
	for key, val := range this.dict {
		checkArgExpFuncDict[key] = val.GetIsExpressionArg
	}
	fdefine, err := ParseFuncDefine(exp, checkArgExpFuncDict)
	if err != nil {
		return err
	}

	err = InitFuncDefine(this.dict, fdefine)
	if err != nil {
		return err
	}

	this.funMinor = fdefine
	return nil
}

func (this *FuncExecNode) Calc(ctx interface{}) (interface{}, error) {
	var err error
	var v interface{}
	v, err = this.funMinor.Excute(ctx)
	if err != nil {
		return nil, err
	}
	return v, nil
}

//func formatSegIndex(index int) string {
//	return fmt.Sprintf("func[%02d]", index)
//}

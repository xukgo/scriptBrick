package scriptDriver

type IScriptMinor interface {
	Clone() IScriptMinor
	AfterInitCorrectArg(dict map[string]IScriptMinor, index int, arg *FuncNodeArg) error
	GetIsExpressionArg(int) bool
	CheckArgCount(int) bool
	Eval(interface{}, ...interface{}) (interface{}, error)
}

type CheckExpressionArgFunc func(int) bool

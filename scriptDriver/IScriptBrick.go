package scriptDriver

type IScriptBrick interface {
	CloneBasic() IScriptBrick
	AfterInitCorrectArg(dict map[string]IScriptBrick, index int, arg *BrickArg) error
	GetIsExpressionArg(int) bool
	CheckArgCount(int) bool
	Eval(interface{}, ...interface{}) (interface{}, error)
}

type CheckExpressionArgFunc func(int) bool

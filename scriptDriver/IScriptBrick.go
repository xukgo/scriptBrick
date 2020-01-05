package scriptDriver

type IScriptBrick interface {
	CloneBasic() IScriptBrick
	//AfterInitCorrectArg(dict map[string]IScriptBrick, index int, arg *BrickArg) error
	CheckArgCount(int) bool
	Eval(interface{}, ...interface{}) (interface{}, error)
}

type IBrickMountCallback interface {
	AfterInitCorrectArg(dict map[string]IScriptBrick, index int, arg *BrickArg) error
}

type CheckExpressionArgFunc func(int) bool

package scriptDriver

type IScriptBrick interface {
	CloneBasic() IScriptBrick
	//AfterInitCorrectArg(dict map[string]IScriptBrick, index int, arg *BrickArg) error
	SurplusContext() bool //第一个context是否多余的，执行的时候用不到
	CheckArgCount(int) bool
	Eval(interface{}, ...interface{}) (interface{}, error)
}

type IBrickMountCallback interface {
	AfterInitCorrectArg(dict map[string]IScriptBrick, index int, arg *BrickArg) error
}

type CheckExpressionArgFunc func(int) bool

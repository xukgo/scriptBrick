package funcField

type IScriptObjectMinor interface {
	Eval(interface{}, ...interface{}) (interface{}, error)
	CheckArgCount(int) bool
}

package funcField

type IScriptObjectMinor interface {
	IArgCountChecker
	Eval(interface{}, ...interface{}) (interface{}, error)
	CheckArgValid(interface{}, ...interface{}) error
}

type IScriptStringMinor interface {
	IArgCountChecker
	Eval(interface{}, ...string) (interface{}, error)
	CheckArgValid(interface{}, ...string) error
}

type IArgCountChecker interface {
	CheckArgCount(int) bool
}

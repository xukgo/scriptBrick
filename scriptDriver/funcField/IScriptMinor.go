package funcField

type IScriptObjectMinor interface {
	EvalInstance(interface{}, ...interface{}) (interface{}, error)
}

type IScriptStringMinor interface {
	EvalScript(interface{}, ...string) (interface{}, error)
}

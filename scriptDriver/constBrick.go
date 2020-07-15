package scriptDriver

type ConstBrick struct {
	constValue interface{}
}

func NewConstBrick(val interface{}) *ConstBrick {
	model := new(ConstBrick)
	model.constValue = val
	return model
}

func (this *ConstBrick) CloneBasic() IScriptBrick {
	model := new(ConstBrick)
	model.constValue = this.constValue
	return model
}
func (this *ConstBrick) SurplusContext() bool {
	return true
}
func (this *ConstBrick) CheckArgCount(count int) bool {
	return count == 0
}

func (this *ConstBrick) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	return this.constValue, nil
}

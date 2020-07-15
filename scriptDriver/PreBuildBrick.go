package scriptDriver

type PreBuildBrick struct {
}

func (this *PreBuildBrick) CloneBasic() IScriptBrick {
	return new(PreBuildBrick)
}
func (this *PreBuildBrick) SurplusContext() bool {
	return true
}
func (this *PreBuildBrick) Eval(ctx interface{}, args ...interface{}) (interface{}, error) {
	return args[0], nil
}

func (this *PreBuildBrick) CheckArgCount(count int) bool {
	return count == 1
}

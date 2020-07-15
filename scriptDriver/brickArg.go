package scriptDriver

type BrickArg struct {
	MType   int
	Value   interface{}
	Content string
	Func    *Brick
}

func NewBrickArg(mtype int, indata interface{}) *BrickArg {
	model := new(BrickArg)
	model.MType = mtype
	if mtype == TYPE_FUNC {
		model.Func = indata.(*Brick)
	} else if mtype == TYPE_STRING {
		model.Content = indata.(string)
	} else {
		model.Value = indata
	}
	return model
}

func (this *BrickArg) CheckIsConstValue() bool {
	if this.MType == TYPE_FUNC {
		return false
	}
	return true
}

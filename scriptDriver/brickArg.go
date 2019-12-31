package scriptDriver

type BrickArg struct {
	MType   int
	Content string
	Func    *Brick
}

func NewBrickArg(mtype int, indata interface{}) *BrickArg {
	model := new(BrickArg)
	model.MType = mtype
	if mtype == TYPE_FUNC {
		model.Func = indata.(*Brick)
	} else {
		model.Content = indata.(string)
	}
	return model
}

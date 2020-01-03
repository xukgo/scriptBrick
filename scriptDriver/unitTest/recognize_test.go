package unitTest

import (
	"fmt"
	"github.com/xukgo/scriptBrick/scriptDriver"
	"testing"
)

func TestRegexp1(t *testing.T) {
	s1 := "2-(1*3/899+getmap(abc,getinn(666,-1,3a),-1,3)/563+9*+getpro(12abc,-1,3))"
	arr, err := scriptDriver.GetFuncSentences(s1)
	if err != nil {
		t.Fail()
	}
	fmt.Println(arr)
}

func TestUnlashExpression1(t *testing.T) {
	exp := "stringJoin  ( abc  ,  sum(100.789,rand(0.001,999.99)) ,  def100)"
	fname, fargs, err := scriptDriver.UnlashScriptExpression(exp)
	if err != nil {
		t.Fail()
	}
	if fname != "stringJoin" {
		t.Fail()
	}
	if len(fargs) != 3 {
		t.Fail()
	}
	if fargs[0] != "abc" {
		t.Fail()
	}
	if fargs[1] != "sum(100.789,rand(0.001,999.99))" {
		t.Fail()
	}
	if fargs[2] != "def100" {
		t.Fail()
	}
}

func TestSplitExpression1(t *testing.T) {
	exp := "stringJoin  ( abc  ,  sum(100.789,rand(0.001,999.99)) ,  def100).split(aaa,22).len()"
	funSens, err := scriptDriver.GetFuncSentences(exp)
	if err != nil {
		t.Fail()
	}

	if len(funSens) != 3{
		t.Fail()
	}
}

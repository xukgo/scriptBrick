package mathEngine

import (
	"fmt"
	"testing"
)

func TestExecA(t *testing.T) {
	exp := "1+2"
	exec(exp)
}

func TestExecB(t *testing.T) {
	exp := "1+2-4"
	exec(exp)
}

func TestExecC(t *testing.T) {
	exp := "1+2-4*3-8"
	exec(exp)
}

func TestExecD(t *testing.T) {
	exp := "1+2-(4*3-8)"
	exec(exp)
}

func TestExecE(t *testing.T) {
	exp := "1+2-(4*3+(1-8))"
	exec(exp)
}

func TestExecF(t *testing.T) {
	exp := "1+(2-(4*3+(1-8)))"
	exec(exp)
}

func TestExecG(t *testing.T) {
	exp := "((1-2)*(3-8))*((((9+2222))))"
	exec(exp)
}

func TestExecH(t *testing.T) {
	exp := "0.8888-0.1 * 444         -0.2"
	exec(exp)
}

func TestExecI(t *testing.T) {
	exp := "0.8888-0.1 * (444         -0.2)"
	exec(exp)
}

func TestExecJ(t *testing.T) {
	exp := "1_234_567*2-3"
	exec(exp)
}

func TestExecK(t *testing.T) {
	exp := "2.3e4*4/3"
	exec(exp)
}

func TestExecL(t *testing.T) {
	exp := "-1+9-88"
	exec(exp)
}

func TestExecM(t *testing.T) {
	exp := "-1+9-88+(88)"
	exec(exp)
}

func TestExecN(t *testing.T) {
	exp := "-1+9-88+(-88)*666-1"
	exec(exp)
}

func TestExecO(t *testing.T) {
	exp := "-(1)+(3)-(-3)*7-((-3))"
	exec(exp)
}

func TestExecP(t *testing.T) {
	exp := "-(-9+3)"
	exec(exp)
}

func TestExecQ(t *testing.T) {
	exp := "2e-3*2+2e2+1"
	exec(exp)
}

// call mathEngine
func exec(exp string) {
	// input text -> []token
	toks, err := Parse(exp)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	// []token -> AST Tree
	ast := NewAST(toks, exp)
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	fmt.Printf("ExprAST: %+v\n", ar)
	// AST traversal -> result
	r := ExprASTResult(ar)
	fmt.Println("progressing ...\t", r)
	fmt.Printf("%s = %v\n", exp, r)
}

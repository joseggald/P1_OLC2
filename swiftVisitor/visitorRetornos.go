package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionReturnVal(ctx *parser.FuncionReturnValContext) interface{} {
	a := e.Visit(ctx.Expression()).(*SwiftValue)
	e.returnValue = a
	fmt.Println("*******************")
	fmt.Println(e.returnValue)
	if (a==nil){
		a=RETURNVOID
		e.returnValue = a
	}
	fmt.Println(e.returnValue)
	return a
}

func (e *VisitorEvalue) VisitFuncionBreak(ctx *parser.FuncionBreakContext) interface{} {
	e.returnBreak=true
	return BREAK
}

func (e *VisitorEvalue) VisitFuncionReturnVoid(ctx *parser.FuncionReturnVoidContext) interface{} {
	e.returnVoid=true
	return RETURNVOID 
}


package swiftVisitor

import (
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionReturnVal(ctx *parser.FuncionReturnValContext) interface{} {
	a := e.Visit(ctx.Expression()).(*SwiftValue)
	e.returnValue = a
	if (a==nil){
		a=RETURNVOID
		e.returnValue = a
	}
	return a
}

func (e *VisitorEvalue) VisitFuncionBreak(ctx *parser.FuncionBreakContext) interface{} {
	e.returnBreak=true
	
	return BREAK
}

func (e *VisitorEvalue) VisitFuncionReturnVoid(ctx *parser.FuncionReturnVoidContext) interface{} {
	e.returnVoid=true
	e.returnValue = nil
	return RETURNVOID 
}
func (e *VisitorEvalue) VisitFuncionContinue(ctx *parser.FuncionContinueContext) interface{} {
	e.returnContinue=true
	return CONTINUE
}

package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitExprCalFunc(ctx *parser.ExprCalFuncContext) interface{} {
	a := e.Visit(ctx.CallFuncstmt())
	return a.(*SwiftValue)
}

func (e *VisitorEvalue) VisitFuncionDeclaFunc(ctx *parser.FuncionDeclaFuncContext) interface{} {
	functionName := ctx.Id().GetText()
	fmt.Printf("Entro VisitFuncionDeclaFunc\n")
	tipe := ctx.TiposAsign().GetText()
	var params []*Param
	if paramList := ctx.ExprListFunc(); paramList != nil {
		for i := 0; i < len(paramList.AllId()); i += 2 {
			idExterior := paramList.Id(i).GetText()
			idInterior := paramList.Id(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
	} else {
		params = nil
	}
	body := ctx.SentenciasFunc()
	function := NewFunction(params, body, nil, tipe)
	e.currentScope.DeclareFunction(functionName, function)
	fmt.Printf("En FuncionFuncDec - Nombre Variable: %v Valor: %v\n", functionName, function.params)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionDeclaFunc2(ctx *parser.FuncionDeclaFunc2Context) interface{} {
	functionName := ctx.Id().GetText()
	fmt.Printf("Entro VisitFuncionDeclaFunc\n")
	var params []*Param
	if paramList := ctx.ExprListFunc(); paramList != nil {
		for i := 0; i < len(paramList.AllId()); i += 2 {
			idExterior := paramList.Id(i).GetText()
			idInterior := paramList.Id(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
	} else {
		params = nil
	}
	body := ctx.SentenciasFunc()
	function := NewFunction(params, body, nil, "")
	e.currentScope.DeclareFunction(functionName, function)
	fmt.Printf("En FuncionFuncDec2 - Nombre Variable: %v Valor: %v\n", functionName, function.params)
	return VOID
}

func (e *VisitorEvalue) VisitSentenciasFunc(ctx *parser.SentenciasFuncContext) interface{} {
	fmt.Printf("Entro VisitSentenciasFunc\n")

	for _, stmtCtx := range ctx.AllStatement() {
		e.Visit(stmtCtx)
		if e.returnValue != nil {
			fmt.Println(e.returnValue)
			return e.returnValue
		} else if e.returnVoid {
			return RETURNVOID
		} else if e.returnBreak {
			return BREAK
		}
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionCallFunc(ctx *parser.FuncionCallFuncContext) interface{} {
	fmt.Printf("Entering VisitFuncionCallFunc\n")
	functionName := ctx.Id().GetText()
	if function := e.currentScope.FindFunction(functionName); function != nil {
		var args []*SwiftValue
		if exprList := ctx.ExprListCallFunc(); exprList != nil {
			for _, expr := range exprList.AllExpression() {
				argValue := e.Visit(expr).(*SwiftValue)
				args = append(args, argValue)
			}
			ret := function.invoke(e.currentScope, args)
			fmt.Println("FUNCION")
			fmt.Println(ret)
			return ret.(*SwiftValue)
		} else {
			args = nil
			ret := function.invoke(e.currentScope, args)
			fmt.Println(ret)
			res := ret
			if ret == nil {
				ret = RETURNVOID
				e.returnValue = ret
				res = e.returnValue
			} else {
				res = ret.(*SwiftValue)
			}
			return res
		}
	} else {
		fmt.Printf("Function not found: %v\n", functionName)
	}

	return NULL
}

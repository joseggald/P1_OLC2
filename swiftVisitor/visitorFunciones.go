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
		for i := 0; i < len(paramList.AllTiposId()); i += 2 {
			idExterior := paramList.TiposId(i).GetText()
			idInterior := paramList.TiposId(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
	}else if paramList := ctx.ExprListFuncBajo(); paramList != nil {
		for i := 0; i < len(paramList.AllTiposId()); i += 1 {
			idInterior := paramList.TiposId(i).GetText()
			typ := paramList.TiposAsign(i).GetText()
			params = append(params, &Param{
				idExterior: "_",
				idInterior: idInterior,
				typ:        typ,
			})
		}
	} 
	body := ctx.Sentencias()
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
		for i := 0; i < len(paramList.AllTiposId()); i += 2 {
			idExterior := paramList.TiposId(i).GetText()
			idInterior := paramList.TiposId(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
		
	}else if paramList := ctx.ExprListFuncBajo(); paramList != nil {
		for i := 0; i < len(paramList.AllTiposId()); i += 1 {
			idInterior := paramList.TiposId(i).GetText()
			typ := paramList.TiposAsign(i).GetText()
			params = append(params, &Param{
				idExterior: "_",
				idInterior: idInterior,
				typ:        typ,
			})
		}
	} 
	body := ctx.Sentencias()
	function := NewFunction(params, body, nil, "")
	e.currentScope.DeclareFunction(functionName, function)
	fmt.Printf("En FuncionFuncDec2 - Nombre Variable: %v Valor: %v\n", functionName, function.params)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionCallFunc(ctx *parser.FuncionCallFuncContext) interface{} {
	fmt.Printf("Entering VisitFuncionCallFunc\n")
	functionName := ctx.Id().GetText()
	if function := e.currentScope.FindFunction(functionName); function != nil {
		var args []*SwiftValue
		var idExt []string
		if exprList := ctx.ExprListCallFunc(); exprList != nil {
			for _, expr := range exprList.AllExpression() {
				argValue := e.Visit(expr).(*SwiftValue)
				fmt.Println(argValue.value)
				args = append(args, argValue)
			}
			for _, ids := range function.params{
				idExt = append(idExt, ids.idExterior)
			}
			paso:=false
			for i, ids := range exprList.AllTiposId(){
				if idExt[i]==ids.GetText(){
					paso=true
				}else{
					paso=false
					break
				}
			}
			if paso{
				ret := function.invoke(e.currentScope, args)
				return ret.(*SwiftValue)
			}else{
				fmt.Println("error de var funcion")
			}
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
		fmt.Printf("La Función no existe: %v\n", functionName)
	}
	return NULL
}
func (e *VisitorEvalue) VisitFuncionCallFunc2(ctx *parser.FuncionCallFunc2Context) interface{} {
	fmt.Printf("Entering VisitFuncionCallFunc2\n")
	functionName := ctx.Id().GetText()
	if function := e.currentScope.FindFunction(functionName); function != nil {
		var args []*SwiftValue
		if exprList := ctx.ExprVector(); exprList != nil {
			for _, expr := range exprList.AllExpression() {
				argValue := e.Visit(expr).(*SwiftValue)
				args = append(args, argValue)
			}
			ret := function.invoke(e.currentScope, args)
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
		fmt.Printf("La Función no existe: %v\n", functionName)
	}

	return NULL
}
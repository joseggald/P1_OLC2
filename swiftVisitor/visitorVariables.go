package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionAsigExp(ctx *parser.FuncionAsigExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	if newVal.isInt(){
		e.currentScope.DeclareVariable(name, newVal,"Int",false)
	}else if newVal.isBool(){
		e.currentScope.DeclareVariable(name, newVal,"Bool",false)
	}else if newVal.isString(){
		e.currentScope.DeclareVariable(name, newVal,"String",false)
	}else if newVal.isNumber(){
		e.currentScope.DeclareVariable(name, newVal,"Float",false)
	}
	
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoNil(ctx *parser.FuncionAsigTipoNilContext) interface{} {
	fmt.Printf("Entro FuncionAsigNilExp\n")
	newVal := NULL
	name := ctx.Id().GetText()
	println(ctx.TiposAsign().GetText())
	e.currentScope.DeclareVariableNil(name, newVal,ctx.TiposAsign().GetText(),false)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v \n", name)
	
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoExp(ctx *parser.FuncionAsigTipoExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigTipoExp\n")
	valVar := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	existe:=e.currentScope.FindVariable(name)
	if existe!=VOID{
		fmt.Printf("Variable '%v' ya existe en el ámbito actual o en ámbitos padres.\n", name)
	} else {
		if valVar.isString() {
			if ctx.TiposAsign().GetText() == "String" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),false)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isInt() {
			if ctx.TiposAsign().GetText() == "Int" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),false)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isDouble() {
			if ctx.TiposAsign().GetText() == "Float" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),false)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isBool() {
			if ctx.TiposAsign().GetText() == "Bool" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),false)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		}
		fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasign(ctx *parser.FuncionReasignContext) interface{} {
	fmt.Printf("Entro FuncionReasign\n")

	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	a:=e.currentScope.FindVariable(name)
	if a!=VOID{
		if newVal.isInt(){
			e.currentScope.ReassignVariable(name,newVal,"Int")
		}else if newVal.isString(){
			e.currentScope.ReassignVariable(name,newVal,"String")
		}else if newVal.isDouble(){
			e.currentScope.ReassignVariable(name,newVal,"Float")
		}else if newVal.isDouble(){
			e.currentScope.ReassignVariable(name,newVal,"Bool")
		}
		
	}else if e.currentScope.FindVariable(name)== nil{
		println("error no existe variable")
	}
	fmt.Printf("En FuncionReasign - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	
	return VOID
}

package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionAsigExp(ctx *parser.FuncionAsigExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	e.currentScope.DeclareVariable(name, newVal)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoNil(ctx *parser.FuncionAsigTipoNilContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	newVal := NULL
	name := ctx.Id().GetText()
	e.currentScope.DeclareVariable(name, newVal)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoExp(ctx *parser.FuncionAsigTipoExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigTipoExp\n")
	valVar := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	fmt.Println(valVar.isString())
	if valVar.isString() {
		if ctx.TiposAsign().GetText() == "String" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isInt() {
		if ctx.TiposAsign().GetText() == "Int" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isDouble() {
		if ctx.TiposAsign().GetText() == "Float" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isBool() {
		if ctx.TiposAsign().GetText() == "Bool" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	}
	return VOID
}

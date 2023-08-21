package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionIncremento(ctx *parser.FuncionIncrementoContext) interface{} {
	fmt.Printf("Entro FuncionIncremento\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	a:=e.currentScope.FindVariable(name)
	if a!=VOID{
		if newVal.isInt(){
			suma:=newVal.asInt()
			suma=suma+a.asInt()
			e.currentScope.ReassignVariable(name,&SwiftValue{suma},"Int")
		}else if newVal.isString(){
			suma:=newVal.asString()
			suma=suma+a.asString()
			e.currentScope.ReassignVariable(name,&SwiftValue{suma},"String")
		}else if newVal.isDouble(){
			suma:=newVal.asDouble()
			suma=suma+a.asDouble()
			e.currentScope.ReassignVariable(name,&SwiftValue{suma},"Float")
		}else if newVal.isBool(){
			println("error bool")
		}
	}else if e.currentScope.FindVariable(name)== nil{
		println("error no existe variable")
	}
	fmt.Printf("En FuncionIncremento - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	
	return VOID
}
func (e *VisitorEvalue) VisitFuncionDecremento(ctx *parser.FuncionDecrementoContext) interface{} {
	fmt.Printf("Entro FuncionIncremento\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	a:=e.currentScope.FindVariable(name)
	if a!=VOID{
		if newVal.isInt(){
			suma:=newVal.asInt()
			suma=suma-a.asInt()
			e.currentScope.ReassignVariable(name,&SwiftValue{suma},"Int")
		}else if newVal.isString(){
			println("error string")
		}else if newVal.isDouble(){
			suma:=newVal.asDouble()
			suma=suma-a.asDouble()
			e.currentScope.ReassignVariable(name,&SwiftValue{suma},"Float")
		}else if newVal.isBool(){
			println("error bool")
		}
	}else if e.currentScope.FindVariable(name)== nil{
		println("error no existe variable")
	}
	fmt.Printf("En FuncionIncremento - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	
	return VOID
}
func (e *VisitorEvalue) VisitFuncionAsigExp(ctx *parser.FuncionAsigExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	c:=false
	
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	cons:=ctx.TipoInit().GetText()
	if cons=="let"{
		c=true
	}
	if newVal.isInt(){
		e.currentScope.DeclareVariable(name, newVal,"Int",c)
	}else if newVal.isBool(){
		e.currentScope.DeclareVariable(name, newVal,"Bool",c)
	}else if newVal.isString(){
		e.currentScope.DeclareVariable(name, newVal,"String",c)
	}else if newVal.isNumber(){
		e.currentScope.DeclareVariable(name, newVal,"Float",c)
	}
	
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoNil(ctx *parser.FuncionAsigTipoNilContext) interface{} {
	fmt.Printf("Entro FuncionAsigNilExp\n")
	newVal := NULL
	name := ctx.Id().GetText()
	println(ctx.TiposAsign().GetText())
	c:=false
	cons:=ctx.TipoInit().GetText()
	if cons=="let"{
		c=true
	}
	e.currentScope.DeclareVariableNil(name, newVal,ctx.TiposAsign().GetText(),c)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v \n", name)
	
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoExp(ctx *parser.FuncionAsigTipoExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigTipoExp\n")
	valVar := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	existe:=e.currentScope.FindVariable(name)
	c:=false
	cons:=ctx.TipoInit().GetText()
	if cons=="let"{
		c=true
	}
	if existe!=VOID{
		fmt.Printf("Variable '%v' ya existe en el ámbito actual o en ámbitos padres.\n", name)
	} else {
		if valVar.isString() {
			if ctx.TiposAsign().GetText() == "String" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),c)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isInt() {
			if ctx.TiposAsign().GetText() == "Int" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),c)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isDouble() {
			if ctx.TiposAsign().GetText() == "Float" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),c)
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				fmt.Println("Error")
			}
		} else if valVar.isBool() {
			if ctx.TiposAsign().GetText() == "Bool" {
				e.currentScope.DeclareVariable(name, valVar,ctx.TiposAsign().GetText(),c)
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
		}else if newVal.isBool(){
			e.currentScope.ReassignVariable(name,newVal,"Bool")
		}
		
	}else if e.currentScope.FindVariable(name)== nil{
		println("error no existe variable")
	}
	fmt.Printf("En FuncionReasign - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	
	return VOID
}

package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

type returnStruct struct{
	name string
	variables []*AtributoVariable
	structs []*Struct
}

func (e *VisitorEvalue) VisitFuncionDefStruct(ctx *parser.FuncionDefStructContext) interface{} {
	id := ctx.IdMayus().GetText()
	var funcs []*FunctionStruct
	var variables []*AtributoVariable
	fmt.Println(id)
	for _, index := range ctx.AllAtributosLista() {
		argValue := e.Visit(index).(*AtributoVariable)
		fmt.Println(argValue.name)
		variables = append(variables, argValue)
	}
	dataStruct := NewStruct(funcs, variables)
	e.globalScope.DeclareStruct(id, dataStruct)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAtributosListTipo(ctx *parser.FuncionAtributosListTipoContext) interface{} {
	consta:= false
	var id string
	if ctx.Id() != nil {
		id = ctx.Id().GetText()
	} else if ctx.IdMayus() != nil {
		id = ctx.IdMayus().GetText()
	}
	tipo := ctx.TiposAsign().GetText()
	var dato *SwiftValue
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		consta = false
	case parser.SwiftLanLexerLet:
		consta = true
	default:
		return NULL
	}
	dato=NULL
	returndato := NewAtributoVariable(dato, tipo, consta)
	returndato.name=id
	return returndato
}

func (e *VisitorEvalue) VisitFuncionAtributosListExp(ctx *parser.FuncionAtributosListExpContext) interface{} {
	consta:= false
	var id string
	if ctx.Id() != nil {
		id = ctx.Id().GetText()
	} else if ctx.IdMayus() != nil {
		id = ctx.IdMayus().GetText()
	}
	value:=e.Visit(ctx.Expression())
	dato:=value.(*SwiftValue)
	tipo:=""
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		consta = false
	case parser.SwiftLanLexerLet:
		consta = true
	default:
		return NULL
	}
	if dato.isInt() {
		tipo = "Int"
	} else if dato.isString() {
		if dato.isChar() {
			tipo = "Char"
		} else {
			tipo = "String"
		}
	} else if dato.isBool() {
		tipo = "Bool"
	} else if dato.isDouble() {
		tipo = "Float"
	}
	returndato := NewAtributoVariable(dato, tipo, consta)
	returndato.name=id
	return returndato
}

func (e *VisitorEvalue) VisitDefStructExpression(ctx *parser.DefStructExpressionContext) interface{} {
	name:=ctx.IdMayus().GetText()
	fmt.Println(name)
	data:=e.Visit(ctx.ExprListStruct()).(returnStruct)
	fmt.Println(data.name)
	return returnStruct{
		name:name,
		variables:data.variables,
		structs:data.structs,
	}
}

func (e *VisitorEvalue) VisitListAtibStruct(ctx *parser.ListAtibStructContext) interface{} {
	var atributos []*AtributoVariable
	var atributosStructs []*Struct
	var value []*SwiftValue
	var names []string
	var tipos []string
	for _, expr := range ctx.AllExpression() {
		fmt.Println(e.Visit(expr))
		argValue := e.Visit(expr).(*SwiftValue)
		value = append(value, argValue)
		if argValue.isInt() {
			tipos = append(tipos, "Int")
		}else if argValue.isString() {
			if argValue.isChar() {
				tipos = append(tipos, "Char")
			}else{
				tipos = append(tipos, "String")
			}
		}else if argValue.isBool() {
			tipos = append(tipos, "Bool")
		}else if argValue.isDouble() {
			tipos = append(tipos, "Float")
		}
	}
	
	for _,id :=range ctx.AllTiposId(){
		nom:=id.GetText()
		names = append(names, nom)
	}
	
	for _, structs := range ctx.AllStructAsig(){
		structValue:=e.Visit(structs).(*Struct)
		atributosStructs = append(atributosStructs, structValue)
	}
	
	for i := 0; i < len(names); i++ {
		dato:=NewAtributoVariable(value[i],tipos[i],true)
		dato.name=names[i]
		atributos = append(atributos,dato)
	}

	return returnStruct{
		name:"ola",
		variables: atributos,
		structs:atributosStructs,
	}
}


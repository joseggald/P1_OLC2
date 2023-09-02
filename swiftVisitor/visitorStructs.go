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
	var varsStruct []*Struct
	for _, index := range ctx.AllAtributosLista() {
		argValue := e.Visit(index).(*AtributoVariable)
		variables = append(variables, argValue)
	}
	for _, index := range ctx.AllAtributosLista2() {
		argValue := e.Visit(index).(*Struct)
		varsStruct = append(varsStruct, argValue)
	}
	dataStruct := NewStruct(funcs, variables,varsStruct)
	e.globalScope.DeclareStruct(id, dataStruct)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAtributosListTipo(ctx *parser.FuncionAtributosListTipoContext) interface{} {
	consta:= false
	id:= ctx.TiposId().GetText()
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
	id:= ctx.TiposId().GetText()
	
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
	data:=e.Visit(ctx.ExprListStruct()).(returnStruct)
	return returnStruct{
		name:name,
		variables:data.variables,
		structs:data.structs,
	}
}

func (e *VisitorEvalue) VisitListAtibStruct(ctx *parser.ListAtibStructContext) interface{} {
	var atributos []*AtributoVariable
	var value []*SwiftValue
	var names []string
	var tipos []string
	for _, expr := range ctx.AllExpression() {
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
	
	for i := 0; i < len(names); i++ {
		dato:=NewAtributoVariable(value[i],tipos[i],true)
		dato.name=names[i]
		atributos = append(atributos,dato)
	}
	
	return returnStruct{
		name:"",
		variables: atributos,
	}
}

func (e *VisitorEvalue) VisitCallVarStructExpression(ctx *parser.CallVarStructExpressionContext) interface{} {
	id:=ctx.TiposId(0).GetText()
	atrib:=ctx.TiposId(1).GetText()
	contenido:=e.currentScope.findVarStruct(id)
	
	if contenido!=nil{
		for _,atributo := range contenido.variables{
			fmt.Println(atributo.name)
		}
		for _,atributo := range contenido.variables{
			if atributo.name==atrib{
				return atributo.dato
			}
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasigObj(ctx *parser.FuncionReasigObjContext) interface{} {
	id:=ctx.TiposId(0).GetText()
	atrib:=ctx.TiposId(1).GetText()
	contenido:=e.currentScope.findVarStruct(id)
	value:=e.Visit(ctx.Expression()).(*SwiftValue)
	if contenido!=nil{
		for _,atributo := range contenido.variables{
			if atributo.name==atrib{
				e.currentScope.reasigVarStruct(id,value,atrib)
			}
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAtributosStruct(ctx *parser.FuncionAtributosStructContext) interface{} {
		
	return VOID
}
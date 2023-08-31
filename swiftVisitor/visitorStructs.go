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
	id := ctx.Id().GetText()
	var funcs []*FunctionStruct
	var variables []*AtributoVariable

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
	id := ctx.Id().GetText()
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
	id := ctx.Id().GetText()
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
	var atributos []*AtributoVariable
	var atributosStructs []*Struct
	var value []*SwiftValue
	var names []string
	var tipos []string
	if exprList := ctx.ExprListStruct(); exprList != nil {
		for _, expr := range exprList.AllExpression() {
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
		for _, iden := range exprList.AllId(){
			nameValue := iden.GetText()
			names = append(names, nameValue)
		}
		for _, structs := range exprList.AllStructAsig(){
			structValue:=e.Visit(structs).(*Struct)
			atributosStructs = append(atributosStructs, structValue)
		}
	}
	for i := 0; i < len(names); i++ {
		dato:=NewAtributoVariable(value[i],tipos[i],true)
		dato.name=names[i]
		atributos = append(atributos,dato)
	}
	return returnStruct{
		name:name,
		variables: atributos,
		structs:atributosStructs,
	}
}
package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

type returnStruct struct {
	name      string
	variables []*AtributoVariable
	structs   []*Struct
}

type ExpStruct struct {
	tipoStruct   *Struct
	tipoVariable *SwiftValue
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
		fmt.Println(argValue.nameVar)
	}
	for _, index := range ctx.AllFuncStructs() {
		argValue := e.Visit(index).(*FunctionStruct)
		funcs = append(funcs, argValue)
	}
	dataStruct := NewStruct(funcs, variables, varsStruct)
	e.globalScope.DeclareStruct(id, dataStruct)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAtributosListTipo(ctx *parser.FuncionAtributosListTipoContext) interface{} {
	consta := false
	id := ctx.TiposId().GetText()
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
	dato = NULL
	returndato := NewAtributoVariable(dato, tipo, consta)
	returndato.name = id
	return returndato
}

func (e *VisitorEvalue) VisitFuncionAtributosListExp(ctx *parser.FuncionAtributosListExpContext) interface{} {
	consta := false
	id := ctx.TiposId().GetText()

	value := e.Visit(ctx.Expression())
	dato := value.(*SwiftValue)
	tipo := ""
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		consta = false
	case parser.SwiftLanLexerLet:
		consta = true
	default:
		return NULL
	}
	if dato.isNumber() {
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
	returndato.setdata = true
	returndato.name = id
	return returndato
}

func (e *VisitorEvalue) VisitDefStructExpression(ctx *parser.DefStructExpressionContext) interface{} {
	name := ctx.IdMayus().GetText()
	data := e.Visit(ctx.ExprListStruct()).(returnStruct)
	return returnStruct{
		name:      name,
		variables: data.variables,
		structs:   data.structs,
	}
}

func (e *VisitorEvalue) VisitListAtibStruct(ctx *parser.ListAtibStructContext) interface{} {

	var atributos []*AtributoVariable
	var varsStruct []*Struct

	for i, id := range ctx.AllTiposId() {
		nom := id.GetText()
		data := e.Visit(ctx.Expr_struct(i)).(ExpStruct)
		if data.tipoStruct != nil {
			data.tipoStruct.nameVar = nom
			varsStruct = append(varsStruct, data.tipoStruct)
		} else if data.tipoVariable != nil {
			argValue := data.tipoVariable
			tipos := ""
			if argValue.isNumber() {
				tipos = "Int"
			} else if argValue.isString() {
				if argValue.isChar() {
					tipos = "Char"
				} else {
					tipos = "String"
				}
			} else if argValue.isBool() {
				tipos = "Bool"
			} else if argValue.isDouble() {
				tipos = "Float"
			}
			dataVariable := NewAtributoVariable(data.tipoVariable, tipos, false)
			dataVariable.name = nom
			atributos = append(atributos, dataVariable)

		}
	}

	return returnStruct{
		name:      "",
		variables: atributos,
		structs:   varsStruct,
	}
}

func (e *VisitorEvalue) VisitRetornoExpStruct(ctx *parser.RetornoExpStructContext) interface{} {
	if ctx.Expression() != nil {
		// Si se encuentra una expresión, visitarla y obtener su valor
		datoVar := e.Visit(ctx.Expression()).(*SwiftValue)
		defer func() {
			if r := recover(); r != nil {
				cont = cont + 1
				errores.AddError(cont, "error semantico: al recibir la expresion para struct", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}
		}()
		fmt.Println("Expresión encontrada")
		return ExpStruct{
			tipoVariable: datoVar,
			tipoStruct:   nil, // Puedes establecerlo según tus necesidades
		}
	} else if ctx.StructAsig() != nil {
		// Si se encuentra una asignación de estructura, visitarla y obtener su valor
		datoStruct := e.Visit(ctx.StructAsig()).(returnStruct)
		var funcs []*FunctionStruct
		retStruct := NewStruct(funcs, datoStruct.variables, datoStruct.structs)
		fmt.Println("Asignación de estructura encontrada")
		return ExpStruct{
			tipoVariable: nil, // Puedes establecerlo según tus necesidades
			tipoStruct:   retStruct,
		}
	} else {
		cont = cont + 1
		errores.AddError(cont, "error semantico: parametros de struct no validos", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		return nil // Otra opción, dependiendo de tus necesidades
	}
}

func (e *VisitorEvalue) VisitCallVarStructExpression(ctx *parser.CallVarStructExpressionContext) interface{} {
	id := ctx.TiposId(0).GetText()
	atrib := ctx.TiposId(1).GetText()
	contenido := e.currentScope.findVarStruct(id)
	contenido2 := e.currentScope.verifyStructVar(id, atrib)
	if contenido != nil {
		if contenido2 == nil {
			for _, atributo := range contenido.variables {
				fmt.Println(atributo.name)
			}
			for _, atributo := range contenido.variables {
				if atributo.name == atrib {
					return atributo.dato
				}
			}
		} else {
			atrib2 := ctx.TiposId(2).GetText()
			for _, atributo := range contenido2.variables {
				fmt.Println(atributo.name)
			}
			for _, atributo := range contenido2.variables {
				if atributo.name == atrib2 {
					return atributo.dato
				}
			}
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasigObj(ctx *parser.FuncionReasigObjContext) interface{} {
	id := ctx.TiposId(0).GetText()
	atrib := ctx.TiposId(1).GetText()
	contenido := e.currentScope.findVarStruct(id)
	value := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para struct", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	contenido2 := e.currentScope.verifyStructVar(id, atrib)
	if contenido != nil {
		if contenido2 == nil {
			for _, atributo := range contenido.variables {
				if atributo.name == atrib {
					e.currentScope.reasigVarStruct(id, value, atrib)
				}
			}
		} else {
			atrib2 := ctx.TiposId(2).GetText()
			e.currentScope.reasigVarStruct2(id, value, atrib,atrib2)
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAtributosStruct(ctx *parser.FuncionAtributosStructContext) interface{} {
	consta := false
	name := ctx.IdMayus().GetText()
	id_var := ctx.TiposId().GetText()
	cont := e.currentScope.findStruct(name)
	result := NewStruct(cont.funciones, cont.variables, cont.structs)
	result.nameVar = id_var
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		consta = false
	case parser.SwiftLanLexerLet:
		consta = true
	default:
		return NULL
	}
	result.constante = consta
	return result
}

func (e *VisitorEvalue) VisitFuncionCrearFunc(ctx *parser.FuncionCrearFuncContext) interface{} {
	id:=ctx.TiposId().GetText()
	body:=ctx.Sentencias()
	funcion:=NewFunction(nil,body,nil,"")
	dataret:=NewFunctionStruct(funcion,false,id)
	return dataret
}

func (e *VisitorEvalue) VisitFuncionCrearFuncMut(ctx *parser.FuncionCrearFuncMutContext) interface{} {
	id:=ctx.TiposId().GetText()
	body:=ctx.Sentencias()
	funcion:=NewFunction(nil,body,nil,"")
	dataret:=NewFunctionStruct(funcion,false,id)
	return dataret
}

func (e *VisitorEvalue) VisitFuncionCallFuncStrcut(ctx *parser.FuncionCallFuncStrcutContext) interface{} {
	id:=ctx.TiposId(0).GetText()
	idFunc:=ctx.TiposId(1).GetText()
	data:=e.currentScope.findFunctionStruct(id,idFunc)
	data.invoke(e.currentScope,nil)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionSelfReasig(ctx *parser.FuncionSelfReasigContext) interface{} {
	id:=ctx.TiposId().GetText()
	tipoIgual:=ctx.TipoIgual().GetText()
	value:=e.Visit(ctx.Expression()).(*SwiftValue)
	if tipoIgual=="="{
		e.currentScope.reasigVarStruct(e.funcsname,value,id)
	}else if tipoIgual=="+="{
		e.currentScope.increVarStruct(e.funcsname,value,id)
	}else if tipoIgual=="-="{
		e.currentScope.decreVarStruct(e.funcsname,value,id)
	}
	
	return VOID
}

func (e *VisitorEvalue) VisitCallSelfExp(ctx *parser.CallSelfExpContext) interface{} {
	id:=ctx.TiposId().GetText()
	dao:=e.globalScope.findSelfVar(e.funcsname,id)	
	return &SwiftValue{dao}
}

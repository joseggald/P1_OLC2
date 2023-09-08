package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"strconv"
)
func (e *VisitorEvalue) VisitVecStrCallExpression(ctx *parser.VecStrCallExpressionContext) interface{} {
	id:=ctx.TiposId(0).GetText()
	idVar:=ctx.TiposId(1).GetText()
	pos:=e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	dato:=e.globalScope.findFuncVecStr(id,pos.asInt(),idVar)
	return dato
}
func (e *VisitorEvalue) VisitFuncionVectorAsigVarStruct(ctx *parser.FuncionVectorAsigVarStructContext) interface{} {
	id:=ctx.TiposId().GetText()
	var strc []*Struct
	a:=NewVectorStruct(strc)
	e.currentScope.DeclareVectorStruct(id,a)
	simbolos.AddSymbol(id,"Vector Struct",a.nameVar,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	return VOID
}
func (e *VisitorEvalue) VisitFuncionAppendVectorStr(ctx *parser.FuncionAppendVectorStrContext) interface{} {
	
	name:=ctx.TiposId().GetText()
	var variables []*AtributoVariable
	var funcs []*FunctionStruct
	var varsStruct []*Struct
	dataStruct := e.Visit(ctx.StructAsig()).(returnStruct)
	structCont:=e.currentScope.findStruct(dataStruct.name)
	
	for _,vars:=range dataStruct.variables{
		for _,varsS:=range structCont.variables{
			if !varsS.constante{
				if vars.name == varsS.name{
					if vars.tipo == varsS.tipo{
						datoVar:=NewAtributoVariable(vars.dato,vars.tipo,vars.constante)
						datoVar.name=vars.name
						variables = append(variables, datoVar)
					}else{
						cont = cont + 1
						errores.AddError(cont, "error semantico: el tipo de dato no pertenece al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())	
					}
				}
			}
		}
	}
	for _,vars:=range dataStruct.structs{
		varsStruct = append(varsStruct, vars)
	}

	for _,vars:=range structCont.funciones{
		vars.funcion.name=name
		funcs = append(funcs, vars)
	}
	for _,vars:=range structCont.variables{
		if vars.constante || vars.setdata {
			variables = append(variables, vars)
		}
	}
	for _,vars:=range structCont.structs{
		if vars.constante {
			varsStruct = append(varsStruct, vars)
			
		}
	}
	result:=NewStruct(funcs,variables,varsStruct)
	e.currentScope.AddVectorStr(name,result)
	return VOID
}
func (e *VisitorEvalue) VisitFuncionVectorAsig(ctx *parser.FuncionVectorAsigContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	name := ctx.TiposId().GetText()
	vecs := ctx.ExprVector()
	tipo := ctx.TiposAsign().GetText()
	var vectores []*SwiftValue
	if vecs != nil {
		for _, expr := range ctx.ExprVector().AllExpression() {
			argValue := e.Visit(expr).(*SwiftValue)
			defer func() {
				if r := recover(); r != nil {
					cont = cont + 1
					errores.AddError(cont, "error semantico: al recibir la expresion para asignar al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				}
			}()
			vectores = append(vectores, argValue)
		}
	}
	nuevoVector := NewVector(vectores, tipo)
	e.currentScope.DeclareVector(name, nuevoVector)
	simbolos.AddSymbol(name,"Vector",tipo,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	return VOID
}

func (e *VisitorEvalue) VisitFuncionRemoveVec(ctx *parser.FuncionRemoveVecContext) interface{} {
	id := ctx.TiposId().GetText()
	cont2 := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para asignar al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	
	if cont2 != nil {
		datoInt := pos.asInt()
		fmt.Println(datoInt)
		if len(cont2.datos) > datoInt {
			e.currentScope.DelVector(id, datoInt)
		}
	} else {
		fmt.Println("Error posicion no encontrada en el vector")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAppendVector(ctx *parser.FuncionAppendVectorContext) interface{} {
	
	id := ctx.TiposId().GetText()
	cont2 := e.currentScope.FindVector(id)
	val := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para asignar al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	if cont2 != nil {
		if val.isNumber() {
			e.currentScope.AddVector(id, val, "Int")
		} else if val.isString() {
			e.currentScope.AddVector(id, val, "String")
		} else if val.isBool() {
			e.currentScope.AddVector(id, val, "Bool")
		} else if val.isChar() {
			e.currentScope.AddVector(id, val, "Char")
		}
	} else {
		cont = cont + 1
		errores.AddError(cont, "error semantico: posición de vector no encontrada", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionRemoveLastVec(ctx *parser.FuncionRemoveLastVecContext) interface{} {

	id := ctx.TiposId().GetText()
	cont2 := e.currentScope.FindVector(id)
	if cont2 != nil {
		e.currentScope.DelVector(id, len(cont2.datos)-1)
	} else {
		cont = cont + 1
		errores.AddError(cont, "error semantico:no hay datos en el vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionVectorAsigVar(ctx *parser.FuncionVectorAsigVarContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	
	id := ctx.TiposId(0).GetText()
	name2 := ctx.TiposId(1).GetText()
	tipo := ctx.TiposAsign().GetText()
	vectores := e.currentScope.FindVector(name2)
	if vectores != nil {
		tipo2 := e.currentScope.FindTypeVector(name2)
		if tipo2 == tipo {
			e.currentScope.DeclareVector(id, vectores)
			simbolos.AddSymbol(id,"Vector",tipo,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
		} else {
			cont = cont + 1
			errores.AddError(cont, "error semantico:no coinciden los tipos de vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitVecCallExpression(ctx *parser.VecCallExpressionContext) interface{} {
	fmt.Printf("Enter - Vector Expression Statement\n")
	
	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression()).(*SwiftValue)
	var dato interface{}
	if cont != nil {
		datoInt := pos.asInt()
		dato = cont.datos[datoInt]
	} else {
		fmt.Println("Error no hay datos en el vector")
	}
	return dato
}

func (e *VisitorEvalue) VisitFuncionVecReasig(ctx *parser.FuncionVecReasigContext) interface{} {
	fmt.Printf("Enter - Vector Expression Statement\n")
	id := ctx.TiposId().GetText()
	cont2 := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression(0)).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para asignar al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	dataReasig := e.Visit(ctx.Expression(1)).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al recibir la expresion para asignar al vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	var dato interface{}
	if cont2 != nil {
		tipo := e.currentScope.FindTypeVector(id)
		if getValueType(dataReasig.String()) == tipo {
			datoInt := pos.asInt()
			e.currentScope.ReasignVector(id, datoInt, dataReasig)
		} else {
			cont = cont + 1
			errores.AddError(cont, "error semantico: no es del mismo vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		
		}
	} else {
		cont = cont + 1
			errores.AddError(cont, "error semantico: no hay datos en el vector", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		
	}
	return &SwiftValue{dato}
}

func getValueType(value string) string {
	// Intentar convertir a int
	if _, err := strconv.Atoi(value); err == nil {
		return "Int"
	}

	// Intentar convertir a float64
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "Float"
	}

	// Intentar convertir a bool
	if _, err := strconv.ParseBool(value); err == nil {
		return "Bool"
	}

	// Verificar si es un solo carácter
	if len(value) == 1 {
		return "Char"
	}

	// Si no es ningún tipo conocido, entonces considerarlo como string
	return "String"
}


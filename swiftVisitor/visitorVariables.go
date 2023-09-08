package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitFuncionIncremento(ctx *parser.FuncionIncrementoContext) interface{} {
	fmt.Printf("Entro FuncionIncremento\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al asignar expresion", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	name := ctx.TiposId().GetText()
	a := e.currentScope.FindVariable(name)
	if a != VOID {
		if newVal.isInt() {
			suma := newVal.asInt()
			suma = suma + a.asInt()
			e.currentScope.ReassignVariable(name, &SwiftValue{suma}, "Int")
		} else if newVal.isString() {
			suma := newVal.asString()
			suma = suma + a.asString()
			e.currentScope.ReassignVariable(name, &SwiftValue{suma}, "String")
		} else if newVal.isDouble() {
			suma := newVal.asDouble()
			suma = suma + a.asDouble()
			e.currentScope.ReassignVariable(name, &SwiftValue{suma}, "Float")
		} else if newVal.isBool() {
			println("error bool")
		}
	} else if e.currentScope.FindVariable(name) == nil {
		cont = cont + 1
		errores.AddError(cont, "error semantico: no existe la variable", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	fmt.Printf("En FuncionIncremento - Nombre Variable: %v Valor: %v\n", name, newVal.value)

	return VOID
}
func (e *VisitorEvalue) VisitFuncionDecremento(ctx *parser.FuncionDecrementoContext) interface{} {
	fmt.Printf("Entro FuncionIncremento\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al incrementar expresion", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	name := ctx.TiposId().GetText()
	a := e.currentScope.FindVariable(name)
	if a != VOID {
		if newVal.isInt() {
			suma := newVal.asInt()
			suma = suma - a.asInt()
			e.currentScope.ReassignVariable(name, &SwiftValue{suma}, "Int")
		} else if newVal.isString() {
			println("error string")
		} else if newVal.isDouble() {
			suma := newVal.asDouble()
			suma = suma - a.asDouble()
			e.currentScope.ReassignVariable(name, &SwiftValue{suma}, "Float")
		} else if newVal.isBool() {
			println("error bool")
		}
	} else if e.currentScope.FindVariable(name) == nil {
		cont = cont + 1
		errores.AddError(cont, "error semantico: no existe la variable", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	fmt.Printf("En FuncionIncremento - Nombre Variable: %v Valor: %v\n", name, newVal.value)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigExp(ctx *parser.FuncionAsigExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	c := false
	var name string
	Val := e.Visit(ctx.Expression())
	newVal := Val.(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al declarar expresion", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	var dataStr *Struct
	if newVal.isString() {
		dataStr = e.currentScope.findVarStruct(newVal.asString())
	}
	if dataStr != nil {
		e.currentScope.DeclareVarStruct(name, dataStr)
		return VOID
	} else {
		name = ctx.TiposId().GetText()
		switch ctx.GetOp().GetTokenType() {
		case parser.SwiftLanLexerVar:
			c = false
		case parser.SwiftLanLexerLet:
			c = true
		default:
			return NULL
		}
		if newVal.isInt() {
			e.currentScope.DeclareVariable(name, newVal, "Int", c)
			simbolos.AddSymbol(name,"Variable","Int",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
		} else if newVal.isBool() {
			e.currentScope.DeclareVariable(name, newVal, "Bool", c)	
			simbolos.AddSymbol(name,"Variable","Bool",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
		} else if newVal.isString() {
			e.currentScope.DeclareVariable(name, newVal, "String", c)
			simbolos.AddSymbol(name,"Variable","String",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
		} else if newVal.isNumber() {
			e.currentScope.DeclareVariable(name, newVal, "Float", c)
			simbolos.AddSymbol(name,"Variable","Float",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
		}
	}

	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v", name)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoNil(ctx *parser.FuncionAsigTipoNilContext) interface{} {
	fmt.Printf("Entro FuncionAsigNilExp\n")
	newVal := NULL
	name := ctx.TiposId().GetText()
	c := false
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		c = false
	case parser.SwiftLanLexerLet:
		c = true
	default:
		return NULL
	}
	e.currentScope.DeclareVariableNil(name, newVal, ctx.TiposAsign().GetText(), c)
	simbolos.AddSymbol(name,"Variable",ctx.TiposAsign().GetText(),entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
					
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v \n", name)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoExp(ctx *parser.FuncionAsigTipoExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigTipoExp\n")
	valVar := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.TiposId().GetText()
	existe := e.currentScope.FindVariable(name)
	c := false
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerVar:
		c = false
	case parser.SwiftLanLexerLet:
		c = true
	default:
		return NULL
	}
	if existe != VOID {
		fmt.Println("a")
	} else {
		if valVar.isString() {
			if valVar.isChar() {
				if ctx.TiposAsign().GetText() == "Char" {
					e.currentScope.DeclareVariable(name, valVar, ctx.TiposAsign().GetText(), c)
					simbolos.AddSymbol(name,"Variable",ctx.TiposAsign().GetText(),entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
					fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
				}
			} else {
				if ctx.TiposAsign().GetText() == "String" {
					e.currentScope.DeclareVariable(name, valVar, ctx.TiposAsign().GetText(), c)
					simbolos.AddSymbol(name,"Variable",ctx.TiposAsign().GetText(),entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
					
					fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
				} else {
					cont = cont + 1
					errores.AddError(cont, "error semantico: la variable no encaja con su tipo", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
				}
			}

		} else if valVar.isNumber() {
			if ctx.TiposAsign().GetText() == "Int" || ctx.TiposAsign().GetText() == "Float" {
				e.currentScope.DeclareVariable(name, valVar, ctx.TiposAsign().GetText(), c)
				simbolos.AddSymbol(name,"Variable",ctx.TiposAsign().GetText(),entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				cont = cont + 1
				errores.AddError(cont, "error semantico: la variable no encaja con su tipo", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}
		} else if valVar.isBool() {
			if ctx.TiposAsign().GetText() == "Bool" {
				e.currentScope.DeclareVariable(name, valVar, ctx.TiposAsign().GetText(), c)
				simbolos.AddSymbol(name,"Variable",ctx.TiposAsign().GetText(),entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
				fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
			} else {
				cont = cont + 1
				errores.AddError(cont, "error semantico: la variable no encaja con su tipo", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
			}
		}
		
		fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasign(ctx *parser.FuncionReasignContext) interface{} {
	fmt.Printf("Entro FuncionReasign\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al reasignar expresion", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	name := ctx.TiposId().GetText()
	a := e.currentScope.FindVariable(name)

	if a != VOID {
		fmt.Println("entro")
		if newVal.isNumber() {
			fmt.Println("lo envio")
			e.currentScope.ReassignVariable(name, newVal, "Int")
			fmt.Println("regreso")
		} else if newVal.isString() {
			if newVal.isChar() {
				e.currentScope.ReassignVariable(name, newVal, "Char")
			} else {
				e.currentScope.ReassignVariable(name, newVal, "String")
			}
		} else if newVal.isBool() {
			e.currentScope.ReassignVariable(name, newVal, "Bool")
		}

	} else if e.currentScope.FindVariable(name) == VOID {
		cont = cont + 1
		errores.AddError(cont, "error semantico: no existe la variable", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	fmt.Printf("En FuncionAsigStruct - Nombre Variable: %v Valor: %v\n", name, newVal.value)

	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigStruct(ctx *parser.FuncionAsigStructContext) interface{} {
	fmt.Printf("Entro Funcion Asignar Struct\n")

	name := ctx.TiposId().GetText()
	
	var variables []*AtributoVariable
	var funcs []*FunctionStruct
	var varsStruct []*Struct
	dataStruct := e.Visit(ctx.StructAsig()).(returnStruct)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al asignar struct", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	structCont := e.currentScope.findStruct(dataStruct.name)
	defer func() {
		if r := recover(); r != nil {
			cont = cont + 1
			errores.AddError(cont, "error semantico: al asignar struct", entorno, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
		}
	}()
	for _, vars := range dataStruct.variables {
		for _, varsS := range structCont.variables {
			if !varsS.constante {
				if vars.name == varsS.name {
					if vars.tipo == varsS.tipo {
						datoVar := NewAtributoVariable(vars.dato, vars.tipo, vars.constante)
						datoVar.name = vars.name
						variables = append(variables, datoVar)
					}
				}
			}
		}
	}
	for _, vars := range dataStruct.structs {
		varsStruct = append(varsStruct, vars)
	}

	for _, vars := range structCont.funciones {
		vars.funcion.name = name
		funcs = append(funcs, vars)
	}
	for _, vars := range structCont.variables {
		if vars.constante || vars.setdata {
			variables = append(variables, vars)
		}
	}
	for _, vars := range structCont.structs {
		if vars.constante {
			varsStruct = append(varsStruct, vars)
		}
	}
	simbolos.AddSymbol(name,"Variable Struct",dataStruct.name,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	result := NewStruct(funcs, variables, varsStruct)
	e.currentScope.DeclareVarStruct(name, result)
	return VOID
}

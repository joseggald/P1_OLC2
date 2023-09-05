package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"strings"
)

func (e *VisitorEvalue) VisitExprCalFunc(ctx *parser.ExprCalFuncContext) interface{} {
	a := e.Visit(ctx.CallFuncstmt())
	return a.(*SwiftValue)
}

func (e *VisitorEvalue) VisitListaFuncConTipo(ctx *parser.ListaFuncConTipoContext) interface{} {
	var params []*Param
	for _, dataFunc := range ctx.AllDataFuncTipo() {
		returnData := e.Visit(dataFunc).(Param)
		params = append(params, &returnData)
	}
	return params
}

func (e *VisitorEvalue) VisitListaFuncConBarra(ctx *parser.ListaFuncConBarraContext) interface{} {
	var params []*Param
	for _, dataFunc := range ctx.AllDataFuncBajo() {
		returnData := e.Visit(dataFunc).(Param)
		params = append(params, &returnData)
	}
	return params
}

func (e *VisitorEvalue) VisitFuncionDataFuncTipo(ctx *parser.FuncionDataFuncTipoContext) interface{} {
	idExterior := ctx.TiposId(0).GetText()
	idInterior := ctx.TiposId(1).GetText()
	tipo := ctx.TiposAsign().GetText()
	array := ""
	if ctx.INOUT() != nil {
		if ctx.Expr_llave() != nil && ctx.Expr_llave2() != nil {
			if ctx.Expr_llave().GetText() == "[" && ctx.Expr_llave2().GetText() == "]" {
				array = "v"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[" && ctx.Expr_llave2().GetText() == "]]" {
				array = "m"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[[" && ctx.Expr_llave2().GetText() == "]]]" {
				array = "m3"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			}
		} else {
			return Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        tipo,
				array:      "var",
			}
		}
	} else {
		if ctx.Expr_llave() != nil && ctx.Expr_llave2() != nil {
			if ctx.Expr_llave().GetText() == "[" && ctx.Expr_llave2().GetText() == "]" {
				array = "vN"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[" && ctx.Expr_llave2().GetText() == "]]" {
				array = "mN"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[[" && ctx.Expr_llave2().GetText() == "]]]" {
				array = "m3N"
				fmt.Println(array)
				return Param{
					idExterior: idExterior,
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			}
		} else {
			return Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        tipo,
				array:      "varN",
			}
		}
	}
	return Param{
		idExterior: idExterior,
		idInterior: idInterior,
		typ:        tipo,
		array:      "varN",
	}
}

func (e *VisitorEvalue) VisitFuncionDataFuncBajo(ctx *parser.FuncionDataFuncBajoContext) interface{} {
	idInterior := ctx.TiposId().GetText()
	tipo := ctx.TiposAsign().GetText()
	array := ""
	if ctx.INOUT() != nil {
		if ctx.Expr_llave() != nil && ctx.Expr_llave2() != nil {
			if ctx.Expr_llave().GetText() == "[" && ctx.Expr_llave2().GetText() == "]" {
				array = "v"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[" && ctx.Expr_llave2().GetText() == "]]" {
				array = "m"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[[" && ctx.Expr_llave2().GetText() == "]]]" {
				array = "m3"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			}
		} else {

			return Param{
				idExterior: "_",
				idInterior: idInterior,
				typ:        tipo,
				array:      "var",
			}
		}
	} else {
		if ctx.Expr_llave() != nil && ctx.Expr_llave2() != nil {
			if ctx.Expr_llave().GetText() == "[" && ctx.Expr_llave2().GetText() == "]" {
				array = "vN"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[" && ctx.Expr_llave2().GetText() == "]]" {
				array = "mN"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			} else if ctx.Expr_llave().GetText() == "[[[" && ctx.Expr_llave2().GetText() == "]]]" {
				array = "m3N"
				fmt.Println(array)
				return Param{
					idExterior: "_",
					idInterior: idInterior,
					typ:        tipo,
					array:      array,
				}
			}
		} else {
			return Param{
				idExterior: "_",
				idInterior: idInterior,
				typ:        tipo,
				array:      "varN",
			}
		}
	}
	return Param{
		idExterior: "_",
		idInterior: idInterior,
		typ:        tipo,
		array:      "varN",
	}
}

func (e *VisitorEvalue) VisitFuncionDeclaFunc(ctx *parser.FuncionDeclaFuncContext) interface{} {
	functionName := ctx.Id().GetText()
	fmt.Printf("Entro VisitFuncionDeclaFunc\n")
	tipe := ctx.TiposAsign().GetText()
	var params []*Param
	if ctx.ExprListFunc() != nil {
		params = e.Visit(ctx.ExprListFunc()).([]*Param)
	} else if ctx.ExprListFuncBajo() != nil {
		params = e.Visit(ctx.ExprListFuncBajo()).([]*Param)
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
	if ctx.ExprListFunc() != nil {
		params = e.Visit(ctx.ExprListFunc()).([]*Param)
	} else if ctx.ExprListFuncBajo() != nil {
		params = e.Visit(ctx.ExprListFuncBajo()).([]*Param)
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
		var args []*Argumento
		var idExt []string
		if exprList := ctx.ExprListCallFunc(); exprList != nil {
			for i, expr := range exprList.AllExpression() {
				fmt.Printf("Visitando expresión #%d: %s\n", i+1, expr.GetText())
				argValue := e.Visit(expr).(*SwiftValue)
				fmt.Println(argValue.value)
				if argValue.isString() {
					paso := true
					if e.currentScope.FindVector(argValue.asString()) != nil {
						a := e.currentScope.FindVector(argValue.asString())
						args = append(args, &Argumento{
							nomVar:   argValue.asString(),
							variable: nil,
							vector:   a,
							matriz:   nil,
							matriz3d: nil,
						})
						paso = false
					} else if e.currentScope.findMatriz(argValue.asString()) != nil {
						a := e.currentScope.findMatriz(argValue.asString())
						fmt.Println(argValue.asString())
						args = append(args, &Argumento{
							nomVar:   argValue.asString(),
							variable: nil,
							vector:   nil,
							matriz:   a,
							matriz3d: nil,
						})
						paso = false
					} else if e.currentScope.findMatriz3D(argValue.asString()) != nil {
						a := e.currentScope.findMatriz3D(argValue.asString())
						args = append(args, &Argumento{
							nomVar:   argValue.asString(),
							variable: nil,
							vector:   nil,
							matriz:   nil,
							matriz3d: a,
						})
						paso = false
					}
					if paso {
						fmt.Println("paso1")
						args = append(args, &Argumento{
							variable: argValue,
							vector:   nil,
							matriz:   nil,
							matriz3d: nil,
						})
					}
				} else {
					fmt.Println("paso2")
					args = append(args, &Argumento{
						variable: argValue,
						vector:   nil,
						matriz:   nil,
						matriz3d: nil,
					})
				}
			}
			for _, ids := range function.params {
				idExt = append(idExt, ids.idExterior)
			}
			paso := false
			for i, ids := range exprList.AllTiposId() {
				if idExt[i] == ids.GetText() {
					paso = true
				} else {
					paso = false
					break
				}
			}
			if paso {
				ret := function.invoke(e.currentScope, args)
				if ret == nil {
					return VOID
				} else {
					return ret.(*SwiftValue)
				}
			} else {
				fmt.Println("error de var funcion")
			}
		} else {
			args = nil
			ret := function.invoke(e.currentScope, args)
			if ret == nil {
				return VOID
			} else {
				return ret.(*SwiftValue)
			}
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
		var args []*Argumento
		if exprList := ctx.ExprVector(); exprList != nil {

			for _, expr := range exprList.AllExpression() {
				nohay := false
				if strings.Contains(expr.GetText(), "&") {
					
					if strings.Contains(expr.GetText(), "[") {
						if strings.Contains(expr.GetText(), "]") {
							argValue := e.Visit(expr).(*Vector)
							args = append(args, &Argumento{
								nomVar:   "",
								variable: nil,
								vector:   argValue,
								matriz:   nil,
								matriz3d: nil,
							})
						} else {
							nohay = true
						}
					} else {
						nohay = true
					}
				} else {
					nohay = true
				}
				if nohay {
					argValue := e.Visit(expr).(*SwiftValue)
					fmt.Println(argValue.value)
					if argValue.isString() {
						paso := true
						if e.currentScope.FindVector(argValue.asString()) != nil {
							a := e.currentScope.FindVector(argValue.asString())
							args = append(args, &Argumento{
								nomVar:   argValue.asString(),
								variable: nil,
								vector:   a,
								matriz:   nil,
								matriz3d: nil,
							})
							paso = false
						} else if e.currentScope.findMatriz(argValue.asString()) != nil {
							a := e.currentScope.findMatriz(argValue.asString())
							fmt.Println(argValue.asString())
							args = append(args, &Argumento{
								nomVar:   argValue.asString(),
								variable: nil,
								vector:   nil,
								matriz:   a,
								matriz3d: nil,
							})
							paso = false
						} else if e.currentScope.findMatriz3D(argValue.asString()) != nil {
							a := e.currentScope.findMatriz3D(argValue.asString())
							args = append(args, &Argumento{
								nomVar:   argValue.asString(),
								variable: nil,
								vector:   nil,
								matriz:   nil,
								matriz3d: a,
							})
							paso = false
						}
						if paso {
							fmt.Println("paso1")
							args = append(args, &Argumento{
								variable: argValue,
								vector:   nil,
								matriz:   nil,
								matriz3d: nil,
							})
						}
					} else {
						fmt.Println("paso2")
						args = append(args, &Argumento{
							variable: argValue,
							vector:   nil,
							matriz:   nil,
							matriz3d: nil,
						})
					}
				}
			}

			ret := function.invoke(e.currentScope, args)
			if ret == nil {
				return VOID
			} else {
				return ret.(*SwiftValue)
			}

		} else {
			args = nil
			ret := function.invoke(e.currentScope, args)
			if ret == nil {
				return VOID
			} else {
				return ret.(*SwiftValue)
			}
		}
	} else {
		fmt.Printf("La Función no existe: %v\n", functionName)
	}

	return NULL
}

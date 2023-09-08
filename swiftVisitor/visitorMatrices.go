package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

type MatrixResult struct {
	Cols int
	Fila int
	Data []*SwiftValue
}

func (e *VisitorEvalue) VisitFuncionAsignarMatrizNormal(ctx *parser.FuncionAsignarMatrizNormalContext) interface{} {
	fmt.Printf("Enter - Matriz Normal Statement\n")
	name := ctx.Id().GetText()
	tipo := ctx.TiposAsign().GetText()
	var data []*SwiftValue
	var dato *SwiftValue
	exprMatrix := e.Visit(ctx.ExprListMatrixDecla()).(MatrixResult)
	for _,matriz :=range exprMatrix.Data {
		dato = matriz
		fmt.Println(dato.value)
		data = append(data, dato)
	}
	nuevaMatriz := NewMatrix(exprMatrix.Fila, exprMatrix.Cols, tipo)
	nuevaMatriz.SetValuesFromList(data)
	e.currentScope.DeclareMatriz(name, nuevaMatriz)
	simbolos.AddSymbol(name,"Matriz",tipo,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	return VOID
}

func (e *VisitorEvalue) VisitExprListMatrix(ctx *parser.ExprListMatrixContext) interface{} {
	fmt.Printf("Enter - Matrix Expression Normal Statement\n")
	cols := 0
	fila := 0
	var data []*SwiftValue
	for _, index := range ctx.AllExprVector() {
		cols += 1
		for _, expr := range index.AllExpression() {
			fila += 1
			argValue := e.Visit(expr).(*SwiftValue)
			data = append(data, argValue)
		}
	}
	return MatrixResult{
		Cols: cols,
		Fila: fila,
		Data: data,
	}
}

func (e *VisitorEvalue) VisitFuncionAsignarMatriz3D(ctx *parser.FuncionAsignarMatriz3DContext) interface{} {
	fmt.Printf("Enter - Matriz 3D Statement\n")
	name := ctx.Id().GetText()
	tipo := ctx.TiposAsign().GetText()
	var data []*SwiftValue
	cols := 0
	rows := 0
	depth := 0
	var dato *SwiftValue

	for _, index := range ctx.AllExprListMatrixDecla() {
		exprMatrix := e.Visit(index).(MatrixResult)
		for i := 0; i < len(exprMatrix.Data); i++ {
			dato = exprMatrix.Data[i]
			data = append(data, dato)
		}
		if cols == 0 && rows == 0 {
			data = append(data, exprMatrix.Data...)
			rows = exprMatrix.Fila
			cols = exprMatrix.Cols
		} else {
			if exprMatrix.Cols == cols && exprMatrix.Fila == rows {
				data = append(data, exprMatrix.Data...)
				rows = exprMatrix.Fila
				cols = exprMatrix.Cols
			} else {
				println("No cumple con la dimensión")
			}
		}
		depth += 1
	}
	fmt.Println(data)
	nuevaMatriz3D := NewMatrix3D(rows, cols, depth, tipo)
	nuevaMatriz3D.SetDataFromList(data)
	nuevaMatriz3D.Print()
	e.currentScope.DeclareMatriz3D(name, nuevaMatriz3D)
	simbolos.AddSymbol(name,"Matriz",tipo,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	return VOID
}
func (e *VisitorEvalue) VisitFuncionAsignarM3D(ctx *parser.FuncionAsignarM3DContext) interface{} {
	fmt.Printf("Enter - Matriz Normal Statement\n")
	name := ctx.Id().GetText()
	tipo := ctx.TiposAsign(0).GetText()
	tipo_verify := ctx.TiposAsign(1).GetText()
	tipo_verify2 := ctx.TiposAsign(2).GetText()
	tipo_verify3 := ctx.TiposAsign(3).GetText()
	var data []*SwiftValue
	dato := e.Visit(ctx.Expression(0)).(*SwiftValue)
	filaExp := e.Visit(ctx.Expression(3)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(2)).(*SwiftValue)
	depthExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	sizeData := depthExp.asInt() * colExp.asInt() * filaExp.asInt()
	if tipo == tipo_verify && tipo_verify == tipo_verify2 && tipo_verify2 == tipo_verify3 {
		for i := 0; i < sizeData; i++ {
			data = append(data, dato)
		}
		nuevaMatriz := NewMatrix3D(filaExp.asInt(), colExp.asInt(), depthExp.asInt(), tipo)
		nuevaMatriz.SetDataFromList(data)
		nuevaMatriz.Print()
		e.currentScope.DeclareMatriz3D(name, nuevaMatriz)
		simbolos.AddSymbol(name,"Matriz",tipo,entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
	} else {
		fmt.Println("error")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasignMatriz(ctx *parser.FuncionReasignMatrizContext) interface{} {
	fmt.Printf("Enter - Matrix Expression Normal Statement\n")
	name := ctx.Id().GetText()
	filaExp := e.Visit(ctx.Expression(0)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	value := e.Visit(ctx.Expression(2)).(*SwiftValue)
	e.currentScope.ReasignMatriz(name, filaExp.asInt(), colExp.asInt(), value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionReasignMatriz3D(ctx *parser.FuncionReasignMatriz3DContext) interface{} {
	fmt.Printf("Enter - Matrix Expression 3d Statement\n")
	name := ctx.Id().GetText()
	filaExp := e.Visit(ctx.Expression(0)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	depthExp := e.Visit(ctx.Expression(2)).(*SwiftValue)
	value := e.Visit(ctx.Expression(3)).(*SwiftValue)
	e.currentScope.ReasignMatriz3D(name, filaExp.asInt(), colExp.asInt(), depthExp.asInt(), value)
	return VOID
}

func (e *VisitorEvalue) VisitMatriz3DCallExpression(ctx *parser.Matriz3DCallExpressionContext) interface{} {
	fmt.Printf("Enter - Matrix Expression 3d Statement\n")
	name := ctx.TiposId().GetText()
	filaExp := e.Visit(ctx.Expression(0)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	depthExp := e.Visit(ctx.Expression(2)).(*SwiftValue)
	intFila := filaExp.asInt()
	intCol := colExp.asInt()
	intDepth := depthExp.asInt()
	cont := e.currentScope.findMatriz3D(name)
	
	data := cont.GetValue(intFila, intCol, intDepth)
	return data
}

func (e *VisitorEvalue) VisitMatrizCallExpression(ctx *parser.MatrizCallExpressionContext) interface{} {
	fmt.Printf("Enter - Matrix Expression Statement\n")
	name := ctx.TiposId().GetText()
	filaExp := e.Visit(ctx.Expression(0)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	intFila := filaExp.asInt()
	intCol := colExp.asInt()
	cont := e.currentScope.findMatriz(name)
	data := cont.GetValue(intFila, intCol)
	return data
}

func (e *VisitorEvalue) VisitCallMatriz(ctx *parser.CallMatrizContext) interface{} {
	name := ctx.TiposId().GetText()
	filaExp := e.Visit(ctx.Expression(0)).(*SwiftValue)
	colExp := e.Visit(ctx.Expression(1)).(*SwiftValue)
	intFila := filaExp.asInt()
	intCol := colExp.asInt()
	cont := e.currentScope.findMatriz3D(name)
	data:=cont.GetSubMatrixAt(intFila,intCol)
	fina:=NewVector(data,cont.tipo)
	return fina
}
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
	exprMatrix := e.Visit(ctx.ExprListMatrix()).(MatrixResult)
	nuevaMatriz := NewMatrix(exprMatrix.Fila, exprMatrix.Cols, tipo)
	nuevaMatriz.SetValuesFromList(exprMatrix.Data)
	e.currentScope.DeclareMatriz(name, nuevaMatriz)
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
	fmt.Printf("Enter - Matriz Normal Statement\n")
	var data []*SwiftValue
	cols:=0
	rows:=0
	depth:=0
	for _, index := range ctx.AllExprListMatrix() {
		exprMatrix := e.Visit(index).(MatrixResult)
		if cols==0 && rows==0{
			data = append(data, exprMatrix.Data...)
			rows=exprMatrix.Fila
			cols=exprMatrix.Cols
		}else{
			if exprMatrix.Cols == cols && exprMatrix.Fila ==rows{
				data = append(data, exprMatrix.Data...)
				rows=exprMatrix.Fila
				cols=exprMatrix.Cols
			}else{
				println("No cumple con la dimensión")
			}
		}
		depth+=1
	}
	fmt.Println(data)
	nuevaMatriz3D:=NewMatrix3D(rows,cols,depth)
	nuevaMatriz3D.SetDataFromList(data)
	nuevaMatriz3D.Print()
	
	return VOID
}

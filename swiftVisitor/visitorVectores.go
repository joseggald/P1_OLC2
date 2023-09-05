package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"strconv"
)

func (e *VisitorEvalue) VisitFuncionVectorAsig(ctx *parser.FuncionVectorAsigContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	name := ctx.TiposId().GetText()
	vecs := ctx.ExprVector()
	tipo := ctx.TiposAsign().GetText()
	var vectores []*SwiftValue
	if vecs != nil {
		for _, expr := range ctx.ExprVector().AllExpression() {
			argValue := e.Visit(expr).(*SwiftValue)
			vectores = append(vectores, argValue)
		}
	}
	nuevoVector := NewVector(vectores, tipo)
	e.currentScope.DeclareVector(name, nuevoVector)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionRemoveVec(ctx *parser.FuncionRemoveVecContext) interface{} {
	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression()).(*SwiftValue)
	if cont != nil {
		datoInt := pos.asInt()
		fmt.Println(datoInt)
		if len(cont.datos) > datoInt {
			e.currentScope.DelVector(id, datoInt)
		}
	} else {
		fmt.Println("Error posicion no encontrada en el vector")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAppendVector(ctx *parser.FuncionAppendVectorContext) interface{} {
	
	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	val := e.Visit(ctx.Expression()).(*SwiftValue)
	if cont != nil {
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
		fmt.Println("Error posicion no encontrada en el vector")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionRemoveLastVec(ctx *parser.FuncionRemoveLastVecContext) interface{} {

	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	if cont != nil {
		e.currentScope.DelVector(id, len(cont.datos)-1)
	} else {
		fmt.Println("Error no hay datos en el vector")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionVectorAsigVar(ctx *parser.FuncionVectorAsigVarContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	var id string
	id = ctx.TiposId(0).GetText()
	name2 := ctx.TiposId(1).GetText()
	tipo := ctx.TiposAsign().GetText()
	vectores := e.currentScope.FindVector(name2)
	if vectores != nil {
		tipo2 := e.currentScope.FindTypeVector(name2)
		if tipo2 == tipo {
			e.currentScope.DeclareVector(id, vectores)
		} else {
			fmt.Println("NO COINCIDEN LOS TIPOS DE VECTOR")
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
	cont := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression(0)).(*SwiftValue)
	dataReasig := e.Visit(ctx.Expression(1)).(*SwiftValue)

	var dato interface{}
	if cont != nil {
		tipo := e.currentScope.FindTypeVector(id)
		if getValueType(dataReasig.String()) == tipo {
			datoInt := pos.asInt()
			e.currentScope.ReasignVector(id, datoInt, dataReasig)
		} else {
			fmt.Println("Error no es del mismo tipo del vector")
		}
	} else {
		fmt.Println("Error no hay datos en el vector")
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

func (e *VisitorEvalue) VisitFuncionVectorAsigVarStruct(ctx *parser.FuncionVectorAsigVarStructContext) interface{} {

	return VOID
}

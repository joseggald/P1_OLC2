package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"strconv"
)

func (e *VisitorEvalue) VisitFuncionVectorAsig(ctx *parser.FuncionVectorAsigContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	name := ctx.Id().GetText()
	vecs := ctx.ExprVector()
	tipo := ctx.TiposAsign().GetText()
	vectores := []interface{}{}
	if vecs != nil {
		for i := 0; i < len(vecs.AllExpression()); i += 1 {
			dato := vecs.Expression(i).GetText()
			a:=getValueType(dato)
			if a==tipo{
				if tipo=="Int"{
					datoInt,_:=strconv.Atoi(dato)
					vectores = append(vectores, datoInt)
				}else if tipo=="String"{
					
					vectores = append(vectores, dato)
				}else if tipo=="Bool"{
					datoBool,_:= strconv.ParseBool(dato)
					vectores = append(vectores, datoBool)
				}else if tipo=="Char"{
					vectores = append(vectores, dato)
				}else if tipo=="Float"{
					datoFloat,_:=strconv.ParseFloat(dato, 64)
					vectores = append(vectores, datoFloat)
				}
				
			}else{
				fmt.Println("ERROR DE VECTORES")
				return VOID
			}
			
		}
	}
	e.currentScope.DeclareVector(name, VECTO, tipo, vectores)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionRemoveVec(ctx *parser.FuncionRemoveVecContext) interface{}{
	id := ctx.Id().GetText()
	cont := e.currentScope.FindVector(id)
	pos := e.Visit(ctx.Expression()).(*SwiftValue)
	if cont!=nil{
		datoInt:= pos.asInt()
		fmt.Println(datoInt)
		if len(cont)>datoInt{
			e.currentScope.DelVector(id,datoInt)
		}
	}else{
		fmt.Println("Error posicion no encontrada en el vector")
	}
	return VOID
}
func (e *VisitorEvalue) VisitFuncionAppendVector(ctx *parser.FuncionAppendVectorContext) interface{}{
	id := ctx.Id().GetText()
	cont := e.currentScope.FindVector(id)
	val := e.Visit(ctx.Expression()).(*SwiftValue)
	if cont!=nil{
		if val.isInt() {
			e.currentScope.AddVector(id,val.asInt(),"Int")
		}else if val.isDouble() {
			e.currentScope.AddVector(id,val.asDouble(),"Float")
		}else if val.isString() {
			e.currentScope.AddVector(id,val.asString(),"String")
		}else if val.isBool() {
			e.currentScope.AddVector(id,val.asBool(),"Bool")
		}else if val.isChar() {
			e.currentScope.AddVector(id,val.asChar(),"Char")
		}
	}else{
		fmt.Println("Error posicion no encontrada en el vector")
	}
	return VOID
}
func (e *VisitorEvalue) VisitFuncionRemoveLastVec(ctx *parser.FuncionRemoveLastVecContext) interface{}{
	id := ctx.Id().GetText()
	cont := e.currentScope.FindVector(id)

	if cont!=nil{
		e.currentScope.DelVector(id,len(cont)-1)
	}else{
		fmt.Println("Error no hay datos en el vector")
	}
	return VOID
}

func (e *VisitorEvalue) VisitFuncionVectorAsigVar(ctx *parser.FuncionVectorAsigVarContext) interface{} {
	fmt.Printf("Enter - Vector Statement\n")
	name := ctx.Id(0).GetText()
	name2 := ctx.Id(1).GetText()
	tipo := ctx.TiposAsign().GetText()
	vectores:=e.currentScope.FindVector(name2)
	if vectores!=nil{
		tipo2:=e.currentScope.FindTypeVector(name2)
		if tipo2==tipo{
			e.currentScope.DeclareVector(name, VECTO, tipo, vectores)
		}else{
			fmt.Println("NO COINCIDEN LOS TIPOS DE VECTOR")
		}
	}
	return VOID
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
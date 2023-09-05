package swiftVisitor

import (
	"fmt"
	"math"
	"modulo/parser"
	"strconv"
	"strings"
)

func (e *VisitorEvalue) VisitFuncionesEmbeExpression(ctx *parser.FuncionesEmbeExpressionContext) interface{} {
	tipo := ctx.TiposAsign().GetText()
	data := e.Visit(ctx.Expression()).(*SwiftValue)
	if tipo == "Int" {
		if data.isString() {
			num, _ := strconv.Atoi(data.asString())
			return &SwiftValue{num}
		} else if data.isDouble() {
			intNum := int(math.Round(data.asDouble()))
			return &SwiftValue{intNum}
		}
	} else if tipo == "String" {
		strNum := data.String()
		return &SwiftValue{strNum}
	}else if tipo == "Float" {
		if data.isNumber(){
			return &SwiftValue{float64(data.asInt())}
		}else if data.isString(){
			numeroDecimal, _ := strconv.ParseFloat(data.asString(), 64)
			return &SwiftValue{numeroDecimal}
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitExpressionExpression(ctx *parser.ExpressionExpressionContext) interface{} {
	fmt.Printf("Enter - Expression Expression\n")
	val := e.Visit(ctx.Expression())
	return val.(*SwiftValue)
}

func (e *VisitorEvalue) VisitCallArray(ctx *parser.CallArrayContext) interface{} {
	fmt.Printf("Enter - call array\n")
	val := ctx.TiposId().GetText()
	return &SwiftValue{val}
}

func (e *VisitorEvalue) VisitIdExpression(ctx *parser.IdExpressionContext) interface{} {
	id := ctx.Id().GetText()
	variable := e.currentScope.FindVariable(id)
	if variable == VOID {
		str := e.currentScope.findVarStruct(id)
		if str != nil {
			fmt.Println("struct")
			return &SwiftValue{id}
		} else {
			fmt.Printf("Error: Variable '%s' no definida.\n", id)
			return INVALID
		}
	} else {
		return variable
	}
}

func (e *VisitorEvalue) VisitCountExpression(ctx *parser.CountExpressionContext) interface{} {
	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	tam := len(cont.datos)
	return &SwiftValue{tam}
}

func (e *VisitorEvalue) VisitConcatenarExpression(ctx *parser.ConcatenarExpressionContext) interface{} {
	out := ""
	for _, expre := range ctx.AllExpression() {
		visit := e.Visit(expre).(*SwiftValue)
		if len(out) > 1 {
			out = out + " " + fmt.Sprintf("%v", visit.value)
		} else {
			out = fmt.Sprintf("%v", visit.value)
		}
	}
	return &SwiftValue{out}
}

func (e *VisitorEvalue) VisitFuncionNot(ctx *parser.FuncionNotContext) interface{} {
	value := e.Visit(ctx.Expression()).(*SwiftValue)
	nuevoVal := false
	if value.isBool() {
		if value.asBool() {
			nuevoVal = false
		} else {
			nuevoVal = true
		}
	}
	return &SwiftValue{nuevoVal}
}

func (e *VisitorEvalue) VisitEmptyVecExpression(ctx *parser.EmptyVecExpressionContext) interface{} {
	id := ctx.TiposId().GetText()
	cont := e.currentScope.FindVector(id)
	tam := len(cont.datos)
	ret := false
	if tam < 1 {
		ret = true
	} else {
		ret = false
	}
	return &SwiftValue{ret}
}

func (e *VisitorEvalue) VisitNumberExpression(ctx *parser.NumberExpressionContext) interface{} {
	numero, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return &SwiftValue{numero}
}

func (e *VisitorEvalue) VisitBoolExpression(ctx *parser.BoolExpressionContext) interface{} {
	valueBool := ctx.GetText() == "true"
	return &SwiftValue{valueBool}
}

func (e *VisitorEvalue) VisitStringExpression(ctx *parser.StringExpressionContext) interface{} {
	texto := ctx.GetText()
	texto = texto[1 : len(texto)-1]
	return &SwiftValue{texto}
}
func (e *VisitorEvalue) VisitNilExpression(ctx *parser.NilExpressionContext) interface{} {
	return NULL
}

func (e *VisitorEvalue) VisitExpressionSumRes(ctx *parser.ExpressionSumResContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerSuma:
		return e.suma(left, right)
	case parser.SwiftLanLexerResta:
		return e.resta(left, right)
	default:
		return NULL
	}
}

func (e *VisitorEvalue) VisitExpressionMultDivMod(ctx *parser.ExpressionMultDivModContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerMult:
		return e.mult(left, right)
	case parser.SwiftLanLexerDiv:
		return e.div(left, right)
	case parser.SwiftLanLexerMod:
		return e.mod(left, right)
	default:
		return NULL
	}
}

func (e *VisitorEvalue) suma(left *SwiftValue, right *SwiftValue) interface{} {
	var er = true
	if left.isNumber() && right.isNumber() {
		er = false
		return &SwiftValue{left.asDouble() + right.asDouble()}
	}
	if left.isString() && right.isString() {
		er = false
		return &SwiftValue{left.asString() + right.asString()}
	}
	if er {
		fmt.Println("Error de operación suma")
	}
	return NULL
}

func (e *VisitorEvalue) resta(left *SwiftValue, right *SwiftValue) interface{} {
	return &SwiftValue{left.asDouble() - right.asDouble()}
}

func (e *VisitorEvalue) mult(left *SwiftValue, right *SwiftValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() * right.asDouble()}
	}
	if left.isString() && right.isNumber() {
		times := right.asInt()
		sb := strings.Builder{}
		for i := 0; i < times; i++ {
			sb.WriteString(left.asString())
		}
		return &SwiftValue{sb.String()}
	}
	return NULL
}

func (e *VisitorEvalue) div(left *SwiftValue, right *SwiftValue) interface{} {
	return &SwiftValue{left.asDouble() / right.asDouble()}
}

func (e *VisitorEvalue) mod(left *SwiftValue, right *SwiftValue) interface{} {
	var er = true
	if left.isInt() && right.isInt() {
		er = false
		return &SwiftValue{left.asInt() % right.asInt()}
	}
	if er {
		fmt.Println("Error de operación mod")
	}
	return NULL
}

func (e *VisitorEvalue) VisitFuncionUnariaExp(ctx *parser.FuncionUnariaExpContext) interface{} {
	valor := e.Visit(ctx.Expression()).(*SwiftValue)
	if valor.isInt() {
		return &SwiftValue{value: -valor.asInt()}
	}
	if valor.isDouble() {
		return &SwiftValue{value: -valor.asDouble()}
	}
	return NULL
}

func (e *VisitorEvalue) VisitFuncionPowExp(ctx *parser.FuncionPowExpContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	return &SwiftValue{math.Pow(left.asDouble(), right.asDouble())}
}

func (e *VisitorEvalue) VisitFuncionCompExp(ctx *parser.FuncionCompExpContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerMenorQue:
		return e.menorQue(left, right)
	case parser.SwiftLanLexerMinQEquals:
		return e.menorIgualQue(left, right)
	case parser.SwiftLanLexerMayorQue:
		return e.mayorQue(left, right)
	case parser.SwiftLanLexerMayQEquals:
		return e.mayorIgualQue(left, right)
	default:
		return NULL
	}
}

func (e *VisitorEvalue) menorQue(left *SwiftValue, right *SwiftValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() < right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &SwiftValue{left.asString() < right.asString()}
	}
	return false
}

func (e *VisitorEvalue) menorIgualQue(left *SwiftValue, right *SwiftValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() <= right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &SwiftValue{left.asString() <= right.asString()}
	}
	return false
}

func (e *VisitorEvalue) mayorQue(left *SwiftValue, right *SwiftValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() > right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &SwiftValue{left.asString() > right.asString()}
	}
	return false
}

func (e *VisitorEvalue) mayorIgualQue(left *SwiftValue, right *SwiftValue) interface{} {
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() >= right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &SwiftValue{left.asString() >= right.asString()}
	}
	return false
}

func (e *VisitorEvalue) VisitFuncionEqExp(ctx *parser.FuncionEqExpContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	switch ctx.GetOp().GetTokenType() {
	case parser.SwiftLanLexerEquals:
		return &SwiftValue{value: left.equals(right)}
	case parser.SwiftLanLexerNotEquals:
		return &SwiftValue{value: !left.equals(right)}
	}
	return NULL
}

func (e *VisitorEvalue) VisitFuncionAndExp(ctx *parser.FuncionAndExpContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	return &SwiftValue{left.asBool() && right.asBool()}
}

func (e *VisitorEvalue) VisitFuncionOrExp(ctx *parser.FuncionOrExpContext) interface{} {
	left := e.Visit(ctx.Expression(0)).(*SwiftValue)
	right := e.Visit(ctx.Expression(1)).(*SwiftValue)
	return &SwiftValue{left.asBool() || right.asBool()}
}

func (e *VisitorEvalue) VisitFuncionTernaryExp(ctx *parser.FuncionTernaryExpContext) interface{} {
	condition := e.Visit(ctx.Expression(0)).(*SwiftValue)
	if condition.asBool() {
		return e.Visit(ctx.Expression(1)).(*SwiftValue)
	} else {
		return e.Visit(ctx.Expression(2)).(*SwiftValue)
	}
}

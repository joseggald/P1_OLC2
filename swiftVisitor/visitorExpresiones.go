package swiftVisitor

import (
	"fmt"
	"math"
	"modulo/parser"
	"strconv"
	"strings"
)


func (e *VisitorEvalue) VisitIdExpression(ctx *parser.IdExpressionContext) interface{} {
	id := ctx.Id().GetText()
	variable := e.currentScope.FindVariable(id)
	if variable == VOID {
		fmt.Printf("Error: Variable '%s' no definida.\n", id)
		return INVALID
	}
	return variable
}

func (e *VisitorEvalue) VisitCountExpression(ctx *parser.CountExpressionContext) interface{} {
	id := ctx.Id().GetText()
	cont := e.currentScope.FindVector(id)
	tam:=len(cont)
	return &SwiftValue{tam}
}

func (e *VisitorEvalue) VisitEmptyVecExpression(ctx *parser.EmptyVecExpressionContext) interface{} {
	id := ctx.Id().GetText()
	cont := e.currentScope.FindVector(id)
	tam:=len(cont)
	ret:=false
	if tam<1{
		ret=true
	}else{
		ret=false
	}
	return &SwiftValue{ret}
}

func (e *VisitorEvalue) VisitEnteroExpression(ctx *parser.EnteroExpressionContext) interface{} {
	numero, _ := strconv.Atoi(ctx.GetText())
	return &SwiftValue{numero}
}

func (e *VisitorEvalue) VisitFloatExpression(ctx *parser.FloatExpressionContext) interface{} {
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
	if left.isNumber() && right.isNumber() {
		return &SwiftValue{left.asDouble() + right.asDouble()}
	}
	if left.isString() && right.isString() {
		return &SwiftValue{left.asString() + right.asString()}
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
	a := int(left.asDouble())
	b := int(right.asDouble())
	return &SwiftValue{a % b}
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
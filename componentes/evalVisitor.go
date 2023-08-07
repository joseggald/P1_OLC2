package componentes

import (
	"fmt"
	"modulo/parser"
	"reflect"
	"strconv"
	"github.com/antlr4-go/antlr/v4"
)
var out=""
type EvalVisitor struct {
	parser.BaseSwiftLanVisitor
	scope     *Scope
	functions map[string]*Function
}

func OutData()string{
	return out;
}

func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{
		scope:     NewScope(),
		functions: make(map[string]*Function),
	}
}

func (e *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("Visitando: %v\n", reflect.TypeOf(tree))
	switch t := tree.(type) {	
		case *antlr.ErrorNodeImpl:
			fmt.Printf("Error: - %v\n", t.GetText())
			return nil
		default:
			return tree.Accept(e)
	}
}

func (e *EvalVisitor) VisitInicio(ctx *parser.InicioContext) interface{} {
	fmt.Printf("Calculating Program: %s\n", ctx.GetText())
	out=""
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		// fmt.Printf("> %s\n", n.(antlr.ParseTree).GetText())
		e.Visit(n.(antlr.ParseTree))
	}
	return VOID
}

func (e *EvalVisitor) VisitSentencias(ctx *parser.SentenciasContext) interface{} {
	fmt.Printf("Entro - Sentencias\n")
	for _, sctx := range ctx.AllStatement() {
		e.Visit(sctx)
	}
	return VOID
}

func (e *EvalVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Printf("Entro VisitStatement\n")
	return e.VisitChildren(ctx)
}

func (e *EvalVisitor) VisitPrintLlamadaPrint(ctx *parser.PrintLlamadaPrintContext) interface{} {
	fmt.Printf("Entro VisitPrint\n")
	if expr := ctx.Expression(); expr != nil {
		fmt.Printf("%s\n", e.Visit(expr).(*TLValue).String())
		out=out + e.Visit(expr).(*TLValue).String() +"\n"
	} else {
		fmt.Println()
	}
	return VOID
}


func (e *EvalVisitor) VisitNumberExpression(ctx *parser.NumberExpressionContext) interface{} {
	num, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return &TLValue{num}
}

func (e *EvalVisitor) VisitBoolExpression(ctx *parser.BoolExpressionContext) interface{} {
	val := ctx.GetText() == "true"
	return &TLValue{val}
}

func (e *EvalVisitor) VisitStringExpression(ctx *parser.StringExpressionContext) interface{} {
	text := ctx.GetText()
	text = text[1 : len(text)-1]
	return &TLValue{text}
}
package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"reflect"
	"github.com/antlr4-go/antlr/v4"
)

var out = ""

type VisitorEvalue struct {
	parser.BaseSwiftLanVisitor
	globalScope  *Scope
	currentScope *Scope
	returnValue  interface{}
	returnVoid 	bool
	returnBreak 	bool
	returnContinue bool
}

func OutData() string {
	return out
}

func NewVisitorEvalue() *VisitorEvalue {
	globalScope := NewScope()
	return &VisitorEvalue{
		globalScope:  globalScope,
		currentScope: globalScope,
	}
}

func (e *VisitorEvalue) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("Visitando: %v\n", reflect.TypeOf(tree))
	switch t := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("Error: - %v\n", t.GetText())
		return nil
	default:
		return tree.Accept(e)
	}
}

func (e *VisitorEvalue) VisitInicio(ctx *parser.InicioContext) interface{} {
	fmt.Printf("Calculating Program: %s\n", ctx.GetText())
	out = ""
	return e.VisitChildren(ctx)
}

func (e *VisitorEvalue) VisitChildren(node antlr.RuleNode) interface{} {
	for _, n := range node.GetChildren() {
		e.Visit(n.(antlr.ParseTree))
	}
	return VOID
}

func (e *VisitorEvalue) VisitSentencias(ctx *parser.SentenciasContext) interface{} {
	fmt.Printf("Entro - Sentencias\n")
	for _, StamentsCtx := range ctx.AllStatement() {
		e.Visit(StamentsCtx)
		if e.returnValue != nil && !e.returnVoid{
			fmt.Println(e.returnValue)
			return e.returnValue
		}else if e.returnVoid{
			return RETURNVOID
		}else if e.returnBreak{
			return BREAK
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Printf("Entro VisitStatement\n")
	childCount := ctx.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := ctx.GetChild(i).(antlr.ParseTree)
		e.Visit(child)
	}
	return VOID

}

func (e *VisitorEvalue) VisitFuncionPrint(ctx *parser.FuncionPrintContext) interface{} {
	fmt.Printf("Entro VisitPrint\n")
	if expr := ctx.Expression(); expr != nil {
		fmt.Printf("%s\n", e.Visit(expr).(*SwiftValue).String())
		out = out + e.Visit(expr).(*SwiftValue).String() + "\n"
	} else {
		fmt.Println()
	}
	return VOID
}

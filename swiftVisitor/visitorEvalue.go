package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"reflect"

	"github.com/antlr4-go/antlr/v4"
)

var out = ""
var errores=NewErrorList()
var simbolos=NewSymbolList()
var cont int
var entorno string

type VisitorEvalue struct {
	parser.BaseSwiftLanVisitor
	globalScope    *Scope
	currentScope   *Scope
	returnValue    interface{}
	returnVoid     bool
	returnBreak    bool
	returnContinue bool
	funcsname string
}

func OutData() string {
	return out
}
func OutErrors() *ErrorList {
	return errores
}
func OutSimbolos() *SymbolList {
	return simbolos
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
		cont=cont+1
		fmt.Printf("Error: - %v\n", t.GetText())
		errores.AddError(cont,"Error Sintactico:"+t.GetText(),entorno,t.GetSymbol().GetLine(),t.GetSymbol().GetColumn())
		return nil
	default:
		return tree.Accept(e)
	}
}

func (e *VisitorEvalue) VisitInicio(ctx *parser.InicioContext) interface{} {
	fmt.Printf("Calculating Program: %s\n", ctx.GetText())
	errores.ResetErrors()
	simbolos.ResetSymbols()
	cont=0
	out = ""
	entorno="Global"
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
		if e.returnValue != nil{
			fmt.Println(e.returnValue)
			return e.returnValue
		} else if e.returnVoid {
			return RETURNVOID
		} else if e.returnBreak {
			return BREAK
		} else if e.returnContinue {
			return CONTINUE
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
		a := e.Visit(expr).(*SwiftValue)
		defer func() {
			if r := recover(); r != nil {
				cont=cont+1
				errores.AddError(cont,"error semantico: al recibir la expresion a imprimir",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
			}
		}()
		if a == NULL {
			out = out + "nil" + "\n"
		} else {
			out = out + a.String() + "\n"
		}
	}
	if expr := ctx.ConcaExp(); expr != nil {
		a := e.Visit(expr).(*SwiftValue)
		defer func() {
			if r := recover(); r != nil {
				cont=cont+1
				errores.AddError(cont,"error al recibir la expresion a imprimir",entorno,ctx.GetStart().GetLine(),ctx.GetStart().GetColumn())
			}
		}()
		if a == NULL {
			out = out + "nil" + "\n"
		} else {
			out = out + a.String() + "\n"
		}
	}
	return VOID
}
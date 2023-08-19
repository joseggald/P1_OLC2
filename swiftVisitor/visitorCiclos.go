package swiftVisitor

import (
	"fmt"
	"modulo/parser"
	"strconv"
	"strings"
)

func (e *VisitorEvalue) VisitFuncionForstmt(ctx *parser.FuncionForstmtContext) interface{} {
	fmt.Printf("Enter - For Statement\n")
	rango := ctx.EnteroRange().GetText()
	println(rango)
	parts := strings.Split(rango, "...")
	start,_ := strconv.Atoi(parts[0])

	end,_ := strconv.Atoi(parts[1])

	forScope := e.currentScope.CreateChildScope()
	id := ctx.Id().GetText()
	for i := start; i <= end; i++ {
        loopVarValue := &SwiftValue{value: i}
        forScope.DeclareVariable(id, loopVarValue,"Int",false)
        evalVisitor := &VisitorEvalue{currentScope: forScope, globalScope: e.currentScope}
        eval:=evalVisitor.Visit(ctx.Sentencias())
		if eval==RETURNVOID{
			return RETURNVOID
		}else if eval==BREAK{
			break
		}else if eval==CONTINUE{
			continue
		} 
    }

	return VOID
}
func (e *VisitorEvalue) VisitFuncionWhilestmt(ctx *parser.FuncionWhilestmtContext) interface{} {
	fmt.Printf("Enter - While Statement\n")
	whileScope := e.currentScope.CreateChildScope()
	for e.Visit(ctx.Expression()).(*SwiftValue).asBool() {
		evalVisitor := &VisitorEvalue{currentScope: whileScope, globalScope: e.currentScope}
        eval:=evalVisitor.Visit(ctx.Sentencias())
		if eval==RETURNVOID{
			return RETURNVOID
		}else if eval==BREAK{
			break
		}else if eval==CONTINUE{
			continue
		} 
	}

	return VOID
}

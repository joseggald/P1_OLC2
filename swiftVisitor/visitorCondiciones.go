package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
    fmt.Printf("Entro VisitIf\n")
    condition := e.Visit(ctx.Ifstat().(*parser.IfstatContext).Expression()).(*SwiftValue)
    if condition.isBool() && condition.asBool() {
        ifScope := e.currentScope.CreateChildScope()
        e.currentScope = ifScope
        defer func() {
            e.currentScope = ifScope.parent
        }()
		e.Visit(ctx.Ifstat().(*parser.IfstatContext).Sentencias())

        if e.returnValue != nil {
            return e.returnValue
        }else if e.returnVoid{
			return RETURNVOID
		}else if e.returnBreak{
			return BREAK
		}else if e.returnContinue{
			return BREAK
		}
        return VOID
    }

    for _, elseIfStat := range ctx.AllElseifstmt() {
        elseIfCondition := e.Visit(elseIfStat.(*parser.ElseifstmtContext).Expression()).(*SwiftValue)
        if elseIfCondition.isBool() && elseIfCondition.asBool() {
            elseIfScope := e.currentScope.CreateChildScope()
            e.currentScope = elseIfScope
            defer func() {
                e.currentScope = elseIfScope.parent
            }()
            e.Visit(elseIfStat.(*parser.ElseifstmtContext).Sentencias())
            if e.returnValue != nil {
                return e.returnValue
            }
            return VOID
        }
    }

    if elseContext := ctx.Elsestmt(); elseContext != nil {
        e.Visit(elseContext.Sentencias())
        if e.returnValue != nil {
            return e.returnValue
        }
    }

    return VOID
}
package swiftVisitor

import (
	"fmt"
	"modulo/parser"
)

func (e *VisitorEvalue) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
    fmt.Printf("Entro VisitIf\n")
    entorno="If condition"
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
    entorno="ElseIf condition"
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
    entorno="Else condition"
    if elseContext := ctx.Elsestmt(); elseContext != nil {
        e.Visit(elseContext.Sentencias())
        if e.returnValue != nil {
            return e.returnValue
        }
    }

    return VOID
}

func (e *VisitorEvalue) VisitGuardstmt(ctx *parser.GuardstmtContext) interface{} {
    fmt.Printf("Entro VisitGuard\n")
    entorno="Guard"
    condition := e.Visit(ctx.Expression()).(*SwiftValue)
    if condition.isBool(){
        if condition.asBool(){
            fmt.Println("Guard condition not met")
            return VOID
        }  
    }
    
    guardScope := e.currentScope.CreateChildScope()
    e.currentScope = guardScope
    defer func() {
        e.currentScope = guardScope.parent
    }()
    
    e.Visit(ctx.Sentencias())
    
    if e.returnValue != nil {
        return e.returnValue
    } else if e.returnVoid {
        return RETURNVOID
    } else if e.returnBreak {
        return BREAK
    } else if e.returnContinue {
        return CONTINUE
    }
    
    return VOID
}
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
	parts := strings.Split(rango, "...")
	fmt.Println(parts[0])
	var start int
	var end int
	a := e.currentScope.FindVariable(parts[0])
	b := e.Visit(ctx.Expression()).(*SwiftValue)
	if a != VOID {
		dat := a.asInt()
		fmt.Println(dat)
		s := dat
		start = s
	} else {
		s, _ := strconv.Atoi(parts[0])
		start = s
	}
	if b != VOID {
		dat := b.asInt()
		fmt.Println(dat)
		s := dat
		end = s
	} else {
		s, _ := strconv.Atoi(parts[1])
		end = s
	}

	forScope := e.currentScope.CreateChildScope()
	var id string
	if ctx.TiposId() != nil {
		id = ctx.TiposId().GetText()
	} else {
		fmt.Println("salto")
		id = ""
	}

	forScope.DeclareVariable(id, &SwiftValue{value: start}, "Int", false)
	for i := start; i <= end; i++ {
		loopVarValue := &SwiftValue{value: i}
		forScope.DeclareVariable(id, loopVarValue, "Int", false)
		evalVisitor := &VisitorEvalue{currentScope: forScope, globalScope: e.currentScope}
		eval := evalVisitor.Visit(ctx.Sentencias())
		if eval == RETURNVOID {
			return RETURNVOID
		} else if eval == BREAK {
			break
		} else if eval == CONTINUE {
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
		eval := evalVisitor.Visit(ctx.Sentencias())
		if eval == RETURNVOID {
			return RETURNVOID
		} else if eval == BREAK {
			break
		} else if eval == CONTINUE {
			continue
		}
	}

	return VOID
}
func (e *VisitorEvalue) VisitFuncionSwitchstmt(ctx *parser.FuncionSwitchstmtContext) interface{} {
	fmt.Println("Enter - Switch Statement")

	switchExprValue := e.Visit(ctx.Expression()).(*SwiftValue)
	switchScope := e.currentScope.CreateChildScope()
	defa := false

	for _, caseCtx := range ctx.AllBloqueCase() {
		caseExprValue := e.Visit(caseCtx.Expression()).(*SwiftValue)
		if switchExprValue.equals(caseExprValue) {
			caseScope := switchScope.CreateChildScope()
			evalVisitor := &VisitorEvalue{currentScope: caseScope, globalScope: e.currentScope}
			evalVisitor.Visit(caseCtx.Sentencias())
			defa = false
			return VOID
		} else {
			defa = true
		}
	}

	if defa {
		defaultScope := switchScope.CreateChildScope()
		evalVisitor := &VisitorEvalue{currentScope: defaultScope, globalScope: e.currentScope}
		evalVisitor.Visit(ctx.Sentencias())
		return VOID
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionForExpstmt(ctx *parser.FuncionForExpstmtContext) interface{} {
	fmt.Printf("Enter - For Statement\n")
	forScope := e.currentScope.CreateChildScope()
	id := ctx.TiposId().GetText()
	dataEnd := ctx.String_().GetText()
	dataStr := strings.Trim(dataEnd, "\"")
	fmt.Println("String")
	b := dataStr
	letters := []rune(b)
	fmt.Println(id)
	fmt.Println(b)
	forScope.DeclareVariable(id, &SwiftValue{value: "a"}, "String", false)
	for _, char := range letters {
		c := string(char)
		fmt.Println(c)
		loopVarValue := &SwiftValue{value: c}
		forScope.ReassignVariable(id, loopVarValue, "String")
		evalVisitor := &VisitorEvalue{currentScope: forScope, globalScope: e.currentScope}
		eval := evalVisitor.Visit(ctx.Sentencias())
		if eval == RETURNVOID {
			return RETURNVOID
		} else if eval == BREAK {
			break
		} else if eval == CONTINUE {
			continue
		}
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionForIdstmt(ctx *parser.FuncionForIdstmtContext) interface{} {
	fmt.Printf("Enter - For Statement\n")
	forScope := e.currentScope.CreateChildScope()
	id := ctx.TiposId(0).GetText()
	dataEnd := ctx.TiposId(1).GetText()
	start := 0
	fmt.Println(dataEnd)
	var end int
	if e.currentScope.FindVector(dataEnd) != nil {
		a := e.currentScope.FindVector(dataEnd)
		end = len(a.datos) - 1
		tipo := e.currentScope.FindTypeVector(dataEnd)
		for i := start; i <= end; i++ {
			loopVarValue := a.datos[i]
			if i == 0 {
				forScope.DeclareVariable(id, loopVarValue, tipo, false)
			} else {
				forScope.ReassignVariable(id, loopVarValue, tipo)
			}
			evalVisitor := &VisitorEvalue{currentScope: forScope, globalScope: e.currentScope}
			eval := evalVisitor.Visit(ctx.Sentencias())
			if eval == RETURNVOID {
				return RETURNVOID
			} else if eval == BREAK {
				break
			} else if eval == CONTINUE {
				continue
			}
		}
	}

	return VOID
}

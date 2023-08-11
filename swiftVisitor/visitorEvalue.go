package swiftVisitor

import (
	"fmt"
	"math"
	"modulo/parser"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

var out = ""

type VisitorEvalue struct {
	parser.BaseSwiftLanVisitor
	globalScope  *Scope
	currentScope *Scope
	returnValue  interface{}
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
		// fmt.Printf("> %s\n", n.(antlr.ParseTree).GetText())
		e.Visit(n.(antlr.ParseTree))
	}
	return VOID
}

func (e *VisitorEvalue) VisitSentencias(ctx *parser.SentenciasContext) interface{} {
	fmt.Printf("Entro - Sentencias\n")
	for _, StamentsCtx := range ctx.AllStatement() {
		e.Visit(StamentsCtx)
		if e.returnValue != nil {
			fmt.Println(e.returnValue)
			return e.returnValue
		}
	}
	return VOID
}

func (e *VisitorEvalue) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Printf("Entro VisitStatement\n")
	childCount := ctx.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := ctx.GetChild(i).(antlr.ParseTree)
		result := e.Visit(child)
		if _, isReturn := result.(*ReturnStatement); isReturn {
			// Si encontramos un ReturnStatement, detener el procesamiento y devolver el resultado
			return result
		}
	}
	return VOID

}
func (e *VisitorEvalue) VisitFuncionDeclaFunc(ctx *parser.FuncionDeclaFuncContext) interface{} {
	functionName := ctx.Id().GetText()
	fmt.Printf("Entro VisitFuncionDeclaFunc\n")
	tipe := ctx.TiposAsign().GetText()
	var params []*Param
	if paramList := ctx.ExprListFunc(); paramList != nil {
		for i := 0; i < len(paramList.AllId()); i += 2 {
			idExterior := paramList.Id(i).GetText()
			idInterior := paramList.Id(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
	}else{
		params=nil
	}

	body := ctx.SentenciasFunc()
	function := NewFunction(params, body, nil, tipe)
	e.currentScope.DeclareFunction(functionName, function)
	fmt.Printf("En FuncionFuncDec - Nombre Variable: %v Valor: %v\n", functionName, function.params)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionDeclaFunc2(ctx *parser.FuncionDeclaFunc2Context) interface{} {
	functionName := ctx.Id().GetText()
	fmt.Printf("Entro VisitFuncionDeclaFunc\n")
	var params []*Param
	if paramList := ctx.ExprListFunc(); paramList != nil {
		for i := 0; i < len(paramList.AllId()); i += 2 {
			idExterior := paramList.Id(i).GetText()
			idInterior := paramList.Id(i + 1).GetText()
			typ := paramList.TiposAsign(i / 2).GetText()
			params = append(params, &Param{
				idExterior: idExterior,
				idInterior: idInterior,
				typ:        typ,
			})
		}
	}

	body := ctx.SentenciasFunc()

	function := NewFunction(params, body, nil, VOID.String())
	e.currentScope.DeclareFunction(functionName, function)
	fmt.Printf("En FuncionFuncDec2 - Nombre Variable: %v Valor: %v\n", functionName, function.params)
	return VOID
}

func (e *VisitorEvalue) VisitSentenciasFunc(ctx *parser.SentenciasFuncContext) interface{} {
	fmt.Printf("Entro VisitSentenciasFunc\n")

	for _, stmtCtx := range ctx.AllStatement() {
		e.Visit(stmtCtx)
		// Verificar si hay un ReturnStatement y retornar si es necesario
		if e.returnValue != nil {
			fmt.Println(e.returnValue)
			return e.returnValue
		}
	}

	return VOID
}

func (e *VisitorEvalue) VisitFuncionReturnVal(ctx *parser.FuncionReturnValContext) interface{} {
	a := e.Visit(ctx.Expression()).(*SwiftValue)
	e.returnValue = a
	fmt.Println("*******************")
	fmt.Println(e.returnValue)
	if (a==nil){
		a=RETURNVOID
		e.returnValue = a
	}
	fmt.Println(e.returnValue)
	return a // Visit the expression node and return its value
}

func (e *VisitorEvalue) VisitExprCalFunc(ctx *parser.ExprCalFuncContext) interface{} {
	a := e.Visit(ctx.CallFuncstmt())
	return a.(*SwiftValue) // Visit the expression node and return its value
}

func (e *VisitorEvalue) VisitFuncionCallFunc(ctx *parser.FuncionCallFuncContext) interface{} {
	fmt.Printf("Entering VisitFuncionCallFunc\n")
	functionName := ctx.Id().GetText()
	if function := e.currentScope.FindFunction(functionName); function != nil {
		var args []*SwiftValue
		if exprList := ctx.ExprListCallFunc(); exprList != nil {
			for _, expr := range exprList.AllExpression() {
				argValue := e.Visit(expr).(*SwiftValue)
				args = append(args, argValue)
			}
			// Llamar a la función y obtener el valor de retorno
			
			fmt.Println("FUNCION")
			fmt.Println(ret)
			return ret.(*SwiftValue)
		} else {
			args=nil
			ret := function.invoke(e.currentScope, args)
			fmt.Println(ret)
		}
	} else {
		fmt.Printf("Function not found: %v\n", functionName)
	}

	return NULL
}

func (e *VisitorEvalue) VisitIfstmt(ctx *parser.IfstmtContext) interface{} {
    fmt.Printf("Entro VisitIf\n")

    // Evaluar la condición del if
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
        }
        return VOID
    }

    // Evaluar las condiciones de los else if
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

    // Evaluar el bloque else si existe
    if elseContext := ctx.Elsestmt(); elseContext != nil {
        e.Visit(elseContext.Sentencias())
        if e.returnValue != nil {
            return e.returnValue
        }
    }

    return VOID
}



func (e *VisitorEvalue) VisitFuncionAsigExp(ctx *parser.FuncionAsigExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	newVal := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	e.currentScope.DeclareVariable(name, newVal)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoNil(ctx *parser.FuncionAsigTipoNilContext) interface{} {
	fmt.Printf("Entro FuncionAsigExp\n")
	newVal := NULL
	name := ctx.Id().GetText()
	e.currentScope.DeclareVariable(name, newVal)
	fmt.Printf("En FuncionAsigExp - Nombre Variable: %v Valor: %v\n", name, newVal.value)
	return VOID
}

func (e *VisitorEvalue) VisitFuncionAsigTipoExp(ctx *parser.FuncionAsigTipoExpContext) interface{} {
	fmt.Printf("Entro FuncionAsigTipoExp\n")
	valVar := e.Visit(ctx.Expression()).(*SwiftValue)
	name := ctx.Id().GetText()
	fmt.Println(valVar.isString())
	if valVar.isString() {
		if ctx.TiposAsign().GetText() == "String" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isInt() {
		if ctx.TiposAsign().GetText() == "Int" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isDouble() {
		if ctx.TiposAsign().GetText() == "Float" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
	} else if valVar.isBool() {
		if ctx.TiposAsign().GetText() == "Bool" {
			e.currentScope.DeclareVariable(name, valVar)
			fmt.Printf("En FuncionAsigTipoExp - Nombre Variable: %v Valor: %v\n", name, valVar.value)
		} else {
			fmt.Println("Error")
		}
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

func (e *VisitorEvalue) VisitIdExpression(ctx *parser.IdExpressionContext) interface{} {
	id := ctx.Id().GetText()
	variable := e.currentScope.FindVariable(id)
	if variable == nil {
		fmt.Printf("Error: Variable '%s' no definida.\n", id)
		return INVALID
	}
	return variable
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
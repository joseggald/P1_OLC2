package swiftVisitor

import (
	"github.com/antlr4-go/antlr/v4"
)

type Param struct {
	idExterior string
	idInterior string
	typ        string
}

type Function struct {
	params    []*Param
	body      antlr.ParserRuleContext
	valReturn *SwiftValue
	tipo	string
}

func NewFunction(params []*Param, body antlr.ParserRuleContext, valReturn *SwiftValue,tipo string) *Function {
	return &Function{
		params:    params,
		body:      body,
		valReturn: valReturn,
		tipo:	tipo,
	}
}

func (f *Function) invoke(scope *Scope, args []*SwiftValue) interface{}{
	childScope := scope.CreateChildScope()
	

	// Asigna los valores de los argumentos a las variables
	for i, arg := range args {
		childScope.DeclareVariable(f.params[i].idInterior, arg,f.tipo,false)
	}

	evalVisitor := &VisitorEvalue{currentScope: childScope, globalScope: scope}
	evalVisitor.Visit(f.body)
	vala:=evalVisitor.returnValue
	
	// Retorna el valor de retorno configurado en el ámbito hijo
	return vala
}

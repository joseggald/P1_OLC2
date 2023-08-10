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
}

func NewFunction(params []*Param, body antlr.ParserRuleContext, valReturn *SwiftValue) *Function {
	return &Function{
		params:    params,
		body:      body,
		valReturn: valReturn,
	}
}

func (f *Function) invoke(scope *Scope, args []*SwiftValue) interface{}{
	childScope := scope.CreateChildScope()

	// Declara las variables utilizando los nombres e tipos de los parámetros
	for _, param := range f.params {
		childScope.DeclareVariable(param.idInterior, &SwiftValue{value: nil})
	}

	// Asigna los valores de los argumentos a las variables
	for i, arg := range args {
		childScope.DeclareVariable(f.params[i].idInterior, arg)
	}

	// Configura el valor de retorno en el ámbito hijo antes de la visita al cuerpo
	childScope.SetReturnValue(f.valReturn)
	f.valReturn=childScope.GetReturnValue()
	// Realiza la visita al cuerpo de la función
	evalVisitor := &VisitorEvalue{currentScope: childScope, globalScope: scope}
	evalVisitor.Visit(f.body)
	vala:=evalVisitor.returnValue
	
	// Retorna el valor de retorno configurado en el ámbito hijo
	return vala
}

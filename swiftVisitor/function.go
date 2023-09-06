package swiftVisitor

import (
	"github.com/antlr4-go/antlr/v4"
)

type Param struct {
	idExterior string
	idInterior string
	typ        string
	array      string
}

type Function struct {
	name	string
	params    []*Param
	body      antlr.ParserRuleContext
	tipo	string
}
type Argumento struct {
	nomVar		string
	variable    *SwiftValue
	vector		*Vector
	matriz		*Matrix
	matriz3d	*Matrix3D
}
func NewFunction(params []*Param, body antlr.ParserRuleContext, valReturn *SwiftValue,tipo string) *Function {
	return &Function{
		params:    params,
		body:      body,
		tipo:	tipo,
	}
}

func (f *Function) invoke(scope *Scope, args []*Argumento) interface{}{
	childScope := scope.CreateChildScope()

	for i, arg := range args {
		if f.params[i].array=="v" || f.params[i].array=="vN"{
			childScope.DeclareVector(f.params[i].idInterior, arg.vector)
		}else if f.params[i].array=="m" || f.params[i].array=="mN"{
			childScope.DeclareMatriz(f.params[i].idInterior, arg.matriz)
		}else if f.params[i].array=="m3" || f.params[i].array=="m3N"{
			childScope.DeclareMatriz3D(f.params[i].idInterior, arg.matriz3d)
		}else if f.params[i].array=="var" || f.params[i].array=="varN"{
			childScope.DeclareVariable(f.params[i].idInterior, arg.variable,f.params[i].typ,false)
		}
	}

	evalVisitor := &VisitorEvalue{currentScope: childScope, globalScope: scope}
	
	if len(f.name)>1{
		evalVisitor.funcsname=f.name
	}
	
	evalVisitor.Visit(f.body)
	for i, arg := range args {
		if f.params[i].array=="v"{
			cont:=childScope.FindVector(f.params[i].idInterior)
			childScope.setVector(arg.nomVar,cont.datos)
		}else if f.params[i].array=="m"{
			cont:=childScope.findMatriz(f.params[i].idInterior)
			childScope.setMatriz(arg.nomVar,cont)
		}else if f.params[i].array=="m3"{
			cont:=childScope.findMatriz3D(f.params[i].idInterior)
			childScope.setMatriz3D(arg.nomVar,cont)
		}else if f.params[i].array=="var"{
			cont:=childScope.FindVariable(f.params[i].idInterior)
			childScope.ReassignVariable(f.params[i].idInterior,cont,f.params[i].typ)
		}
	}
	vala:=evalVisitor.returnValue

	return vala
}

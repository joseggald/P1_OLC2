package swiftVisitor

type Struct struct {
	funciones		[]*FunctionStruct
	variables		[]*AtributoVariable
}

func NewStruct(funciones []*FunctionStruct, variables[]*AtributoVariable) *Struct {
	return &Struct{
		funciones:    funciones,
		variables:      variables,
	}
}



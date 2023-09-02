package swiftVisitor

type Struct struct {
	funciones		[]*FunctionStruct
	variables		[]*AtributoVariable
	structs			[]*Struct
}

func NewStruct(funciones []*FunctionStruct, variables[]*AtributoVariable,structs []*Struct) *Struct {
	return &Struct{
		funciones:    funciones,
		variables:      variables,
		structs:      structs,
	}
}



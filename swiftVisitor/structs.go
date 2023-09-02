package swiftVisitor

type Struct struct {
	nameVar			string
	funciones		[]*FunctionStruct
	variables		[]*AtributoVariable
	structs			[]*Struct
	constante		bool
}

func NewStruct(funciones []*FunctionStruct, variables[]*AtributoVariable,structs []*Struct) *Struct {
	return &Struct{
		funciones:    funciones,
		variables:      variables,
		structs:      structs,
	}
}



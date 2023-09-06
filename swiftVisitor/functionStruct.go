package swiftVisitor

type FunctionStruct struct {
	name 		string
	funcion		*Function
	mutanting	bool
}

func NewFunctionStruct(funcion *Function, mutanting bool,name string) *FunctionStruct {
	return &FunctionStruct{
		name: name,
		funcion:    funcion,
		mutanting:      mutanting,
	}
}
package swiftVisitor

type FunctionStruct struct {
	funcion		*Function
	mutanting	bool
}

func NewFunctionStruct(funcion *Function, mutanting bool) *FunctionStruct {
	return &FunctionStruct{
		funcion:    funcion,
		mutanting:      mutanting,
	}
}
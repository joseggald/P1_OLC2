package swiftVisitor

type AtributoVariable struct {
	name	string
	dato *SwiftValue
	tipo   string
	constante bool
	setdata   bool
}

func NewAtributoVariable(value *SwiftValue,tipo string,constante bool) *AtributoVariable {
	return &AtributoVariable {
		dato: value,
		tipo:	tipo,
		constante: constante,
	}
}
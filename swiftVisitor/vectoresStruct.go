package swiftVisitor

type VectorStruct struct {
	nameVar			string
	structs			[]*Struct
}

func NewVectorStruct(structs []*Struct) *VectorStruct {
	return &VectorStruct{
		structs:      structs,
	}
}

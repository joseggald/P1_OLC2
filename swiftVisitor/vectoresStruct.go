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
func (v *VectorStruct) apendVec(value *Struct) {
    v.structs = append(v.structs,value)
}
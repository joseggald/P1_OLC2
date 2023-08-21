package swiftVisitor

type Vector struct {
	datos  []*SwiftValue
	tipo   string
}
func NewVector(params []*SwiftValue,tipo string) *Vector {
	return &Vector{
		datos: params,
		tipo:	tipo,
	}
}

func (v *Vector) RemoveAt(index int) {
    if index < 0 || index >= len(v.datos) {
        // Manejar el índice fuera de rango aquí si es necesario
        return
    }  
    v.datos = append(v.datos[:index], v.datos[index+1:]...)
	for _, expr := range v.datos {
		println(expr.asInt())
	}
}

func (v *Vector) apendVec(value *SwiftValue) {
    v.datos = append(v.datos,value)
}

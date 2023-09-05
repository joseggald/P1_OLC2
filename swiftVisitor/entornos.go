package swiftVisitor

import (
	"fmt"
)

type Scope struct {
	parent             *Scope
	variables          map[string]*SwiftValue
	variablesAtributos map[string]*AtributoVariable
	functions          map[string]*Function
	vectores           map[string]*Vector
	matrices           map[string]*Matrix
	matrices3D         map[string]*Matrix3D
	structs            map[string]*Struct
	varsStruct         map[string]*Struct
}

func NewScope() *Scope {
	return &Scope{
		parent:             &Scope{},
		variables:          map[string]*SwiftValue{},
		variablesAtributos: map[string]*AtributoVariable{},
		functions:          map[string]*Function{},
		vectores:           map[string]*Vector{},
		matrices:           map[string]*Matrix{},
		matrices3D:         map[string]*Matrix3D{},
		structs:            map[string]*Struct{},
		varsStruct:         map[string]*Struct{},
	}
}

func (s *Scope) CreateChildScope() *Scope {
	childScope := NewScope()
	childScope.parent = s
	return childScope
}

func (s *Scope) DeclareVariable(name string, value *SwiftValue, tipo string, constante bool) {

	if value.isNumber() {
		if tipo == "Int" || tipo == "Float" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
		} else {
			println("error de tipado")
		}
	} else if value.isBool() {
		if tipo == "Bool" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
		} else {
			println("error de tipado")
		}
	} else if value.isString() {
		if tipo == "String" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
		} else {
			println("error de tipado")
		}
	} else if value.isChar() {
		if tipo == "Char" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
		} else {
			println("error de tipado")
		}
	}
}

func (s *Scope) DeclareVariableNil(name string, value *SwiftValue, tipo string, constante bool) {
	_, exists := s.variables[name]
	if exists {
		fmt.Println("error")
	} else {
		s.variables[name] = value
		s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
	}
}

func (s *Scope) ReassignVariable(name string, value *SwiftValue, tipo string) {
	varValue := s.variablesAtributos[name]
	_, exists := s.variables[name]
	if exists {
		if !varValue.constante {
			s.variablesAtributos[name].dato = value
			s.variables[name]=value
		}
	} else if s.parent != nil {
		s.parent.ReassignVariable(name, value, tipo)
	}

}

func (s *Scope) DeclareVector(name string, contenido *Vector) {
	_, exists := s.vectores[name]
	if exists {
		fmt.Println("error")
	} else {
		s.vectores[name] = contenido
	}
}

func (s *Scope) FindVector(name string) *Vector {
	vectorValue, exists := s.vectores[name]
	if exists {
		return vectorValue
	} else if s.parent != nil {
		return s.parent.FindVector(name)
	}
	return nil
}

func (s *Scope) FindTypeVector(name string) string {
	cont, exists := s.vectores[name]
	if exists {
		return cont.tipo
	} else if s.parent != nil {
		return s.parent.FindTypeVector(name)
	}
	return "nil"
}

func (s *Scope) AddVector(name string, dato *SwiftValue, tipo string) {
	cont, exists := s.vectores[name]
	if exists {
		if cont.tipo == tipo {
			cont.apendVec(dato)
		} else {
			fmt.Println("error no es el mismo tipo del vector")
		}
	} else if s.parent != nil {
		s.parent.AddVector(name, dato, tipo)
	}
}

func (s *Scope) setVector(name string, value []*SwiftValue) {
	cont, exists := s.vectores[name]
	if exists {
		cont.setearDatos(value)
	} else if s.parent != nil {
		s.parent.setVector(name,value)
	}
}

func (s *Scope) setMatriz(name string, value *Matrix) {
	cont, exists := s.matrices[name]
	if exists {
		cont.setearDatos(value)
	} else if s.parent != nil {
		s.parent.setMatriz(name,value)
	}
}

func (s *Scope) setMatriz3D(name string, value *Matrix3D) {
	cont, exists := s.matrices3D[name]
	if exists {
		cont.setearDatos(value)
	} else if s.parent != nil {
		s.parent.setMatriz3D(name,value)
	}
}

func (s *Scope) DelVector(name string, pos int) {
	cont, exists := s.vectores[name]
	if exists {
		cont.RemoveAt(pos)
	} else if s.parent != nil {
		s.parent.DelVector(name, pos)
	}
}

func (s *Scope) ReasignVector(name string, pos int, value *SwiftValue) {
	cont, exists := s.vectores[name]
	if exists {
		cont.datos[pos] = value
	} else if s.parent != nil {
		s.parent.ReasignVector(name, pos, value)
	}
}
func (s *Scope) FindVariable(name string) *SwiftValue {
	val, exists := s.variables[name]
	if exists {
		return val
	} else if s.parent != nil {
		return s.parent.FindVariable(name)
	}
	return VOID
}

func (s *Scope) DeclareFunction(name string, value *Function) {
	_, exists := s.functions[name]
	if exists {
		fmt.Println("error")
	} else {
		s.functions[name] = value
	}

}

func (s *Scope) FindFunction(name string) *Function {
	funcValue, exists := s.functions[name]
	if exists {
		return funcValue
	} else if s.parent != nil {
		return s.parent.FindFunction(name)
	}
	return nil
}

func (s *Scope) DeclareMatriz(name string, value *Matrix) {
	_, exists := s.matrices[name]
	if exists {
		fmt.Println("error")
	} else {
		s.matrices[name] = value
	}
}

func (s *Scope) DeclareMatriz3D(name string, value *Matrix3D) {
	_, exists := s.matrices3D[name]
	if exists {
		fmt.Println("error")
	} else {
		s.matrices3D[name] = value
	}
}

func (s *Scope) ReasignMatriz(name string, row, col int, value *SwiftValue) {
	_, exists := s.matrices[name]
	if exists {
		s.matrices[name].SetValue(row, col, value)
	} else if s.parent != nil {
		s.parent.ReasignMatriz(name, row, col, value)
	}
}

func (s *Scope) ReasignMatriz3D(name string, row, col, depth int, value *SwiftValue) {
	_, exists := s.matrices3D[name]
	if exists {
		s.matrices3D[name].SetValue(row, col, depth, value)
	} else if s.parent != nil {
		s.parent.ReasignMatriz3D(name, row, col, depth, value)
	}
}

func (s *Scope) findMatriz(name string) *Matrix {
	cont, exists := s.matrices[name]
	if exists {
		return cont
	} else if s.parent != nil {
		return s.parent.findMatriz(name)
	}
	return nil
}

func (s *Scope) findMatriz3D(name string) *Matrix3D {
	cont, exists := s.matrices3D[name]
	if exists {
		return cont
	} else if s.parent != nil {
		return s.parent.findMatriz3D(name)
	}
	return nil
}

func (s *Scope) DeclareStruct(name string, value *Struct) {
	_, exists := s.structs[name]
	if exists {
		fmt.Println("error")
	} else {
		s.structs[name] = value
	}
}

func (s *Scope) findStruct(name string) *Struct {
	cont, exists := s.structs[name]
	if exists {
		return cont
	} else if s.parent != nil {
		return s.parent.findStruct(name)
	}
	return nil
}

func (s *Scope) DeclareVarStruct(name string, value *Struct) {
	_, exists := s.varsStruct[name]
	if exists {
		fmt.Println("error")
	} else {
		s.varsStruct[name] = value
		fmt.Println(s.varsStruct[name].variables[0].name)
		fmt.Println(s.varsStruct[name].variables[0].dato.value)
	}
}

func (s *Scope) findVarStruct(name string) *Struct {
	cont, exists := s.varsStruct[name]
	if exists {
		return cont
	} else if s.parent != nil {
		return s.parent.findStruct(name)
	}
	return nil
}

func (s *Scope) reasigVarStruct(name string, value *SwiftValue, atributo string) {
	_, exists := s.varsStruct[name]
	if exists {
		for _, index := range s.varsStruct[name].variables {
			if index.name == atributo {
				fmt.Println("llego2")
				index.dato = value
			}
		}
	} else if s.parent != nil {
		s.parent.reasigVarStruct(name, value, atributo)
	}
}

func (s *Scope) reasigVarStruct2(name string, value *SwiftValue, atributo, atributo2 string) {
	_, exists := s.varsStruct[name]
	if exists {
		for _, index := range s.varsStruct[name].structs {
			if index.nameVar == atributo {
				for _, vars := range index.variables {
					if vars.name == atributo2 {
						vars.dato = value
					}
				}
			}
		}
	} else if s.parent != nil {
		s.parent.reasigVarStruct(name, value, atributo)
	}
}

func (s *Scope) verifyStructVar(name string, atributo string) *Struct {
	_, exists := s.varsStruct[name]
	if exists {
		for _, index := range s.varsStruct[name].structs {
			if index.nameVar == atributo {
				return index
			}
		}
	} else if s.parent != nil {
		s.parent.verifyStructVar(name, atributo)
	}
	return nil
}

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
	isVector           bool
	isFunction         bool
	constante          bool
}

func NewScope() *Scope {
	return &Scope{
		parent:             &Scope{},
		variables:          map[string]*SwiftValue{},
		variablesAtributos: map[string]*AtributoVariable{},
		functions:          map[string]*Function{},
		vectores:           map[string]*Vector{},
		isFunction:         false,
		isVector:           false,
		constante:          false,
	}
}

func (s *Scope) CreateChildScope() *Scope {
	childScope := NewScope()
	childScope.parent = s
	return childScope
}

func (s *Scope) DeclareVariable(name string, value *SwiftValue, tipo string, constante bool) {
	if value.isInt() {
		if tipo == "Int" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isBool() {
		if tipo == "Bool" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isNumber() {
		if tipo == "Float" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isString() {
		if tipo == "String" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isChar() {
		if tipo == "Char" {
			s.variables[name] = value
			s.variablesAtributos[name] = NewAtributoVariable(value, tipo, constante)
			return
		} else {
			println("error de tipado")
		}
	}
}
func (s *Scope) DeclareVariableNil(name string, value *SwiftValue, tipo string, constante bool) {
	s.constante = constante
	s.variables[name] = value

}

func (s *Scope) ReassignVariable(name string, value *SwiftValue, tipo string) {
	varValue := s.variablesAtributos[name]
	if !varValue.constante {
		if _, exists := s.variables[name]; exists {
			if tipo == varValue.tipo {
				s.variables[name] = value
			} else if tipo == "Int" && varValue.tipo == "Float" {
				s.variables[name] = value
			} else {
				println("error de reasignación tipado")
			}
		} else if s.parent != nil {
			s.parent.ReassignVariable(name, value, tipo)
		}
	} else {
		println("error es constante no puede reasignar")
	}

}

func (s *Scope) DeclareVector(name string, contenido *Vector) {
	s.vectores[name] = contenido
	s.isVector = true
	println(contenido.datos)
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
		if s.isVector {
			if cont.tipo == tipo {
				cont.apendVec(dato)
			} else {
				fmt.Println("error no es el mismo tipo del vector")
			}
		} else {
			fmt.Println("error no es vector")
		}

	} else if s.parent != nil {
		s.parent.AddVector(name, dato, tipo)
	}
}

func (s *Scope) DelVector(name string, pos int) {
	cont, exists := s.vectores[name]
	if exists {
		if s.isVector {
			cont.RemoveAt(pos)
		}
	} else if s.parent != nil {
		s.parent.DelVector(name, pos)
	}
}

func (s *Scope) ReasignVector(name string, pos int, value *SwiftValue) {
	cont, exists := s.vectores[name]
	if exists {
		if s.isVector {
			cont.datos[pos] = value
		}
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
	s.functions[name] = value
}

func (s *Scope) FindFunction(name string) *Function {
	funcValue, exists := s.functions[name]
	if exists {
		return funcValue
	} else if !s.isFunction && s.parent != nil {
		return s.parent.FindFunction(name)
	}
	return nil
}

package swiftVisitor

import (
	"fmt"
)

type Scope struct {
	parent     *Scope
	variables  map[string]*SwiftValue
	functions  map[string]*Function
	contenido  []interface{}
	isVector   bool
	isFunction bool
	constante  bool
	tipo       string
}

func NewScope() *Scope {
	return &Scope{
		parent:     &Scope{},
		variables:  map[string]*SwiftValue{},
		functions:  map[string]*Function{},
		isFunction: false,
		isVector: false,
		contenido:  make([]interface{}, 0),
		constante:  false,
		tipo:       "",
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
			s.tipo = tipo
			s.constante = constante
			s.variables[name] = value
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isBool() {
		if tipo == "Bool" {
			s.tipo = tipo
			s.constante = constante
			s.variables[name] = value
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isNumber() {
		if tipo == "Float" {
			s.tipo = tipo
			s.constante = constante
			s.variables[name] = value
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isString() {
		if tipo == "String" {
			s.tipo = tipo
			s.constante = constante
			s.variables[name] = value
			return
		} else {
			println("error de tipado")
		}
	}
	if value.isChar() {
		if tipo == "Char" {
			s.tipo = tipo
			s.constante = constante
			s.variables[name] = value
			return
		} else {
			println("error de tipado")
		}
	}
	println(s.tipo)
}
func (s *Scope) DeclareVariableNil(name string, value *SwiftValue, tipo string, constante bool) {
	s.tipo = tipo
	s.constante = constante
	s.variables[name] = value
	
}

func (s *Scope) ReassignVariable(name string, value *SwiftValue, tipo string) {
	if !s.constante {
		if _, exists := s.variables[name]; exists {
			if tipo == s.tipo {
				s.variables[name] = value
			}else if tipo=="Int" && s.tipo=="Float"{
				s.variables[name] = value
			}else {
				println("error de reasignación tipado")
			}
		} else if s.parent != nil {
			s.parent.ReassignVariable(name, value, tipo)
		}
	} else {
		println("error es constante no puede reasignar")
	}

}

func (s *Scope) DeclareVector(name string, value *SwiftValue, tipo string, contenido []interface{}) {
	s.tipo = tipo
	s.constante = false
	s.variables[name] = value
	s.contenido = contenido
	s.isVector=true
	fmt.Println(contenido)
}

func (s *Scope) FindVector(name string) []interface{} {
	_, exists := s.variables[name]
	if exists {
		return s.contenido
	} else if s.parent != nil {
		return s.parent.FindVector(name)
	}
	return nil
}

func (s *Scope) FindTypeVector(name string) string {
	_, exists := s.variables[name]
	if exists {
		return s.tipo
	} else if s.parent != nil {
		return s.parent.FindTypeVector(name)
	}
	return "nil"
}



func (s *Scope) AddVector(name string, dato interface{},tipo string){
	_, exists := s.variables[name]
	if exists {
		if s.isVector{
			if s.tipo==tipo{
				s.contenido = append(s.contenido, dato)
				fmt.Println(s.contenido)
			}else{
				fmt.Println("error no es el mismo tipo del vector")
			}
		}else{
			fmt.Println("error no es vector")
		}

	} else if s.parent != nil {
		s.parent.AddVector(name,dato,tipo)
	}
}

func (s *Scope) DelVector(name string, pos int){
	_, exists := s.variables[name]
	if exists {
		if s.isVector{
			indexToRemove := pos
			if indexToRemove >= 0 && indexToRemove < len(s.contenido) {
				// Crear un nuevo slice que excluya el elemento en la posición indexToRemove
				s.contenido = append(s.contenido[:indexToRemove], s.contenido[indexToRemove+1:]...)
				fmt.Println("Slice después de eliminar:", s.contenido)
			} else {
				fmt.Println("Índice inválido para eliminación")
			}

		}
	} else if s.parent != nil {
		s.parent.DelVector(name,pos)
	}
}
func (s *Scope) ReasignVector(name string, pos int,value interface{}){
	_, exists := s.variables[name]
	if exists {
		if s.isVector{
			s.contenido[pos]=value
		}
	} else if s.parent != nil {
		s.parent.DelVector(name,pos)
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

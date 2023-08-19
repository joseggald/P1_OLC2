package swiftVisitor

type Scope struct {
	parent      *Scope
	variables   map[string]*SwiftValue
	functions   map[string]*Function
	isFunction  bool
	constante bool
	tipo string
}

func NewScope() *Scope {
	return &Scope{
		parent:      &Scope{},
		variables:   map[string]*SwiftValue{},
		functions:   map[string]*Function{},
		isFunction:  false,
		constante: false,
		tipo: "",
	}
}

func (s *Scope) CreateChildScope() *Scope {
	childScope := NewScope()
	childScope.parent = s
	return childScope
}

func (s *Scope) DeclareVariable(name string, value *SwiftValue,tipo string,constante bool) {
	if value.isInt(){
		if tipo=="Int"{
			s.tipo=tipo
			s.constante=constante
			s.variables[name] = value
		}else{
			println("error de tipado")
		}
	}
	if value.isBool(){
		if tipo=="Bool"{
			s.tipo=tipo
			s.constante=constante
			s.variables[name] = value
		}else{
			println("error de tipado")
		}
	}
	if value.isNumber(){
		if tipo=="Float"{
			s.tipo=tipo
			s.constante=constante
			s.variables[name] = value
		}else{
			println("error de tipado")
		}
	}
	if value.isString(){
		if tipo=="String"{
			s.tipo=tipo
			s.constante=constante
			s.variables[name] = value
		}else{
			println("error de tipado")
		}
	}
	println(s.tipo)
}
func (s *Scope) DeclareVariableNil(name string, value *SwiftValue,tipo string,constante bool) {
	s.tipo=tipo
	s.constante=constante
	s.variables[name] = value
}
func (s *Scope) ReassignVariable(name string, value *SwiftValue,tipo string){
	if _, exists := s.variables[name]; exists {
		println(s.tipo)
		println(tipo)
		if tipo==s.tipo{
			s.variables[name] = value
		}else{
			println("error de reasignación tipado")
		}
	} else if s.parent != nil {
		s.parent.ReassignVariable(name, value,tipo)
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

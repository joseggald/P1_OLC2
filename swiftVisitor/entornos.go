package swiftVisitor

type Scope struct {
	parent     *Scope
	variables  map[string]*SwiftValue
	functions  map[string]*Function
	isFunction bool
	returnValue  *SwiftValue 
}

func NewScope() *Scope {
	return &Scope{
		parent:     &Scope{},
		variables:  map[string]*SwiftValue{},
		functions:  map[string]*Function{},
		isFunction: false,
		returnValue:  &SwiftValue{},
	}
}

func (s *Scope) CreateChildScope() *Scope {
	childScope := NewScope()
	childScope.parent = s
	return childScope
}

func (s *Scope) DeclareVariable(name string, value *SwiftValue) {
	s.variables[name] = value
}

func (s *Scope) FindVariable(name string) *SwiftValue {
	val, exists := s.variables[name]
	if exists {
		return val
	} else if s.parent != nil {
		return s.parent.FindVariable(name)
	}
	return nil
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

func (s *Scope) SetReturnValue(value *SwiftValue) {
	s.returnValue = value
}

func (s *Scope) GetReturnValue() *SwiftValue {
	return s.returnValue
}
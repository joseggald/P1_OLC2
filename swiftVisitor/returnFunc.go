package swiftVisitor

type ReturnStatement struct {
	value *SwiftValue
}

func NewReturnStatement(value *SwiftValue) *ReturnStatement {
	return &ReturnStatement{value: value}
}

func (r *ReturnStatement) GetValue() *SwiftValue {
	return r.value
}
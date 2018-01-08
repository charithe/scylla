package symbols

import (
	"errors"
	"strconv"
)

var ErrInvalidScalarType = errors.New("Invalid scalar type")

type ValueType int

const (
	VARIABLE_VALUE ValueType = iota
	INT_VALUE
	FLOAT_VALUE
	STRING_VALUE
	BOOLEAN_VALUE
	NULL_VALUE
	ENUM_VALUE
	LIST_VALUE
	OBJECT_VALUE
)

// Value is an interface representing various GraphQL value types
type Value interface {
	Type() ValueType
	Value() interface{}
}

func CreateScalarValue(tpe ValueType, data string) (Value, error) {
	switch tpe {
	case VARIABLE_VALUE:
		return &VariableValue{Name: data}, nil
	case INT_VALUE:
		intVal, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil, err
		}
		return &IntValue{Val: intVal}, nil
	case FLOAT_VALUE:
		floatVal, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return nil, err
		}
		return &FloatValue{Val: floatVal}, nil
	case STRING_VALUE:
		return &StringValue{Val: data}, nil
	case BOOLEAN_VALUE:
		boolVal, err := strconv.ParseBool(data)
		if err != nil {
			return nil, err
		}
		return &BooleanValue{Val: boolVal}, nil
	case NULL_VALUE:
		return &NullValue{}, nil
	case ENUM_VALUE:
		return &EnumValue{Val: data}, nil
	default:
		return nil, ErrInvalidScalarType
	}
}

// VariableValue holds a GraphQL variable
type VariableValue struct {
	Name string
}

func (v *VariableValue) Type() ValueType {
	return VARIABLE_VALUE
}

func (v *VariableValue) Value() interface{} {
	return v.Name
}

// IntValue holds a GraphQL integer value
type IntValue struct {
	Val int64
}

func (i *IntValue) Type() ValueType {
	return INT_VALUE
}

func (i *IntValue) Value() interface{} {
	return i.Val
}

// FloatValue holds a GraphQL float value
type FloatValue struct {
	Val float64
}

func (f *FloatValue) Type() ValueType {
	return FLOAT_VALUE
}

func (f *FloatValue) Value() interface{} {
	return f.Val
}

// StringValue holds a GraphQL string value
type StringValue struct {
	Val string
}

func (s *StringValue) Type() ValueType {
	return STRING_VALUE
}

func (s *StringValue) Value() interface{} {
	return s.Val
}

// BooleanValue holds a GraphQL boolean value
type BooleanValue struct {
	Val bool
}

func (b *BooleanValue) Type() ValueType {
	return BOOLEAN_VALUE
}

func (b *BooleanValue) Value() interface{} {
	return b.Val
}

// NullValue holds a GraphQL null value
type NullValue struct{}

func (n *NullValue) Type() ValueType {
	return NULL_VALUE
}

func (n *NullValue) Value() interface{} {
	return nil
}

// EnumValue holds a GraphQL enum value
type EnumValue struct {
	Val string
}

func (e *EnumValue) Type() ValueType {
	return ENUM_VALUE
}

func (e *EnumValue) Value() interface{} {
	return e.Val
}

// ListValue holds a GraphQL list value
type ListValue struct {
	Items []Value
}

func (l *ListValue) Type() ValueType {
	return LIST_VALUE
}

func (l *ListValue) Value() interface{} {
	return l.Items
}

// ObjectValue holds a GraphQL object value
type ObjectValue struct {
	Items map[string]Value
}

func (o *ObjectValue) Type() ValueType {
	return OBJECT_VALUE
}

func (o *ObjectValue) Value() interface{} {
	return o.Items
}

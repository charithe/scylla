package symbols

import (
	"encoding/json"
	"fmt"
)

type SymbolType int

const (
	ALIAS SymbolType = iota
	NAME
	FIELD
	SELECTION_SET
	DOCUMENT
	KEY_VALUE
	KEY
	SCALAR
	ARGUMENTS
	LIST
	OBJECT
	DIRECTIVE
)

type Symbol interface {
	Type() SymbolType
	Add(tpe SymbolType, sym interface{}) error
}

// Stringify returns the string representation of the given symbol
func Stringify(s Symbol) string {
	str, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Sprintf("ERR %+v", s)
	}
	return string(str)
}

// Document represents a GraphQL document
type Document struct {
	SelectionSets []*SelectionSet
}

func (d *Document) Type() SymbolType {
	return DOCUMENT
}

func (d *Document) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case SELECTION_SET:
		if ss, ok := sym.(*SelectionSet); ok {
			d.SelectionSets = append(d.SelectionSets, ss)
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// SelectionSet represents a GraphQL selection set
type SelectionSet struct {
	Fields []*Field
}

func (ss *SelectionSet) Type() SymbolType {
	return SELECTION_SET
}

func (ss *SelectionSet) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case FIELD:
		if f, ok := sym.(*Field); ok {
			ss.Fields = append(ss.Fields, f)
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// Field represents a GraphQL field
type Field struct {
	Alias      string
	Name       string
	Args       []*KeyValue
	Directives []*Directive
	Selections *SelectionSet
}

func (f *Field) Type() SymbolType {
	return FIELD
}

func (f *Field) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case ALIAS:
		if a, ok := sym.(string); ok {
			f.Alias = a
			return nil
		}
		return ErrIllegalSymbolValue
	case NAME:
		if n, ok := sym.(string); ok {
			f.Name = n
			return nil
		}
		return ErrIllegalSymbolValue
	case ARGUMENTS:
		if args, ok := sym.(*Arguments); ok {
			f.Args = args.KeyValues
			return nil
		}
		return ErrIllegalSymbolValue
	case DIRECTIVE:
		if dir, ok := sym.(*Directive); ok {
			f.Directives = append(f.Directives, dir)
			return nil
		}
		return ErrIllegalSymbolValue
	case SELECTION_SET:
		if ss, ok := sym.(*SelectionSet); ok {
			f.Selections = ss
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// Arguments is a container for a list of arguments
type Arguments struct {
	KeyValues []*KeyValue
}

func (a *Arguments) Type() SymbolType {
	return ARGUMENTS
}

func (a *Arguments) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case KEY_VALUE:
		if kv, ok := sym.(*KeyValue); ok {
			a.KeyValues = append(a.KeyValues, kv)
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// Scalar is a container for a GraphQL scalar value type
type Scalar struct {
	Value Value
}

func (v *Scalar) Type() SymbolType {
	return SCALAR
}

func (v *Scalar) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case SCALAR:
		if val, ok := sym.(Value); ok {
			v.Value = val
			return nil
		}
		return ErrIllegalSymbolValue

	default:
		return ErrIllegalSymbolType
	}
}

// KeyValue is a container for a key-value pair
type KeyValue struct {
	Key   string
	Value Value
}

func (kv *KeyValue) Type() SymbolType {
	return KEY_VALUE
}

func (kv *KeyValue) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case KEY:
		if k, ok := sym.(string); ok {
			kv.Key = k
			return nil
		}
		return ErrIllegalSymbolValue
	case SCALAR:
		if v, ok := sym.(*Scalar); ok {
			kv.Value = v.Value
			return nil
		}
		return ErrIllegalSymbolValue
	case LIST:
		if l, ok := sym.(*List); ok {
			kv.Value = &ListValue{Items: l.Items}
			return nil
		}
		return ErrIllegalSymbolValue
	case OBJECT:
		if o, ok := sym.(*Object); ok {
			kv.Value = &ObjectValue{Items: o.Items}
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// List is a container for a list value
type List struct {
	Items []Value
}

func (l *List) Type() SymbolType {
	return LIST
}

func (l *List) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case SCALAR:
		if v, ok := sym.(*Scalar); ok {
			l.Items = append(l.Items, v.Value)
			return nil
		}
		return ErrIllegalSymbolValue
	case LIST:
		if lst, ok := sym.(*List); ok {
			l.Items = append(l.Items, &ListValue{Items: lst.Items})
			return nil
		}
		return ErrIllegalSymbolValue
	case OBJECT:
		if o, ok := sym.(*Object); ok {
			l.Items = append(l.Items, &ObjectValue{Items: o.Items})
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// Object is a containter for an object value
type Object struct {
	Items map[string]Value
}

func (o *Object) Type() SymbolType {
	return OBJECT
}

func (o *Object) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case KEY_VALUE:
		if kv, ok := sym.(*KeyValue); ok {
			if o.Items == nil {
				o.Items = make(map[string]Value)
			}
			o.Items[kv.Key] = kv.Value
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

// Directive is a container for a directive
type Directive struct {
	Name string
	Args []*KeyValue
}

func (d *Directive) Type() SymbolType {
	return DIRECTIVE
}

func (d *Directive) Add(tpe SymbolType, sym interface{}) error {
	switch tpe {
	case NAME:
		if n, ok := sym.(string); ok {
			d.Name = n
			return nil
		}
		return ErrIllegalSymbolValue
	case ARGUMENTS:
		if args, ok := sym.(*Arguments); ok {
			d.Args = args.KeyValues
			return nil
		}
		return ErrIllegalSymbolValue
	default:
		return ErrIllegalSymbolType
	}
}

package symbols

import (
	"bytes"
	"errors"
)

var (
	ErrStackEmpty  = errors.New("Empty stack")
	ErrCannotMerge = errors.New("Cannot merge symbols")
)

type Stack struct {
	items []Symbol
	top   int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]Symbol, 0, 8),
		top:   -1,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) PushNewSymbol(tpe SymbolType) {
	switch tpe {
	case FIELD:
		s.Push(&Field{})
	case SELECTION_SET:
		s.Push(&SelectionSet{})
	case DOCUMENT:
		s.Push(&Document{})
	case ARGUMENTS:
		s.Push(&Arguments{})
	case SCALAR:
		s.Push(&Scalar{})
	case KEY_VALUE:
		s.Push(&KeyValue{})
	case LIST:
		s.Push(&List{})
	case OBJECT:
		s.Push(&Object{})
	case DIRECTIVE:
		s.Push(&Directive{})
	default:
		panic("Unhandled symbol type in PushNewSymbol")
	}
}

func (s *Stack) Push(sym Symbol) {
	s.top += 1
	if s.top >= len(s.items) {
		s.items = append(s.items, sym)
	} else {
		s.items[s.top] = sym
	}
}

func (s *Stack) Pop() (Symbol, error) {
	if s.IsEmpty() {
		return nil, ErrStackEmpty
	}

	sym := s.items[s.top]
	s.items[s.top] = nil
	s.top -= 1
	return sym, nil
}

func (s *Stack) Peek() (Symbol, error) {
	if s.IsEmpty() {
		return nil, ErrStackEmpty
	}

	return s.items[s.top], nil
}

func (s *Stack) MergeDown() error {
	if s.top < 1 {
		return ErrCannotMerge
	}

	child, _ := s.Pop()
	s.items[s.top].Add(child.Type(), child)
	return nil
}

func (s *Stack) String() string {
	var buf bytes.Buffer
	for i := s.top; i > -1; i-- {
		(&buf).WriteString(Stringify(s.items[i]))
		(&buf).WriteString("\n-----\n")
	}

	return (&buf).String()
}

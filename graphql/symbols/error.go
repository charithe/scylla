package symbols

import (
	"errors"
	"fmt"
)

var (
	ErrParseFailure       = errors.New("Parse failure")
	ErrUnexpectedEOF      = errors.New("Unexpected EOF")
	ErrIllegalSymbolType  = errors.New("Illegal symbol")
	ErrIllegalSymbolValue = errors.New("Illegal symbol value")
)

type ParseError struct {
	Cause    error
	Consumed []byte
	Pos      int
}

func (pe *ParseError) Error() string {
	var context string
	consumedLen := len(pe.Consumed)
	if consumedLen > 25 {
		context = "... " + string(pe.Consumed[consumedLen-25:])
	} else {
		context = string(pe.Consumed)
	}

	return fmt.Sprintf("%+v [Pos=%d Context='%s']", pe.Cause, pe.Pos, context)
}

package internal

import "fmt"
import "github.com/charithe/scylla/graphql/symbols"

%%{
    machine graphql;
    write data;
    access rlp.;
    variable p rlp.p;
    variable pe rlp.pe;
}%%


type RagelParser struct {
    data []byte
    p, pe, cs, top, te, ts, act int
    stack [32]int
    opStart, opEnd int
    tmpPtr int
    name string
    symStack *symbols.Stack
    Err error
}

func NewRagelParser(data []byte) *RagelParser {
    rlp := &RagelParser{
        data: data,
        pe: len(data),
        symStack: symbols.NewStack(),
    }

    %% write init;

    return rlp
}

func (rlp *RagelParser) Start() (symbols.Symbol, error) {
    data := rlp.data
    eof := len(rlp.data) 
    rlp.documentStart();

     %%{
        action arguments_start { 
            rlp.startSymbol(symbols.ARGUMENTS); 
            fcall arguments; 
        }

        action end_symbol { 
            rlp.endSymbol(false); 
            fret; 
        }
        
        action end_symbol_with_rewind { 
            rlp.endSymbol(true); 
            fret; 
        }

        action illegal_character { fmt.Printf("Illegal character\n"); fbreak; }
        
        action key_value_start { 
            rlp.startSymbol(symbols.KEY_VALUE); 
            fcall key_value; 
        }
        
        action name_end { 
            rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
        }
        
        action name_start { 
            rlp.tmpPtr = rlp.p; 
        }

        action rewind {
            rlp.rewind()
        }
        
        action selection_set_start { 
            rlp.startSymbol(symbols.SELECTION_SET); 
            fcall selection_set; 
        }
        

        name_begin = [_a-zA-Z];

        name = ([_a-zA-Z][_a-zA-Z0-9]*) >name_start %name_end;

        ws = space+;

        alias = (name ws? ':' ws?);

        key = alias;

        list_sep = ','+;

        variable_value = '$' name;

        int_value = '-'? (0 | [1-9][0-9]*);

        fraction = '.' digit+;

        exponent = [eE] [+\-]? digit+;

        float_value = int_value (fraction | exponent | fraction exponent);

        string_value := |*
            '\\"';
            '"' => { rlp.addScalarValueToSymbol(symbols.STRING_VALUE, string(rlp.data[rlp.tmpPtr:rlp.p])); fgoto after_scalar_value; };
            (any - [\r\n]);
        *|;

        special_names = ('true' | 'false' | 'null');

        boolean_value = ('true' | 'false');

        list_value := |*
            ws;
            list_sep;
            ']'         => end_symbol;
            any         => { rlp.rewind(); fcall value; };
        *|;

        object_value := |*
            ws;
            list_sep;
            '}'                => end_symbol;
            name_begin %rewind => key_value_start;
        *|;

        after_scalar_value := |*
            ws;
            list_sep => end_symbol;
            any      => end_symbol_with_rewind;
        *|;

        value := |*
            boolean_value          => { rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.BOOLEAN_VALUE); fgoto after_scalar_value; };
            'null'                 => { rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.NULL_VALUE); fgoto after_scalar_value; };
            (name - special_names) => { rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.ENUM_VALUE);  fgoto after_scalar_value; };
            variable_value         => { rlp.startSymbol(symbols.SCALAR); rlp.addScalarValueToSymbol(symbols.VARIABLE_VALUE, rlp.name);  fgoto after_scalar_value; };
            int_value              => { rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.INT_VALUE);  fgoto after_scalar_value; };
            float_value            => { rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.FLOAT_VALUE);  fgoto after_scalar_value; };
            '"'                    => { rlp.startSymbol(symbols.SCALAR); rlp.tmpPtr = rlp.p+1; fgoto string_value; };
            '['                    => { rlp.startSymbol(symbols.LIST); fgoto list_value; };
            '{'                    => { rlp.startSymbol(symbols.OBJECT); fgoto object_value; };
        *|;

        key_value := |*
            key          => { rlp.addNameToSymbolAs(symbols.KEY); fcall value; };
            any          => end_symbol_with_rewind;
        *|;

        arguments := |*
            ws;
            list_sep;
            ')'                => end_symbol;
            name_begin %rewind => key_value_start;
        *|;

        inside_directive := |*
            ws;
            '('      => arguments_start;
            list_sep => end_symbol;
            any      => end_symbol_with_rewind;
        *|;

        directives := |*
            ws;
            '@' name => { rlp.startSymbol(symbols.DIRECTIVE); rlp.addNameToSymbolAs(symbols.NAME); fcall inside_directive; };
            any      => { rlp.rewind(); fret; };
        *|;

        inside_field := |*
            ws;
            '('      => arguments_start;
            '@'      => { rlp.rewind(); fcall directives; };
            '{'      => selection_set_start;
            list_sep => end_symbol;
            any      => end_symbol_with_rewind;
        *|;

        field := |*
            alias    => { rlp.addNameToSymbolAs(symbols.ALIAS) };
            name     => { rlp.addNameToSymbolAs(symbols.NAME); fgoto inside_field; };
            list_sep => end_symbol;
            any      => end_symbol_with_rewind;
        *|;

        selection_set := |*
            ws;
            '}'                => end_symbol;
            name_begin %rewind => { rlp.startSymbol(symbols.FIELD); fcall field; };
        *|;
        
        main := |*
            ws;
            '{'        => selection_set_start;
        *|;

     
         write exec;
     }%%

     if rlp.cs < graphql_first_final  {
        if rlp.Err == nil {
            rlp.error(symbols.ErrParseFailure)
        }

        return nil, rlp.Err
     }

     if rlp.symStack.Size() > 1 {
        return nil, symbols.ErrUnexpectedEOF
     }

    return rlp.symStack.Pop()
}


func (rlp *RagelParser) rewind() {
    rlp.p -= 1
}

func (rlp *RagelParser) startSymbol(tpe symbols.SymbolType) {
    rlp.symStack.PushNewSymbol(tpe)
}

func (rlp *RagelParser) endSymbol(rewind bool) {
    if err := rlp.symStack.MergeDown(); err != nil {
        rlp.error(err)
        return
    }

    if(rewind) {
        rlp.rewind()
    }
}

func (rlp *RagelParser) addNameToSymbolAs(tpe symbols.SymbolType) {
    parent, err := rlp.symStack.Peek()
    if err != nil {
        rlp.error(err)
        return
    }

    parent.Add(tpe, rlp.name)
}

func (rlp *RagelParser) createScalarValueFromMatch(tpe symbols.ValueType) {
    rlp.addScalarValueToSymbol(tpe, string(rlp.data[rlp.ts:rlp.te]))
}

func (rlp *RagelParser) addScalarValueToSymbol(tpe symbols.ValueType, valueStr string) {
    val, err := symbols.CreateScalarValue(tpe, valueStr)
    if err != nil {
        rlp.error(err)
        return
    }

    parent, err := rlp.symStack.Peek()
    if err != nil {
        rlp.error(err)
        return
    }
    parent.Add(symbols.SCALAR, val)
}

func (rlp *RagelParser) mergeSymbolDown() {
    if err := rlp.symStack.MergeDown(); err != nil {
        rlp.error(err)
    }
}

func (rlp *RagelParser) documentStart() {
    rlp.symStack.PushNewSymbol(symbols.DOCUMENT)
}

func (rlp *RagelParser) error(err error) {
    rlp.Err = &symbols.ParseError{
        Cause: err,
        Consumed: rlp.data[:rlp.p],
        Pos: rlp.p,
    }

    rlp.cs = graphql_error
}


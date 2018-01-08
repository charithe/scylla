package internal

import "fmt"
import "github.com/charithe/scylla/graphql/symbols"

var graphql_start  int  = 20
var graphql_first_final  int  = 20
var graphql_error  int  = 0
var graphql_en_string_value  int  = 22
var graphql_en_list_value  int  = 24
var graphql_en_object_value  int  = 27
var graphql_en_after_scalar_value  int  = 31
var graphql_en_value  int  = 34
var graphql_en_key_value  int  = 51
var graphql_en_arguments  int  = 54
var graphql_en_inside_directive  int  = 58
var graphql_en_directives  int  = 61
var graphql_en_inside_field  int  = 65
var graphql_en_field  int  = 68
var graphql_en_selection_set  int  = 72
var graphql_en_main  int  = 20
var _graphql_nfa_targs [] int8  = [] int8  { 0, 0  }
var _graphql_nfa_offsets [] int8  = [] int8  { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0  }
var _graphql_nfa_push_actions [] int8  = [] int8  { 0, 0  }
var _graphql_nfa_pop_trans [] int8  = [] int8  { 0, 0  }
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
	
	
	{
		rlp.cs = int(graphql_start);
		rlp.top = 0;
		rlp.ts = 0;
		rlp.te = 0;
		rlp.act = 0;
	}
	return rlp
}

func (rlp *RagelParser) Start() (symbols.Symbol, error) {
	data := rlp.data
	eof := len(rlp.data) 
	rlp.documentStart();
	
	
	{
		if (rlp.p) == (rlp.pe) {
			goto _test_eof;
			
		}
		goto _resume;
		
		_again:
		switch rlp.cs  {
			case 20:
			goto st20;
			case 0:
			goto st0;
			case 21:
			goto st21;
			case 1:
			goto st1;
			case 2:
			goto st2;
			case 3:
			goto st3;
			case 4:
			goto st4;
			case 5:
			goto st5;
			case 6:
			goto st6;
			case 7:
			goto st7;
			case 8:
			goto st8;
			case 9:
			goto st9;
			case 10:
			goto st10;
			case 11:
			goto st11;
			case 22:
			goto st22;
			case 23:
			goto st23;
			case 24:
			goto st24;
			case 25:
			goto st25;
			case 26:
			goto st26;
			case 27:
			goto st27;
			case 28:
			goto st28;
			case 29:
			goto st29;
			case 30:
			goto st30;
			case 31:
			goto st31;
			case 32:
			goto st32;
			case 33:
			goto st33;
			case 34:
			goto st34;
			case 35:
			goto st35;
			case 12:
			goto st12;
			case 36:
			goto st36;
			case 13:
			goto st13;
			case 14:
			goto st14;
			case 37:
			goto st37;
			case 15:
			goto st15;
			case 38:
			goto st38;
			case 16:
			goto st16;
			case 39:
			goto st39;
			case 40:
			goto st40;
			case 41:
			goto st41;
			case 42:
			goto st42;
			case 43:
			goto st43;
			case 44:
			goto st44;
			case 45:
			goto st45;
			case 46:
			goto st46;
			case 47:
			goto st47;
			case 48:
			goto st48;
			case 49:
			goto st49;
			case 50:
			goto st50;
			case 51:
			goto st51;
			case 52:
			goto st52;
			case 17:
			goto st17;
			case 53:
			goto st53;
			case 18:
			goto st18;
			case 54:
			goto st54;
			case 55:
			goto st55;
			case 56:
			goto st56;
			case 57:
			goto st57;
			case 58:
			goto st58;
			case 59:
			goto st59;
			case 60:
			goto st60;
			case 61:
			goto st61;
			case 62:
			goto st62;
			case 63:
			goto st63;
			case 64:
			goto st64;
			case 65:
			goto st65;
			case 66:
			goto st66;
			case 67:
			goto st67;
			case 68:
			goto st68;
			case 69:
			goto st69;
			case 70:
			goto st70;
			case 19:
			goto st19;
			case 71:
			goto st71;
			case 72:
			goto st72;
			case 73:
			goto st73;
			case 74:
			goto st74;
			
		}
		_resume:
		switch rlp.cs  {
			case 20:
			goto st_case_20;
			case 0:
			goto st_case_0;
			case 21:
			goto st_case_21;
			case 1:
			goto st_case_1;
			case 2:
			goto st_case_2;
			case 3:
			goto st_case_3;
			case 4:
			goto st_case_4;
			case 5:
			goto st_case_5;
			case 6:
			goto st_case_6;
			case 7:
			goto st_case_7;
			case 8:
			goto st_case_8;
			case 9:
			goto st_case_9;
			case 10:
			goto st_case_10;
			case 11:
			goto st_case_11;
			case 22:
			goto st_case_22;
			case 23:
			goto st_case_23;
			case 24:
			goto st_case_24;
			case 25:
			goto st_case_25;
			case 26:
			goto st_case_26;
			case 27:
			goto st_case_27;
			case 28:
			goto st_case_28;
			case 29:
			goto st_case_29;
			case 30:
			goto st_case_30;
			case 31:
			goto st_case_31;
			case 32:
			goto st_case_32;
			case 33:
			goto st_case_33;
			case 34:
			goto st_case_34;
			case 35:
			goto st_case_35;
			case 12:
			goto st_case_12;
			case 36:
			goto st_case_36;
			case 13:
			goto st_case_13;
			case 14:
			goto st_case_14;
			case 37:
			goto st_case_37;
			case 15:
			goto st_case_15;
			case 38:
			goto st_case_38;
			case 16:
			goto st_case_16;
			case 39:
			goto st_case_39;
			case 40:
			goto st_case_40;
			case 41:
			goto st_case_41;
			case 42:
			goto st_case_42;
			case 43:
			goto st_case_43;
			case 44:
			goto st_case_44;
			case 45:
			goto st_case_45;
			case 46:
			goto st_case_46;
			case 47:
			goto st_case_47;
			case 48:
			goto st_case_48;
			case 49:
			goto st_case_49;
			case 50:
			goto st_case_50;
			case 51:
			goto st_case_51;
			case 52:
			goto st_case_52;
			case 17:
			goto st_case_17;
			case 53:
			goto st_case_53;
			case 18:
			goto st_case_18;
			case 54:
			goto st_case_54;
			case 55:
			goto st_case_55;
			case 56:
			goto st_case_56;
			case 57:
			goto st_case_57;
			case 58:
			goto st_case_58;
			case 59:
			goto st_case_59;
			case 60:
			goto st_case_60;
			case 61:
			goto st_case_61;
			case 62:
			goto st_case_62;
			case 63:
			goto st_case_63;
			case 64:
			goto st_case_64;
			case 65:
			goto st_case_65;
			case 66:
			goto st_case_66;
			case 67:
			goto st_case_67;
			case 68:
			goto st_case_68;
			case 69:
			goto st_case_69;
			case 70:
			goto st_case_70;
			case 19:
			goto st_case_19;
			case 71:
			goto st_case_71;
			case 72:
			goto st_case_72;
			case 73:
			goto st_case_73;
			case 74:
			goto st_case_74;
			
		}
		goto st_out;
		ctr7:
		{{rlp.te = (rlp.p)+1;
				{fmt.Println("MUTATION") }
			}}
		
		
		goto st20;
		ctr11:
		{{rlp.te = (rlp.p)+1;
				{fmt.Println("QUERY") }
			}}
		
		
		goto st20;
		ctr38:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.SELECTION_SET); 
					{rlp.stack[rlp.top] = 20;
						rlp.top+= 1;
						goto st72;}}
			}}
		
		
		goto st20;
		ctr40:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st20;
		st20:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof20;
			
		}
		st_case_20:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st21;
			}
			case 109:
			{
				goto st1;
			}
			case 113:
			{
				goto st8;
			}
			case 123:
			{
				goto ctr38;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st21;
			}
			
		}
		{
			goto st0;
		}
		st_case_0:
		st0:
		rlp.cs = 0;
		goto _out;
		st21:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof21;
			
		}
		st_case_21:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st21;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st21;
			}
			
		}
		{
			goto ctr40;
		}
		st1:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof1;
			
		}
		st_case_1:
		if ( data[(rlp.p)]) == 117  {
			{
				goto st2;
			}
			
		}
		{
			goto st0;
		}
		st2:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof2;
			
		}
		st_case_2:
		if ( data[(rlp.p)]) == 116  {
			{
				goto st3;
			}
			
		}
		{
			goto st0;
		}
		st3:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof3;
			
		}
		st_case_3:
		if ( data[(rlp.p)]) == 97  {
			{
				goto st4;
			}
			
		}
		{
			goto st0;
		}
		st4:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof4;
			
		}
		st_case_4:
		if ( data[(rlp.p)]) == 116  {
			{
				goto st5;
			}
			
		}
		{
			goto st0;
		}
		st5:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof5;
			
		}
		st_case_5:
		if ( data[(rlp.p)]) == 105  {
			{
				goto st6;
			}
			
		}
		{
			goto st0;
		}
		st6:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof6;
			
		}
		st_case_6:
		if ( data[(rlp.p)]) == 111  {
			{
				goto st7;
			}
			
		}
		{
			goto st0;
		}
		st7:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof7;
			
		}
		st_case_7:
		if ( data[(rlp.p)]) == 110  {
			{
				goto ctr7;
			}
			
		}
		{
			goto st0;
		}
		st8:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof8;
			
		}
		st_case_8:
		if ( data[(rlp.p)]) == 117  {
			{
				goto st9;
			}
			
		}
		{
			goto st0;
		}
		st9:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof9;
			
		}
		st_case_9:
		if ( data[(rlp.p)]) == 101  {
			{
				goto st10;
			}
			
		}
		{
			goto st0;
		}
		st10:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof10;
			
		}
		st_case_10:
		if ( data[(rlp.p)]) == 114  {
			{
				goto st11;
			}
			
		}
		{
			goto st0;
		}
		st11:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof11;
			
		}
		st_case_11:
		if ( data[(rlp.p)]) == 121  {
			{
				goto ctr11;
			}
			
		}
		{
			goto st0;
		}
		ctr41:
		{{rlp.te = (rlp.p)+1;
			}}
		
		
		goto st22;
		ctr42:
		{{rlp.te = (rlp.p)+1;
				{rlp.addScalarValueToSymbol(symbols.STRING_VALUE, string(rlp.data[rlp.tmpPtr:rlp.p])); {goto st31;}}
			}}
		
		
		goto st22;
		ctr45:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st22;
		ctr46:
		{{rlp.te = (rlp.p)+1;
			}}
		
		
		goto st22;
		st22:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof22;
			
		}
		st_case_22:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 10:
			{
				goto st0;
			}
			case 13:
			{
				goto st0;
			}
			case 34:
			{
				goto ctr42;
			}
			case 92:
			{
				goto st23;
			}
			
		}
		{
			goto ctr41;
		}
		st23:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof23;
			
		}
		st_case_23:
		if ( data[(rlp.p)]) == 34  {
			{
				goto ctr46;
			}
			
		}
		{
			goto ctr45;
		}
		ctr47:
		{{rlp.te = (rlp.p)+1;
				{rlp.rewind(); {rlp.stack[rlp.top] = 24;
						rlp.top+= 1;
						goto st34;}}
			}}
		
		
		goto st24;
		ctr50:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st24;
		ctr52:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st24;
		ctr54:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st24;
		st24:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof24;
			
		}
		st_case_24:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st25;
			}
			case 44:
			{
				goto st26;
			}
			case 93:
			{
				goto ctr50;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st25;
			}
			
		}
		{
			goto ctr47;
		}
		st25:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof25;
			
		}
		st_case_25:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st25;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st25;
			}
			
		}
		{
			goto ctr52;
		}
		st26:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof26;
			
		}
		st_case_26:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st26;
			}
			
		}
		{
			goto ctr54;
		}
		ctr58:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st27;
		ctr60:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st27;
		ctr62:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st27;
		ctr64:
		{rlp.rewind()
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.KEY_VALUE); 
					{rlp.stack[rlp.top] = 27;
						rlp.top+= 1;
						goto st51;}}
			}}
		
		
		goto st27;
		st27:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof27;
			
		}
		st_case_27:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st28;
			}
			case 44:
			{
				goto st29;
			}
			case 95:
			{
				goto st30;
			}
			case 125:
			{
				goto ctr58;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto st28;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st30;
					}
					
				}
			} 
			default:
			{
				goto st30;
			}
			
		}
		{
			goto st0;
		}
		st28:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof28;
			
		}
		st_case_28:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st28;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st28;
			}
			
		}
		{
			goto ctr60;
		}
		st29:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof29;
			
		}
		st_case_29:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st29;
			}
			
		}
		{
			goto ctr62;
		}
		st30:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof30;
			
		}
		st_case_30:
		{
			goto ctr64;
		}
		ctr65:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st31;
		ctr69:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st31;
		ctr71:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st31;
		st31:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof31;
			
		}
		st_case_31:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st32;
			}
			case 44:
			{
				goto st33;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st32;
			}
			
		}
		{
			goto ctr65;
		}
		st32:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof32;
			
		}
		st_case_32:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st32;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st32;
			}
			
		}
		{
			goto ctr69;
		}
		st33:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof33;
			
		}
		st_case_33:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st33;
			}
			
		}
		{
			goto ctr71;
		}
		ctr13:
		{{(rlp.p) = ((rlp.te))-1;
				{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.INT_VALUE);  {goto st31;}}
			}}
		
		
		goto st34;
		ctr103:
		{{switch rlp.act  {
					case 15 :
					(rlp.p) = ((rlp.te))-1;
					{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.BOOLEAN_VALUE); {goto st31;}}
					
					break;
					case 16 :
					(rlp.p) = ((rlp.te))-1;
					{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.NULL_VALUE); {goto st31;}}
					
					break;
					case 19 :
					(rlp.p) = ((rlp.te))-1;
					{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.INT_VALUE);  {goto st31;}}
					
					break;
					case 20 :
					(rlp.p) = ((rlp.te))-1;
					{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.FLOAT_VALUE);  {goto st31;}}
					
					break;
					
				}
			}
		}
		
		
		goto st34;
		ctr72:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.SCALAR); rlp.tmpPtr = rlp.p+1; {goto st22;}}
			}}
		
		
		goto st34;
		ctr76:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.LIST); {goto st24;}}
			}}
		
		
		goto st34;
		ctr80:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.OBJECT); {goto st27;}}
			}}
		
		
		goto st34;
		ctr91:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.INT_VALUE);  {goto st31;}}
			}}
		
		
		goto st34;
		ctr87:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.FLOAT_VALUE);  {goto st31;}}
			}}
		
		
		goto st34;
		ctr89:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.SCALAR); rlp.addScalarValueToSymbol(symbols.VARIABLE_VALUE, rlp.name);  {goto st31;}}
			}}
		
		
		goto st34;
		ctr112:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.SCALAR); rlp.createScalarValueFromMatch(symbols.ENUM_VALUE);  {goto st31;}}
			}}
		
		
		goto st34;
		st34:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof34;
			
		}
		st_case_34:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 0:
			{
				goto ctr21;
			}
			case 34:
			{
				goto ctr72;
			}
			case 36:
			{
				goto st15;
			}
			case 45:
			{
				goto st16;
			}
			case 91:
			{
				goto ctr76;
			}
			case 95:
			{
				goto ctr75;
			}
			case 102:
			{
				goto ctr77;
			}
			case 110:
			{
				goto ctr78;
			}
			case 116:
			{
				goto ctr79;
			}
			case 123:
			{
				goto ctr80;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 49 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto ctr22;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto ctr75;
					}
					
				}
			} 
			default:
			{
				goto ctr75;
			}
			
		}
		{
			goto st0;
		}
		ctr21:
		{{rlp.te = (rlp.p)+1;
			}}
		{{rlp.act = 19;
			}}
		
		
		goto st35;
		st35:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof35;
			
		}
		st_case_35:
		switch ( data[(rlp.p)])  {
			case 46:
			{
				goto st12;
			}
			case 69:
			{
				goto st13;
			}
			case 101:
			{
				goto st13;
			}
			
		}
		{
			goto ctr91;
		}
		st12:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof12;
			
		}
		st_case_12:
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto ctr14;
			}
			
		}
		{
			goto ctr13;
		}
		ctr14:
		{{rlp.te = (rlp.p)+1;
			}}
		{{rlp.act = 20;
			}}
		
		
		goto st36;
		st36:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof36;
			
		}
		st_case_36:
		switch ( data[(rlp.p)])  {
			case 69:
			{
				goto st13;
			}
			case 101:
			{
				goto st13;
			}
			
		}
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto ctr14;
			}
			
		}
		{
			goto ctr87;
		}
		st13:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof13;
			
		}
		st_case_13:
		switch ( data[(rlp.p)])  {
			case 43:
			{
				goto st14;
			}
			case 45:
			{
				goto st14;
			}
			
		}
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto st37;
			}
			
		}
		{
			goto ctr103;
		}
		st14:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof14;
			
		}
		st_case_14:
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto st37;
			}
			
		}
		{
			goto ctr103;
		}
		st37:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof37;
			
		}
		st_case_37:
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto st37;
			}
			
		}
		{
			goto ctr87;
		}
		st15:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof15;
			
		}
		st_case_15:
		if ( data[(rlp.p)]) == 95  {
			{
				goto ctr20;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto ctr20;
					}
					
				}
			} 
			case ( data[(rlp.p)]) >= 65 :
			{
				goto ctr20;
			}
			
		}
		{
			goto st0;
		}
		ctr20:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st38;
		st38:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof38;
			
		}
		st_case_38:
		if ( data[(rlp.p)]) == 95  {
			{
				goto st38;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st38;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st38;
					}
					
				}
			} 
			default:
			{
				goto st38;
			}
			
		}
		{
			goto ctr89;
		}
		st16:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof16;
			
		}
		st_case_16:
		if ( data[(rlp.p)]) == 0  {
			{
				goto ctr21;
			}
			
		}
		if 49 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto ctr22;
			}
			
		}
		{
			goto st0;
		}
		ctr22:
		{{rlp.te = (rlp.p)+1;
			}}
		{{rlp.act = 19;
			}}
		
		
		goto st39;
		st39:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof39;
			
		}
		st_case_39:
		switch ( data[(rlp.p)])  {
			case 46:
			{
				goto st12;
			}
			case 69:
			{
				goto st13;
			}
			case 101:
			{
				goto st13;
			}
			
		}
		if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
			{
				goto ctr22;
			}
			
		}
		{
			goto ctr91;
		}
		ctr75:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st40;
		st40:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof40;
			
		}
		st_case_40:
		if ( data[(rlp.p)]) == 95  {
			{
				goto st40;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		ctr77:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st41;
		st41:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof41;
			
		}
		st_case_41:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 97:
			{
				goto st42;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 98 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st42:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof42;
			
		}
		st_case_42:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 108:
			{
				goto st43;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st43:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof43;
			
		}
		st_case_43:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 115:
			{
				goto st44;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st44:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof44;
			
		}
		st_case_44:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 101:
			{
				goto ctr102;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		ctr102:
		{{rlp.te = (rlp.p)+1;
			}}
		{{rlp.act = 15;
			}}
		
		
		goto st45;
		ctr109:
		{{rlp.te = (rlp.p)+1;
			}}
		{{rlp.act = 16;
			}}
		
		
		goto st45;
		st45:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof45;
			
		}
		st_case_45:
		if ( data[(rlp.p)]) == 95  {
			{
				goto st40;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr103;
		}
		ctr78:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st46;
		st46:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof46;
			
		}
		st_case_46:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 117:
			{
				goto st47;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st47:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof47;
			
		}
		st_case_47:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 108:
			{
				goto st48;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st48:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof48;
			
		}
		st_case_48:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 108:
			{
				goto ctr109;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		ctr79:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st49;
		st49:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof49;
			
		}
		st_case_49:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 114:
			{
				goto st50;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		st50:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof50;
			
		}
		st_case_50:
		switch ( data[(rlp.p)])  {
			case 95:
			{
				goto st40;
			}
			case 117:
			{
				goto st44;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st40;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st40;
					}
					
				}
			} 
			default:
			{
				goto st40;
			}
			
		}
		{
			goto ctr112;
		}
		ctr27:
		{{(rlp.p) = ((rlp.te))-1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st51;
		ctr113:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st51;
		ctr116:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st51;
		ctr118:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.addNameToSymbolAs(symbols.KEY); {rlp.stack[rlp.top] = 51;
						rlp.top+= 1;
						goto st34;}}
			}}
		
		
		goto st51;
		st51:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof51;
			
		}
		st_case_51:
		{{rlp.ts = (rlp.p);
			}}
		if ( data[(rlp.p)]) == 95  {
			{
				goto ctr114;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto ctr114;
					}
					
				}
			} 
			case ( data[(rlp.p)]) >= 65 :
			{
				goto ctr114;
			}
			
		}
		{
			goto ctr113;
		}
		ctr114:
		{{rlp.te = (rlp.p)+1;
			}}
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st52;
		st52:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof52;
			
		}
		st_case_52:
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto ctr28;
			}
			case 58:
			{
				goto ctr30;
			}
			case 95:
			{
				goto st18;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 48 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto ctr28;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 57 :
			{
				switch {
					case ( data[(rlp.p)]) > 90 :
					{
						if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
							{
								goto st18;
							}
							
						}
					} 
					case ( data[(rlp.p)]) >= 65 :
					{
						goto st18;
					}
					
				}
			} 
			default:
			{
				goto st18;
			}
			
		}
		{
			goto ctr116;
		}
		ctr28:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		
		
		goto st17;
		st17:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof17;
			
		}
		st_case_17:
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st17;
			}
			case 58:
			{
				goto st53;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st17;
			}
			
		}
		{
			goto ctr27;
		}
		ctr30:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		
		
		goto st53;
		st53:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof53;
			
		}
		st_case_53:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st53;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st53;
			}
			
		}
		{
			goto ctr118;
		}
		st18:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof18;
			
		}
		st_case_18:
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto ctr28;
			}
			case 58:
			{
				goto ctr30;
			}
			case 95:
			{
				goto st18;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 48 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto ctr28;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 57 :
			{
				switch {
					case ( data[(rlp.p)]) > 90 :
					{
						if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
							{
								goto st18;
							}
							
						}
					} 
					case ( data[(rlp.p)]) >= 65 :
					{
						goto st18;
					}
					
				}
			} 
			default:
			{
				goto st18;
			}
			
		}
		{
			goto ctr27;
		}
		ctr120:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st54;
		ctr124:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st54;
		ctr126:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st54;
		ctr128:
		{rlp.rewind()
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.KEY_VALUE); 
					{rlp.stack[rlp.top] = 54;
						rlp.top+= 1;
						goto st51;}}
			}}
		
		
		goto st54;
		st54:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof54;
			
		}
		st_case_54:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st55;
			}
			case 41:
			{
				goto ctr120;
			}
			case 44:
			{
				goto st56;
			}
			case 95:
			{
				goto st57;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto st55;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st57;
					}
					
				}
			} 
			default:
			{
				goto st57;
			}
			
		}
		{
			goto st0;
		}
		st55:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof55;
			
		}
		st_case_55:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st55;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st55;
			}
			
		}
		{
			goto ctr124;
		}
		st56:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof56;
			
		}
		st_case_56:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st56;
			}
			
		}
		{
			goto ctr126;
		}
		st57:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof57;
			
		}
		st_case_57:
		{
			goto ctr128;
		}
		ctr129:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st58;
		ctr131:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.ARGUMENTS); 
					{rlp.stack[rlp.top] = 58;
						rlp.top+= 1;
						goto st54;}}
			}}
		
		
		goto st58;
		ctr134:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st58;
		ctr136:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st58;
		st58:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof58;
			
		}
		st_case_58:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st59;
			}
			case 40:
			{
				goto ctr131;
			}
			case 44:
			{
				goto st60;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st59;
			}
			
		}
		{
			goto ctr129;
		}
		st59:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof59;
			
		}
		st_case_59:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st59;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st59;
			}
			
		}
		{
			goto ctr134;
		}
		st60:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof60;
			
		}
		st_case_60:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st60;
			}
			
		}
		{
			goto ctr136;
		}
		ctr137:
		{{rlp.te = (rlp.p)+1;
				{rlp.rewind(); {rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} }
			}}
		
		
		goto st61;
		ctr141:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st61;
		ctr143:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.rewind(); {rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} }
			}}
		
		
		goto st61;
		ctr146:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.DIRECTIVE); rlp.addNameToSymbolAs(symbols.NAME); {rlp.stack[rlp.top] = 61;
						rlp.top+= 1;
						goto st58;}}
			}}
		
		
		goto st61;
		st61:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof61;
			
		}
		st_case_61:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st62;
			}
			case 64:
			{
				goto st63;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st62;
			}
			
		}
		{
			goto ctr137;
		}
		st62:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof62;
			
		}
		st_case_62:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st62;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st62;
			}
			
		}
		{
			goto ctr141;
		}
		st63:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof63;
			
		}
		st_case_63:
		if ( data[(rlp.p)]) == 95  {
			{
				goto ctr144;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto ctr144;
					}
					
				}
			} 
			case ( data[(rlp.p)]) >= 65 :
			{
				goto ctr144;
			}
			
		}
		{
			goto ctr143;
		}
		ctr144:
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st64;
		st64:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof64;
			
		}
		st_case_64:
		if ( data[(rlp.p)]) == 95  {
			{
				goto st64;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 48 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 57  {
					{
						goto st64;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st64;
					}
					
				}
			} 
			default:
			{
				goto st64;
			}
			
		}
		{
			goto ctr146;
		}
		ctr148:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st65;
		ctr150:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.ARGUMENTS); 
					{rlp.stack[rlp.top] = 65;
						rlp.top+= 1;
						goto st54;}}
			}}
		
		
		goto st65;
		ctr152:
		{{rlp.te = (rlp.p)+1;
				{rlp.rewind(); {rlp.stack[rlp.top] = 65;
						rlp.top+= 1;
						goto st61;}}
			}}
		
		
		goto st65;
		ctr153:
		{{rlp.te = (rlp.p)+1;
				{rlp.startSymbol(symbols.SELECTION_SET); 
					{rlp.stack[rlp.top] = 65;
						rlp.top+= 1;
						goto st72;}}
			}}
		
		
		goto st65;
		ctr155:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st65;
		ctr157:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st65;
		st65:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof65;
			
		}
		st_case_65:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st66;
			}
			case 40:
			{
				goto ctr150;
			}
			case 44:
			{
				goto st67;
			}
			case 64:
			{
				goto ctr152;
			}
			case 123:
			{
				goto ctr153;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st66;
			}
			
		}
		{
			goto ctr148;
		}
		st66:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof66;
			
		}
		st_case_66:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st66;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st66;
			}
			
		}
		{
			goto ctr155;
		}
		st67:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof67;
			
		}
		st_case_67:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st67;
			}
			
		}
		{
			goto ctr157;
		}
		ctr32:
		{{(rlp.p) = ((rlp.te))-1;
				{rlp.addNameToSymbolAs(symbols.NAME); {goto st65;}}
			}}
		
		
		goto st68;
		ctr158:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(true); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st68;
		ctr162:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st68;
		ctr164:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.addNameToSymbolAs(symbols.NAME); {goto st65;}}
			}}
		
		
		goto st68;
		ctr169:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.addNameToSymbolAs(symbols.ALIAS) }
			}}
		
		
		goto st68;
		st68:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof68;
			
		}
		st_case_68:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 44:
			{
				goto st69;
			}
			case 95:
			{
				goto ctr160;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto ctr160;
					}
					
				}
			} 
			case ( data[(rlp.p)]) >= 65 :
			{
				goto ctr160;
			}
			
		}
		{
			goto ctr158;
		}
		st69:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof69;
			
		}
		st_case_69:
		if ( data[(rlp.p)]) == 44  {
			{
				goto st69;
			}
			
		}
		{
			goto ctr162;
		}
		ctr160:
		{{rlp.te = (rlp.p)+1;
			}}
		{rlp.tmpPtr = rlp.p; 
		}
		
		
		goto st70;
		ctr166:
		{{rlp.te = (rlp.p)+1;
			}}
		
		
		goto st70;
		st70:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof70;
			
		}
		st_case_70:
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto ctr165;
			}
			case 58:
			{
				goto ctr167;
			}
			case 95:
			{
				goto ctr166;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 48 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto ctr165;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 57 :
			{
				switch {
					case ( data[(rlp.p)]) > 90 :
					{
						if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
							{
								goto ctr166;
							}
							
						}
					} 
					case ( data[(rlp.p)]) >= 65 :
					{
						goto ctr166;
					}
					
				}
			} 
			default:
			{
				goto ctr166;
			}
			
		}
		{
			goto ctr164;
		}
		ctr165:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		
		
		goto st19;
		st19:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof19;
			
		}
		st_case_19:
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st19;
			}
			case 58:
			{
				goto st71;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st19;
			}
			
		}
		{
			goto ctr32;
		}
		ctr167:
		{rlp.name = string(rlp.data[rlp.tmpPtr:rlp.p]); 
		}
		
		
		goto st71;
		st71:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof71;
			
		}
		st_case_71:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st71;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st71;
			}
			
		}
		{
			goto ctr169;
		}
		ctr172:
		{{rlp.te = (rlp.p)+1;
				{rlp.endSymbol(false); 
					{rlp.top -= 1;
						rlp.cs = rlp.stack[rlp.top];
						goto _again;} 
				}
			}}
		
		
		goto st72;
		ctr174:
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
			}}
		
		
		goto st72;
		ctr176:
		{rlp.rewind()
		}
		{{rlp.te = (rlp.p);
				(rlp.p) = (rlp.p) - 1;
				{rlp.startSymbol(symbols.FIELD); {rlp.stack[rlp.top] = 72;
						rlp.top+= 1;
						goto st68;}}
			}}
		
		
		goto st72;
		st72:
		{{rlp.ts = 0;
			}}
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof72;
			
		}
		st_case_72:
		{{rlp.ts = (rlp.p);
			}}
		switch ( data[(rlp.p)])  {
			case 32:
			{
				goto st73;
			}
			case 95:
			{
				goto st74;
			}
			case 125:
			{
				goto ctr172;
			}
			
		}
		switch {
			case ( data[(rlp.p)]) < 65 :
			{
				if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
					{
						goto st73;
					}
					
				}
			} 
			case ( data[(rlp.p)]) > 90 :
			{
				if 97 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 122  {
					{
						goto st74;
					}
					
				}
			} 
			default:
			{
				goto st74;
			}
			
		}
		{
			goto st0;
		}
		st73:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof73;
			
		}
		st_case_73:
		if ( data[(rlp.p)]) == 32  {
			{
				goto st73;
			}
			
		}
		if 9 <= ( data[(rlp.p)]) && ( data[(rlp.p)]) <= 13  {
			{
				goto st73;
			}
			
		}
		{
			goto ctr174;
		}
		st74:
		(rlp.p)+= 1;
		if (rlp.p) == (rlp.pe) {
			goto _test_eof74;
			
		}
		st_case_74:
		{
			goto ctr176;
		}
		st_out:
		_test_eof20: rlp.cs = 20;
		goto _test_eof; 
		_test_eof21: rlp.cs = 21;
		goto _test_eof; 
		_test_eof1: rlp.cs = 1;
		goto _test_eof; 
		_test_eof2: rlp.cs = 2;
		goto _test_eof; 
		_test_eof3: rlp.cs = 3;
		goto _test_eof; 
		_test_eof4: rlp.cs = 4;
		goto _test_eof; 
		_test_eof5: rlp.cs = 5;
		goto _test_eof; 
		_test_eof6: rlp.cs = 6;
		goto _test_eof; 
		_test_eof7: rlp.cs = 7;
		goto _test_eof; 
		_test_eof8: rlp.cs = 8;
		goto _test_eof; 
		_test_eof9: rlp.cs = 9;
		goto _test_eof; 
		_test_eof10: rlp.cs = 10;
		goto _test_eof; 
		_test_eof11: rlp.cs = 11;
		goto _test_eof; 
		_test_eof22: rlp.cs = 22;
		goto _test_eof; 
		_test_eof23: rlp.cs = 23;
		goto _test_eof; 
		_test_eof24: rlp.cs = 24;
		goto _test_eof; 
		_test_eof25: rlp.cs = 25;
		goto _test_eof; 
		_test_eof26: rlp.cs = 26;
		goto _test_eof; 
		_test_eof27: rlp.cs = 27;
		goto _test_eof; 
		_test_eof28: rlp.cs = 28;
		goto _test_eof; 
		_test_eof29: rlp.cs = 29;
		goto _test_eof; 
		_test_eof30: rlp.cs = 30;
		goto _test_eof; 
		_test_eof31: rlp.cs = 31;
		goto _test_eof; 
		_test_eof32: rlp.cs = 32;
		goto _test_eof; 
		_test_eof33: rlp.cs = 33;
		goto _test_eof; 
		_test_eof34: rlp.cs = 34;
		goto _test_eof; 
		_test_eof35: rlp.cs = 35;
		goto _test_eof; 
		_test_eof12: rlp.cs = 12;
		goto _test_eof; 
		_test_eof36: rlp.cs = 36;
		goto _test_eof; 
		_test_eof13: rlp.cs = 13;
		goto _test_eof; 
		_test_eof14: rlp.cs = 14;
		goto _test_eof; 
		_test_eof37: rlp.cs = 37;
		goto _test_eof; 
		_test_eof15: rlp.cs = 15;
		goto _test_eof; 
		_test_eof38: rlp.cs = 38;
		goto _test_eof; 
		_test_eof16: rlp.cs = 16;
		goto _test_eof; 
		_test_eof39: rlp.cs = 39;
		goto _test_eof; 
		_test_eof40: rlp.cs = 40;
		goto _test_eof; 
		_test_eof41: rlp.cs = 41;
		goto _test_eof; 
		_test_eof42: rlp.cs = 42;
		goto _test_eof; 
		_test_eof43: rlp.cs = 43;
		goto _test_eof; 
		_test_eof44: rlp.cs = 44;
		goto _test_eof; 
		_test_eof45: rlp.cs = 45;
		goto _test_eof; 
		_test_eof46: rlp.cs = 46;
		goto _test_eof; 
		_test_eof47: rlp.cs = 47;
		goto _test_eof; 
		_test_eof48: rlp.cs = 48;
		goto _test_eof; 
		_test_eof49: rlp.cs = 49;
		goto _test_eof; 
		_test_eof50: rlp.cs = 50;
		goto _test_eof; 
		_test_eof51: rlp.cs = 51;
		goto _test_eof; 
		_test_eof52: rlp.cs = 52;
		goto _test_eof; 
		_test_eof17: rlp.cs = 17;
		goto _test_eof; 
		_test_eof53: rlp.cs = 53;
		goto _test_eof; 
		_test_eof18: rlp.cs = 18;
		goto _test_eof; 
		_test_eof54: rlp.cs = 54;
		goto _test_eof; 
		_test_eof55: rlp.cs = 55;
		goto _test_eof; 
		_test_eof56: rlp.cs = 56;
		goto _test_eof; 
		_test_eof57: rlp.cs = 57;
		goto _test_eof; 
		_test_eof58: rlp.cs = 58;
		goto _test_eof; 
		_test_eof59: rlp.cs = 59;
		goto _test_eof; 
		_test_eof60: rlp.cs = 60;
		goto _test_eof; 
		_test_eof61: rlp.cs = 61;
		goto _test_eof; 
		_test_eof62: rlp.cs = 62;
		goto _test_eof; 
		_test_eof63: rlp.cs = 63;
		goto _test_eof; 
		_test_eof64: rlp.cs = 64;
		goto _test_eof; 
		_test_eof65: rlp.cs = 65;
		goto _test_eof; 
		_test_eof66: rlp.cs = 66;
		goto _test_eof; 
		_test_eof67: rlp.cs = 67;
		goto _test_eof; 
		_test_eof68: rlp.cs = 68;
		goto _test_eof; 
		_test_eof69: rlp.cs = 69;
		goto _test_eof; 
		_test_eof70: rlp.cs = 70;
		goto _test_eof; 
		_test_eof19: rlp.cs = 19;
		goto _test_eof; 
		_test_eof71: rlp.cs = 71;
		goto _test_eof; 
		_test_eof72: rlp.cs = 72;
		goto _test_eof; 
		_test_eof73: rlp.cs = 73;
		goto _test_eof; 
		_test_eof74: rlp.cs = 74;
		goto _test_eof; 
		
		_test_eof: {}
		if (rlp.p) == eof  {
			{
				switch rlp.cs  {
					case 21:
					goto ctr40;
					case 23:
					goto ctr45;
					case 25:
					goto ctr52;
					case 26:
					goto ctr54;
					case 28:
					goto ctr60;
					case 29:
					goto ctr62;
					case 30:
					goto ctr64;
					case 32:
					goto ctr69;
					case 33:
					goto ctr71;
					case 35:
					goto ctr91;
					case 12:
					goto ctr13;
					case 36:
					goto ctr87;
					case 13:
					goto ctr103;
					case 14:
					goto ctr103;
					case 37:
					goto ctr87;
					case 38:
					goto ctr89;
					case 39:
					goto ctr91;
					case 40:
					goto ctr112;
					case 41:
					goto ctr112;
					case 42:
					goto ctr112;
					case 43:
					goto ctr112;
					case 44:
					goto ctr112;
					case 45:
					goto ctr103;
					case 46:
					goto ctr112;
					case 47:
					goto ctr112;
					case 48:
					goto ctr112;
					case 49:
					goto ctr112;
					case 50:
					goto ctr112;
					case 52:
					goto ctr116;
					case 17:
					goto ctr27;
					case 53:
					goto ctr118;
					case 18:
					goto ctr27;
					case 55:
					goto ctr124;
					case 56:
					goto ctr126;
					case 57:
					goto ctr128;
					case 59:
					goto ctr134;
					case 60:
					goto ctr136;
					case 62:
					goto ctr141;
					case 63:
					goto ctr143;
					case 64:
					goto ctr146;
					case 66:
					goto ctr155;
					case 67:
					goto ctr157;
					case 69:
					goto ctr162;
					case 70:
					goto ctr164;
					case 19:
					goto ctr32;
					case 71:
					goto ctr169;
					case 73:
					goto ctr174;
					case 74:
					goto ctr176;
					
				}
			}
			
			
		}
		_out: {}
	}
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
	fmt.Printf("\n<<<ERROR: %+v>>>\n", err)
	rlp.Err = &symbols.ParseError{
		Cause: err,
		Consumed: rlp.data[:rlp.p],
		Pos: rlp.p,
	}
	
	rlp.cs = graphql_error
}


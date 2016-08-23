package gocc

type Token int

const (
	start Token = iota

	start_lit
	TRUE_LITER
	FALSE_LITER
	INT_LITER
	CHAR_LIT
	STR_LIT
	end_lit
	begin_kw
	BEGIN
	END
	NULL
	READ
	FREE
	RETURN
	EXIT
	PRINT
	IF
	THEN
	ELSE
	FI
	WHILE
	CALL
	end_kw
	begin_type
	INT_T
	CHAR_T
	STRING_T
	end_type
	ID
	begin_symbol
	PLUS,
	MINUS,
	DIV,
	REM,
	MUL,
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	end_symbol
	COMMENT
	ILLEGAL
	end
)


var tok_strings = map[Token]string{
	ILLEGAL: "Illegal",
	COMMENT: "Comment",
	INT_LITER: "Integer",
	TRUE_LITER: "true",
	FALSE_LITER: "false",
	CHAR_LIT: ".",
	STR_LIT: "[\".*?\"]",
	LPAREN:  "(",
	RPAREN:  ")",
	PLUS: "+",
	MINUS: "-",
	MUL: "*",
	DIV: "/",
	REM: "%",
	BEGIN: "begin",
	END: "end",
	NULL: "null",
	READ: "read",
	FREE: "free",
	RETURN: "return",
	EXIT: "exit",
	PRINT: "print",
	IF: "if",
	THEN: "then",
	ELSE: "else",
	FI: "fi",
	WHILE: "while",
	CALL: "call",
	INT_T: "int",
	CHAR_T: "char",
	STRING_T: "string",
	ID: "id",
	LBRACE: "{",
	RBRACE: "}",
}

func (t Token) IsLiteral() bool {
	return t > start_lit && t < end_lit
}

func (t Token) IsOperator() bool {
	return t > begin_symbol && t < end_symbol
}

func (t Token) IsType() bool {
	return t > begin_type && t < end_type
}

func (t Token) IsKeyword() bool {
	return t > begin_kw && t < end_kw
}

func (t Token) IsID() bool {
	return t == ID
}

func Lookup(str string) Token {
	for t, s := range tok_strings {
		if s == str {
			return t
		}
	}
	return ILLEGAL
}

func (t Token) String() string {
	return tok_strings[t]
}

func (t Token) Valid() bool {
	return t > start && t < end
}
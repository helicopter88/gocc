package scanner

import (
	"regexp"
	"fmt"
	"strings"
	"io"
)

type TokenValue struct {
	Tok Token
	Str string
}

func (t TokenValue) String() string {
	return fmt.Sprintf("%s: %s", t.Tok.String(), t.Str)
}

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
	ASSIGN
	PLUS
	MINUS
	DIV
	REM
	MUL
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	end_symbol
	COMMENT
	ILLEGAL
	EOF
	end
)

type GoccScanner struct {
	*File
	currLineTokens []string
	currentToken int
}

func NewScanner(src string) *GoccScanner {
	s := new(GoccScanner)
	file, err := NewFile(src)
	if err != nil {
		panic(err)
	}
	s.File = file
	s.currentToken = 0
	s.currentLn = 0
	return s
}

func parseToken(str string) (TokenValue, error) {
	tok := ILLEGAL
	switch str {
	case "-":
		tok = MINUS
	case "+":
		tok = PLUS
	case "/":
		tok = DIV
	case "*":
		tok = MUL
	case "%":
		tok = REM
	case "=":
		tok = ASSIGN
	case "{":
		tok = LBRACE
	case "}":
		tok = RBRACE
	case "(":
		tok = LPAREN
	case ")":
		tok = RPAREN
	case "true":
		tok = TRUE_LITER
	case "false:":
		tok = FALSE_LITER
	case "begin":
		tok = BEGIN
	case "end":
		tok = END
	case "null":
		tok = NULL
	case "READ":
		tok = READ
	case "int":
		tok = INT_T
	case "char":
		tok = CHAR_T
	case "string":
		tok = STRING_T
	case "if":
		tok = IF
	case "fi":
		tok = FI
	case "else":
		tok = ELSE
	case "while":
		tok = WHILE
	default:
		t, err := parseLiteral(str)
		if err != nil {
			return TokenValue{t, ""}, err
		}
		tok = t
	}
	return TokenValue{tok, str}, nil
}

func (s *GoccScanner) Next() (TokenValue, error) {
	if s.currentLn == s.maxLen - 1 {
		return TokenValue{}, io.EOF
	}
	line := s.Lines[s.currentLn]
	if len(s.currLineTokens) == 0 {
		s.currLineTokens = strings.Split(line, " ")
	}
	tokens := s.currLineTokens
	if s.currentToken == len(s.currLineTokens) {
		fmt.Println("hello")
		s.currentToken = 0
		s.currentLn += 1
		s.currLineTokens = s.currLineTokens[:0]
		return s.Next()
	}
	s.currentToken += 1
	return parseToken(tokens[s.currentToken - 1])
}

func parseLiteral(tok string) (Token, error) {
	if match, err := regexp.MatchString("\".*?\"", tok); match {
		return STR_LIT, err
	}
	if match, err := regexp.MatchString("[0-9]+", tok); match {
		return INT_LITER, err
	}
	if match, err := regexp.MatchString("[a-zA-Z]+", tok); match {
		return ID, err
	}
	return ILLEGAL, fmt.Errorf("Invalid token %s", tok)
}
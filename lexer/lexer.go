package lexer

import "monkey/token"

type Lexer struct {
	input        string // input string
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // read the first character
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace() // skip whitespaces

	switch l.ch {
	case '=':
		if l.PeekChar() == '=' { // if the next character is '='
			ch := l.ch                           // save the current character
			l.readChar()                         // read the next character
			literal := string(ch) + string(l.ch) // concatenate the characters
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.PeekChar() == '=' { // if the next character is '='
			ch := l.ch                           // save the current character
			l.readChar()                         // read the next character
			literal := string(ch) + string(l.ch) // concatenate the characters
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = "" // end of file
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier() // read the identifier
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber() // read the number
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // read the next character
	return tok
}

// newToken creates a new token.Token with the given token.TokenType and byte.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // if we reach the end of input
		l.ch = 0 // ASCII code for "NUL" character
	} else {
		l.ch = l.input[l.readPosition] // read the next character
	}
	l.position = l.readPosition // update the current position
	l.readPosition += 1         // update the reading position
}

// readIdentifier reads the identifier and advances the lexer's position until
// it encounters a non-letter character.
func (l *Lexer) readIdentifier() string {
	position := l.position // save the current position
	for isLetter(l.ch) {
		l.readChar() // read the next character
	}
	return l.input[position:l.position] // return the identifier
}

// isLetter returns true if the given byte is a letter or an underscore.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_'
}

// skipWhitespace skips whitespaces.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' { // if the current character is a whitespace
		l.readChar() // read the next character
	}
}

// readNumber reads the number and advances the lexer's position until it
// encounters a non-digit character.
func (l *Lexer) readNumber() string {
	position := l.position // save the current position
	for isDigit(l.ch) {
		l.readChar() // read the next character
	}
	return l.input[position:l.position] // return the number
}

// isDigit returns true if the given byte is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// PeekChar returns the next character without advancing the lexer's position.
func (l *Lexer) PeekChar() byte {
	if l.readPosition >= len(l.input) { // if we reach the end of input
		return 0 // ASCII code for "NUL" character
	} else {
		return l.input[l.readPosition] // read the next character
	}
}

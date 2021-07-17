package parse

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	input        string
	position     int // current position
	readPosition int // next position
	ch           byte

	lineHead int
	line     int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 is NULL (EOF)
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++

	if isBreak(l.ch) {
		l.lineHead = l.position + 1
		l.line++
	}
}

func (l *Lexer) genPosition() Position {
	return Position{Line: l.line + 1, Column: l.readPosition - l.lineHead + 1}
}

// NextToken return next token
func (l *Lexer) NextToken() (int, string, Position) {
	var (
		token   int
		literal string
	)

	position := l.genPosition()

	switch ch := l.ch; {
	case isWhiteSpace(ch):
		indent, count := l.readIndent()
		if count == 1 {
			//l.readChar()
			return l.NextToken()
		} else {
			token = INDENT
			literal = indent
		}
	case isLetter(ch):
		token = IDENT
		literal = l.readIdentifier()
	case isDigit(ch):
		token = NUMBER
		literal = l.readNumber()
	case isBreak(ch):
		token = LINE_BREAK
		literal = "LF"
		l.readChar()
	default:
		switch ch {
		case 0:
			token = EOF
			literal = ""
		case '[', ']', '(', ')', ';', '+', '-', '*', '%', '=':
			token = int(ch)
			literal = string(ch)
		}

		l.readChar()
	}

	return token, literal, position
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readIndent() (string, int) {
	position := l.position
	count := 0
	for isWhiteSpace(l.ch) {
		l.readChar()
		count++
	}

	return l.input[position:l.position], count
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isBreak(ch byte) bool {
	return ch == '\n'
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}

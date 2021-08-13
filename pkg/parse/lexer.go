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
	return Position{
		Line:   l.line + 1,
		Column: l.readPosition - l.lineHead + 1,
	}
}

func (l *Lexer) NextToken() (int, string, Position) {
	var (
		token   int
		literal string
	)

	position := l.genPosition()

	switch ch := l.ch; {
	// INDENT
	case isWhiteSpace(ch):
		indent, count := l.readIndent()
		if count == 1 {
			return l.NextToken()
		} else {
			token = INDENT
			literal = indent
		}
	// IDENT
	case isLetter(ch):
		token = IDENT
		literal = l.readIdentifier()
	// VALUE
	case ch == '(':
		token = VALUE
		literal = l.readValue()
		l.readChar()
	// [
	case ch == '[':
		//token = LBRA
		token = int(ch)
		literal = "["
		l.readChar()
	// NUMBER
	case isDigit(ch):
		token = NUMBER
		literal = l.readNumber()
	// -
	case ch == '-':
		//token = HYPHEN
		token = int(ch)
		literal = "-"
		l.readChar()
	// ]
	case ch == ']':
		//token = RBRA
		token = int(ch)
		literal = "]"
		l.readChar()
	// LineBREAK
	case isBreak(ch):
		token = LINE_BREAK
		literal = "LF"
		l.readChar()
	default:
		switch ch {
		case 0:
			token = EOF
			literal = ""
		}
		l.readChar()
	}

	return token, literal, position
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

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readValue() string {
	position := l.position + 1
	for {
		l.readChar()
		if isEndValue(l.ch) || l.ch == 0 {
			break
		}
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

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isStartValue(ch byte) bool {
	return ch == '('
}

func isEndValue(ch byte) bool {
	return ch == ')'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isBreak(ch byte) bool {
	return ch == '\n'
}

// Code generated by goyacc parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parse

import __yyfmt__ "fmt"

//line parser.go.y:2

import (
	"fmt"
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

var currentDepth = 0

//line parser.go.y:19
type yySymType struct {
	yys        int
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

const EOF = 57346
const IDENT = 57347
const NUMBER = 57348
const INDENT = 57349
const LINE_BREAK = 57350
const LBRA = 57351
const RBRA = 57352
const HYPHEN = 57353
const VALUE = 57354

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"EOF",
	"IDENT",
	"NUMBER",
	"INDENT",
	"LINE_BREAK",
	"LBRA",
	"RBRA",
	"HYPHEN",
	"VALUE",
	"';'",
	"'['",
	"'-'",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:168

type LexerWrapper struct {
	l          *Lexer
	recentLit  string
	recentPos  Position
	statements []Statement
	Stack      *Stack
}

func (w *LexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := w.l.NextToken()
	if tok == EOF {
		return 0
	}

	fmt.Println(lit)

	lval.tok = Token{
		tok: tok,
		lit: lit,
		pos: pos,
	}
	w.recentLit = lit
	w.recentPos = pos
	return tok
}

func (w *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d, Lit: %s, Error: %s",
		w.recentPos.Line, w.recentPos.Column, w.recentLit, e)
}

func Parse(l *Lexer) *Stack {
	w := LexerWrapper{
		l:     l,
		Stack: &Stack{},
	}

	if yyParse(&w) != 0 {
		panic("Parse error")
	}

	return w.Stack
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 32

var yyAct = [...]int{
	32, 30, 29, 25, 28, 23, 22, 17, 16, 11,
	15, 10, 20, 14, 8, 5, 12, 6, 4, 9,
	31, 27, 26, 24, 21, 19, 18, 13, 1, 3,
	2, 7,
}

var yyPact = [...]int{
	10, -1000, 10, 6, -1000, -3, 11, -1000, -1000, -1000,
	21, -1, -4, -8, 20, 19, -2, 18, -9, -10,
	17, -13, 16, 15, -11, -1000, -14, -15, 14, -1000,
	-1000, -16, -1000,
}

var yyPgo = [...]int{
	0, 28, 30, 29,
}

var yyR1 = [...]int{
	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
}

var yyR2 = [...]int{
	0, 0, 2, 2, 2, 1, 6, 7, 7, 8,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, 8, 5, 7, -1, 8, 13,
	14, 12, 5, 6, 14, 14, 12, 15, 6, 6,
	14, 6, 15, 15, 6, 16, 6, 6, 15, 16,
	16, 6, 16,
}

var yyDef = [...]int{
	1, -2, 1, 0, 5, 0, 0, 2, 3, 4,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 6, 0, 0, 0, 7,
	8, 0, 9,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 15, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 14, 3, 16,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:35
		{
			yyVAL.statements = nil
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:42
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:51
		{
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				yyVAL.statement = &ExpressionStatement{Expression: yyDollar[1].expr}

				sqlExpr := yyDollar[1].expr.(*SQLExpression)
				depth := sqlExpr.Depth
				if depth == 0 {
					l.Stack.Push(sqlExpr)
				} else if currentDepth > depth {
					v := currentDepth - depth
					for i := 0; i < v+1; i++ {
						l.Stack.Pop()
					}
					node := l.Stack.Pop()
					node.Expressions = append(node.Expressions, sqlExpr)

					l.Stack.Push(node)
					l.Stack.Push(sqlExpr)

					currentDepth = depth
				} else {
					parent := l.Stack.Pop()
					if parent != nil {
						parent.Expressions = append(parent.Expressions, sqlExpr)
						l.Stack.Push(parent)
					}
					l.Stack.Push(sqlExpr)

					currentDepth++
				}
			}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:84
		{
			if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
				yyVAL.statement = &ExpressionStatement{Expression: yyDollar[1].expr}

				sqlExpr := yyDollar[1].expr.(*SQLExpression)
				depth := sqlExpr.Depth
				if depth == 0 {
					l.Stack.Push(sqlExpr)
				} else if currentDepth > depth {
					v := currentDepth - depth
					for i := 0; i < v+1; i++ {
						l.Stack.Pop()
					}
					node := l.Stack.Pop()
					node.Expressions = append(node.Expressions, sqlExpr)

					l.Stack.Push(node)
				} else {
					parent := l.Stack.Pop()
					if parent != nil {
						parent.Expressions = append(parent.Expressions, sqlExpr)
						l.Stack.Push(parent)
					}
					l.Stack.Push(sqlExpr)
				}
			}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:114
		{
			yyVAL.expr = &LineBreakExpression{Literal: yyDollar[1].tok.lit}
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:118
		{
			yyVAL.expr = &SQLExpression{
				Literal: yyDollar[1].tok.lit,
				Value:   "",
				PositionExpression: &PositionExpression{
					Start: yyDollar[3].tok.lit,
					End:   yyDollar[5].tok.lit,
				},
				Depth: 0,
			}
		}
	case 7:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:130
		{
			yyVAL.expr = &SQLExpression{
				Literal: yyDollar[1].tok.lit,
				Value:   yyDollar[2].tok.lit,
				PositionExpression: &PositionExpression{
					Start: yyDollar[4].tok.lit,
					End:   yyDollar[6].tok.lit,
				},
				Depth: 0,
			}
		}
	case 8:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:142
		{
			depth := len(yyDollar[1].tok.lit) / 2
			yyVAL.expr = &SQLExpression{
				Literal: yyDollar[2].tok.lit,
				Value:   "",
				PositionExpression: &PositionExpression{
					Start: yyDollar[4].tok.lit,
					End:   yyDollar[6].tok.lit,
				},
				Depth: depth,
			}
		}
	case 9:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:155
		{
			depth := len(yyDollar[1].tok.lit) / 2
			yyVAL.expr = &SQLExpression{
				Literal: yyDollar[2].tok.lit,
				Value:   yyDollar[3].tok.lit,
				PositionExpression: &PositionExpression{
					Start: yyDollar[5].tok.lit,
					End:   yyDollar[7].tok.lit,
				},
				Depth: depth,
			}
		}
	}
	goto yystack /* stack new state and value */
}
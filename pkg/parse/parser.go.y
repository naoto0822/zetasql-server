%{
package parse

import (
	"log"
)

type Token struct {
	tok int
	lit string
	pos Position
}

var currentDepth = 0

%}

%union{
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

%type<statements> statements
%type<statement> statement
%type<expr> expr
%token<tok> EOF IDENT NUMBER INDENT LINE_BREAK LBRA RBRA HYPHEN VALUE

%%

statements
	:
	{
		$$ = nil
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}
	| statement statements
	{
		$$ = append([]Statement{$1}, $2...)
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			l.statements = $$
		}
	}

statement
	: expr LINE_BREAK
	{
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			$$ = &ExpressionStatement{Expression: $1}

			sqlExpr := $1.(*SQLExpression)
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
	| expr ';'
	{
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			$$ = &ExpressionStatement{Expression: $1}

			sqlExpr := $1.(*SQLExpression)
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

expr
	: LINE_BREAK
	{
		$$ = &LineBreakExpression{Literal: $1.lit}
	}
	| IDENT '[' NUMBER '-' NUMBER ']'
	{
		$$ = &SQLExpression{
			Literal: $1.lit,
			Value: "",
			PositionExpression: &PositionExpression{
				Start: $3.lit,
				End:   $5.lit,
			},
			Depth: 0,
		}
	}
	| IDENT VALUE '[' NUMBER '-' NUMBER ']'
	{
		$$ = &SQLExpression{
			Literal: $1.lit,
			Value: $2.lit,
			PositionExpression: &PositionExpression{
				Start: $4.lit,
				End:   $6.lit,
			},
			Depth: 0,
		}
	}
	| INDENT IDENT '[' NUMBER '-' NUMBER ']'
	{
		depth := len($1.lit) / 2
		$$ = &SQLExpression{
			Literal: $2.lit,
			Value: "",
			PositionExpression: &PositionExpression{
				Start: $4.lit,
				End:   $6.lit,
			},
			Depth: depth,
		}
	}
	| INDENT IDENT VALUE '[' NUMBER '-' NUMBER ']'
	{
		depth := len($1.lit) / 2
		$$ = &SQLExpression{
			Literal: $2.lit,
			Value: $3.lit,
			PositionExpression: &PositionExpression{
				Start: $5.lit,
				End:   $7.lit,
			},
			Depth: depth,
		}
	}

%%

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
		l: l,
		Stack: &Stack{},
	}

	if yyParse(&w) != 0 {
		panic("Parse error")
	}

	return w.Stack
}

func Reset() {
	currentDepth = 0
}

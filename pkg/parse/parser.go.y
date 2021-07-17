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

var current_depth = 0

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
%token<tok> EOF IDENT NUMBER INDENT LINE_BREAK

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
			sqlExpr := $1.(*SQLExpression)

			depth := sqlExpr.Depth
			$$ = &ExpressionStatement{Expression: $1}
			se := &SQLExpression{Literal: sqlExpr.Literal, Value: sqlExpr.Value, Depth: depth}
			if depth == 0 {
				l.Stack.Push(se)
			} else if current_depth > depth {
					v := current_depth - depth
					for i := 0; i < v+1; i++ {
							l.Stack.Pop()
					}
					node := l.Stack.Pop()
					node.Expressions = append(node.Expressions, se)

					l.Stack.Push(node)
					l.Stack.Push(se)

					current_depth = depth
			} else {
					parent := l.Stack.Pop()
					if parent != nil {
						parent.Expressions = append(parent.Expressions, se)
						l.Stack.Push(parent)
					}
					l.Stack.Push(se)

					current_depth++
			}
		}
	}
	| expr ';'
	{
		if l, isLexerWrapper := yylex.(*LexerWrapper); isLexerWrapper {
			sqlExpr := $1.(*SQLExpression)

			depth := sqlExpr.Depth
			$$ = &ExpressionStatement{Expression: $1}
			se := &SQLExpression{Literal: sqlExpr.Literal, Value: sqlExpr.Value, Depth: depth}
			if depth == 0 {
				l.Stack.Push(se)
			} else if current_depth > depth {
					v := current_depth - depth
					for i := 0; i < v+1; i++ {
							l.Stack.Pop()
					}
					node := l.Stack.Pop()
					node.Expressions = append(node.Expressions, se)

					l.Stack.Push(node)
			} else {
					parent := l.Stack.Pop()
					if parent != nil {
						parent.Expressions = append(parent.Expressions, se)
						l.Stack.Push(parent)
					}
					l.Stack.Push(se)
			}
		}
	}

expr
	: NUMBER
	{
		$$ = &NumberExpression{Literal: $1.lit}
	}
	| IDENT
	{
		$$ = &IdentifierExpression{Literal: $1.lit}
	}
	| INDENT
	{
		$$ = &IndentExpression{Literal: $1.lit}
	}
	| LINE_BREAK
	{
		$$ = &LineBreakExpression{Literal: $1.lit}
	}
	| IDENT '[' expr '-' expr ']'
	{
		$$ = &SQLExpression{Literal: $1.lit,  Depth: 0}
	}
	| IDENT '(' expr ')' '[' expr '-' expr ']'
	{
		$$ = &SQLExpression{Literal: $1.lit, Value: $3, Depth: 0}
	}
	| INDENT IDENT '[' expr '-' expr ']'
	{
		depth := len($1.lit) / 2
		$$ = &SQLExpression{Literal: $2.lit, Depth: depth}
	}
	| INDENT IDENT '(' expr ')' '[' expr '-' expr ']'
	{
		depth := len($1.lit) / 2
		$$ = &SQLExpression{Literal: $2.lit, Value: $4, Depth: depth}
	}

%%

type LexerWrapper struct {
	s          *Lexer
	recentLit  string
	recentPos  Position
	statements []Statement
	Stack *Stack
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.NextToken()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *LexerWrapper) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Lexer) *Stack {
	l := LexerWrapper{
		s: s,
		Stack: &Stack{},
	}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}

	return l.Stack
}

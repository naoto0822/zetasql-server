package parse

type (
	Statement interface {
		statement()
	}

	Expression interface {
		expression()
	}
)

type (
	ExpressionStatement struct {
		Expression Expression
	}
)

func (x *ExpressionStatement) statement() {}

type (
	NumberExpression struct {
		Literal string
	}

	IdentifierExpression struct {
		Literal string
	}

	IndentExpression struct {
		Literal string
	}

	LineBreakExpression struct {
		Literal string
	}

	SQLExpression struct {
		Literal            string
		Value              Expression
		PositionExpression *PositionExpression
		Depth              int

		Expressions []*SQLExpression
	}

	PositionExpression struct {
		Start int
		End   int
	}
)

func (x *NumberExpression) expression()     {}
func (x *IdentifierExpression) expression() {}
func (x *IndentExpression) expression()     {}
func (x *LineBreakExpression) expression()  {}
func (x *SQLExpression) expression()        {}

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

//
type (
	IdentifierExpression struct {
		Literal string
	}

	ValueExpression struct {
		Literal string
	}

	NumberExpression struct {
		Literal string
	}

	PositionExpression struct {
		Start string
		End   string
	}
)

type (
	IndentExpression struct {
		Literal string
	}

	LineBreakExpression struct {
		Literal string
	}

	SQLExpression struct {
		Literal            string
		Value              string
		PositionExpression *PositionExpression
		Depth              int

		Expressions []*SQLExpression
	}
)

func (x *NumberExpression) expression()     {}
func (x *IdentifierExpression) expression() {}
func (x *IndentExpression) expression()     {}
func (x *LineBreakExpression) expression()  {}
func (x *SQLExpression) expression()        {}
func (e *ValueExpression) expression()      {}
